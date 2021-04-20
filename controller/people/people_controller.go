package people

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	comment2 "ostmfe/controller/comment"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/people"
	"ostmfe/io/comment_io"
	"ostmfe/io/image_io"
	"ostmfe/io/pageData_io"
	"ostmfe/io/people_io"
	"time"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))
	r.Post("/create-comment", CreateCommentHandler(app))
	r.Get("/{peopleId}", PeopleHanler(app))
	r.Get("/category/{categoryId}", PeopleCategoryHanler(app))

	return r
}

func PeopleCategoryHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryId := chi.URLParam(r, "categoryId")
		fmt.Println(categoryId)

		peoples, err := people_io.ReadCategories()
		var bannerImage string
		if err != nil {
			fmt.Println(err, " There is an error when reading all the category")
		}
		pageBanner, err := pageData_io.ReadPageBannerWIthPageName("people-page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = misc.GetBannerImage(pageBanner.BannerId)
		}
		peopleData := getAllPeopleDataOfCategory(categoryId)
		type PageData struct {
			Peoples      []people.Category
			PeopleData   []PeopleBriefData
			PeoplePage   PeoplePage
			PeopleBanner string
		}

		data := PageData{peoples, peopleData, getPageData(), bannerImage}
		files := []string{
			app.Path + "people/people_home.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func CreateCommentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		historyId := r.PostFormValue("historyId")

		if historyId != "" && email != "" && message != "" {
			commentObject := comment.Comment{"", email, name, misc.FormatDateTime(time.Now()), misc.ConvertToByteArray(message), "", false}
			newComment, err := comment_io.CreateComment(commentObject)
			if err != nil {
				fmt.Println("error creating comment")
			} else {
				_, err := comment_io.CreateCommentPeople(comment.CommentPeople{"", historyId, newComment.Id, "false"})
				if err != nil {
					fmt.Println("error creating comment")
				}
				http.Redirect(w, r, "/history/single_history/"+historyId, 301)
			}
		}
		http.Redirect(w, r, "/history/single_history/"+historyId, 301)
	}
}

func PeopleHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peopleId := chi.URLParam(r, "peopleId")
		peopleDataHistory := GetPeopleDataHistory(peopleId)
		var bannerImage string

		//We are checking if the previous method returns nothing, we should redirect people home page
		//TODO we need to implement error reporter on People Home Page
		if peopleDataHistory.History.Id == "" {
			//app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/people", 301)
			return
		}
		peopleComment, err := comment_io.CountCommentPeople(peopleId)
		if err != nil {
			fmt.Print(err, " error counting people comments")
		}
		pageBanner, err := pageData_io.ReadPageBannerWIthPageName("people-page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = misc.GetBannerImage(pageBanner.BannerId)
		}
		type PageData struct {
			PeopleDataHistory PeopleDataHistory
			GalleryString     []string
			CommentNumber     int64
			Comments          []comment.CommentStack
			GalleryImages     []misc.PeopleGalleryImages
			PeopleBanner      string
		}

		data := PageData{peopleDataHistory, GetpeopleGallery(peopleId), peopleComment, comment2.GetAllPeopleComments(peopleId), misc.GetPeopleGallery(peopleId), bannerImage}
		files := []string{
			app.Path + "people/people_single.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
			app.Path + "base_templates/client-gallery.html",
			app.Path + "base_templates/comments.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peoples, err := people_io.ReadCategories()
		if err != nil {
			fmt.Println(err, " error reading peoples.")
		}
		var bannerImage string
		if err != nil {
			fmt.Println(err, " There is an error when reading all the category")
		}
		pageBanner, err := pageData_io.ReadPageBannerWIthPageName("people-page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = misc.GetBannerImage(pageBanner.BannerId)
		}
		peopleData := GetPeopleBriefData()
		type PageData struct {
			Peoples      []people.Category
			PeopleData   []PeopleBriefData
			PeoplePage   PeoplePage
			PeopleBanner string
		}

		data := PageData{peoples, peopleData, getPageData(), bannerImage}
		files := []string{
			app.Path + "people/people_home.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

type PeoplePage struct {
	Banner string
	Intro  string
}

func getPageData() PeoplePage {
	var intro string
	var banner string
	page, err := pageData_io.ReadPageDataWIthName("people-page")
	if err != nil {
		fmt.Println(err, " error reading page")
	} else {
		pageDateSectionObject, err := pageData_io.ReadPageSectionAllOf(page.Id)
		if err != nil {
			fmt.Println(err, " error reading page")
		}
		for _, pageDateSection := range pageDateSectionObject {
			pageSection, err := pageData_io.ReadSection(pageDateSection.SectionId)
			if err != nil {
				fmt.Println(err, " error reading page")
			} else {
				if pageSection.SectionName == "banner" {
					fmt.Println(" banner", pageSection)
					banner = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "intro" { //todo The only section in this page is INTRO!!!
					fmt.Println(" sahoContent", pageSection)
					intro = misc.ConvertingToString(pageDateSection.Content)
				}
			}
		}
	}
	return PeoplePage{banner, intro}
}

func getAllPeopleDataOfCategory(categoryId string) []PeopleBriefData {
	var peopleBriefDatas []PeopleBriefData
	var image image3.Images
	peopleCategorys, err := people_io.ReadPeopleCategoryWithCategoryId(categoryId)
	if err != nil {
		fmt.Println(err, " error reading peopleCategoryByCategoryId")
		return peopleBriefDatas
	}

	for _, peopleCategory := range peopleCategorys {
		people, err := people_io.ReadPeople(peopleCategory.PeopleId)
		if err != nil {
			fmt.Println(err, " couldn't read people")
		} else {
			peopleImage, err := people_io.ReadPeopleImageWithPeopleId(people.Id)
			if err != nil {
				fmt.Println(err, " couldn't read peopleImage")
			}
			image, err = image_io.ReadImage(peopleImage.ImageId)
			if err != nil {
				fmt.Println(err, " couldn't read image")
			} else {
				peopleBriefData := PeopleBriefData{people, image}
				peopleBriefDatas = append(peopleBriefDatas, peopleBriefData)
			}
		}

	}
	return peopleBriefDatas
}
