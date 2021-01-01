package people

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/people"
	"ostmfe/io/pageData_io"
	"ostmfe/io/people_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))
	r.Get("/{peopleId}", PeopleHanler(app))

	return r
}

func PeopleHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peopleId := chi.URLParam(r, "peopleId")
		peopleDataHistory := GetPeopleDataHistory(peopleId)

		//We are checking if the previous method returns nothing, we should redirect people home page
		//TODO we need to implement error reporter on People Home Page
		if peopleDataHistory.History.Id == "" {
			//app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/people", 301)
			return
		}
		type PageData struct {
			PeopleDataHistory PeopleDataHistory
			GalleryString     []string
		}

		data := PageData{peopleDataHistory, GetpeopleGallery(peopleId)}
		files := []string{
			app.Path + "people/people_single.html",
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

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
