package about_us

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	comment2 "ostmfe/controller/comment"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/domain/event"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	member2 "ostmfe/domain/member"
	"ostmfe/domain/pages"
	"ostmfe/domain/user"
	"ostmfe/io/comment_io"
	"ostmfe/io/event_io"
	"ostmfe/io/group_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/member_io"
	"ostmfe/io/pageData_io"
	"ostmfe/io/pages/client"
	"ostmfe/io/user_io"
	"time"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	//r.Get("/", homeHanler(app))
	r.Get("/", homeHandler(app))
	r.Get("/group", GrouphomeHanler(app))
	r.Get("/single/{groupId}", GroupHanler(app))
	r.Post("/create-comment", CreateComment(app))
	r.Post("/create-member", CreateMemberHandler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))
	return r
}

func CreateMemberHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		name := r.PostFormValue("name")
		surname := r.PostFormValue("surname")
		email := r.PostFormValue("email")
		address := r.PostFormValue("address")
		subject := r.PostFormValue("subject")
		Messege := r.PostFormValue("message")
		groupId := r.PostFormValue("groupId")

		if name != "" && email != "" && surname != "" {
			member := member2.Member{"", email, name, surname, subject, Messege, address}
			fmt.Println("email: ", member)
			newMember, err := member_io.CreateMember(member)
			if err != nil {
				fmt.Println("error creating member")
			} else {
				groupMember := group.GroupMember{"", newMember.Id, misc.FormatDateTime(time.Now()), groupId}
				fmt.Println("Group Member: ", groupMember)
				_, err := group_io.CreateGroupMember(groupMember)
				if err != nil {
					fmt.Println("error creating groupMember")
				}
			}
		}
		http.Redirect(w, r, "/about_us/single/"+groupId, 301)
	}
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		aboutUsPageData, err := client.AboutClientPage()
		if err != nil {
			fmt.Println(err, " error reading about us page data")
		}
		type PageData struct {
			Data pages.AboutUsPageData
		}

		data := PageData{aboutUsPageData}
		files := []string{
			app.Path + "about_us/about_us.html",
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

func CreateComment(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		groupId := r.PostFormValue("groupId")

		if name != "" && email != "" && message != "" {
			commentObject := comment.Comment{"", email, name, misc.FormatDateTime(time.Now()), misc.ConvertToByteArray(message), "", false}
			newComment, err := comment_io.CreateComment(commentObject)
			if err != nil {
				fmt.Println("error creating comment")
			} else {
				_, err := comment_io.CreateCommentGroup(comment.CommentGroup{"", groupId, newComment.Id})
				if err != nil {
					fmt.Println("error creating comment")
				}
				http.Redirect(w, r, "/about_us/single/"+groupId, 301)
			}
		}
		http.Redirect(w, r, "/about_us/single/"+groupId, 301)
	}
}

func GrouphomeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var bannerImage string
		banner, err := misc.GetBanner("Group-home")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = banner.Id
		}
		type PageData struct {
			ProjectBanner string
			Groups        []GroupData
			Histories     []misc.HistoryAndProfile
		}
		data := PageData{bannerImage, GetGroupData(), misc.ReadHistoryWithImages()}
		files := []string{
			app.Path + "about_us/group-home.html",
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

func GroupHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bannerImage string
		var files []string
		banner, err := misc.GetBanner("Group-home")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = banner.Id
		}
		groupId := chi.URLParam(r, "groupId")
		groupDataHistory := GetGroupDataHistory(groupId)

		singleData, err := client.GetGroupClientSingleData(groupId)
		if err != nil {
			fmt.Println("error getting groupData")
			app.InfoLog.Println(err, " error getting groupData")
			http.Redirect(w, r, "/about_us/group", 301)
			return
		}
		//We are checking if the previous method returns nothing, we should redirect people home page
		//TODO we need to implement error reporter on People Home Page
		if groupDataHistory.History.Id == "" {
			app.InfoLog.Println(" error getting group History")
			fmt.Println(" error getting group History")
			//app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/about_us/group", 301)
			return
		}
		eventCommentNumber, err := comment_io.CountCommentGroup(groupId)
		if err != nil {
			fmt.Println("error reading counting CommentGroup")
		}
		commentes := comment2.GetGroupComment(groupId)

		fmt.Println("Comments: ", commentes)
		type PageData struct {
			ProjectBanner    string
			PageDatas        pages.GroupClientSingleData
			GroupDataHistory GroupDataHistory
			EventData        []EventData
			GalleryImages    []misc.GroupGalleryImages
			Comments         []comment.CommentStack
			CommentNumber    int64
			NumberOfComments int64
		}
		data := PageData{bannerImage, singleData, groupDataHistory, getEventsData(groupId), misc.GetGroupGallery(groupId), commentes, eventCommentNumber, eventCommentNumber}

		if singleData.Group.Name == "The Phoenix Committee" {
			files = []string{
				app.Path + "about_us/pheonix_group.html",
				app.Path + "base_templates/navigator.html",
				app.Path + "base_templates/footer.html",
				app.Path + "base_templates/comments.html",
			}
		} else {
			files = []string{
				app.Path + "about_us/groups_single.html",
				app.Path + "base_templates/navigator.html",
				app.Path + "base_templates/footer.html",
				app.Path + "base_templates/comments.html",
			}
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

type EventData struct {
	Event event.Event
	Image image.Images
}

//This method returns events with it picture for an groupId.
func getEventsData(groupId string) []EventData {
	var eventData []EventData

	groupEvents, err := event_io.ReadEventGroupWithGroupId(groupId)
	if err != nil {
		fmt.Println(err, " error reading EventGroups")
		return eventData
	}
	for _, groupEvent := range groupEvents {
		event, err := event_io.ReadEvent(groupEvent.EventId)
		if err != nil {
			fmt.Println(err, " error reading Event")
		} else {
			eventImage, err := event_io.ReadEventImageOf(event.Id)
			if err != nil {
				fmt.Println(err, " error reading Event")
			} else {
				image, err := image_io.ReadImage(eventImage.ImageId)
				if err != nil {
					fmt.Println(err, " error reading groups image")
				} else {
					eventData = append(eventData, EventData{event, image})
				}
			}
		}
	}
	return eventData
}

type GroupData struct {
	Group group.Groupes
	Image image.Images
}

//With An eventId this method returns a groupEvent.

//This method returns a list of groups with their picture.
func GetGroupData() []GroupData {
	var groupDataList []GroupData
	groups, err := group_io.ReadGroups()
	if err != nil {
		fmt.Println(err, " error reading groups")
		return groupDataList
	}
	for _, groupData := range groups {
		groupImage, err := group_io.ReadGroupImageWithGroupId(groupData.Id)
		if err != nil {
			fmt.Println(err, " error reading groups image")
		} else {
			image, err := image_io.ReadImage(groupImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading groups image")
			} else {
				groupDataObject := GroupData{groupData, image}
				groupDataList = append(groupDataList, groupDataObject)
				groupDataObject = GroupData{}
			}
		}
	}
	return groupDataList
}

type StaffData struct {
	User    user.Users
	Image   image.ImagesHelper
	History history2.HistoriesHelper
}

func getStaff() []StaffData {
	var staffs []StaffData
	var roleId string

	roles, err := user_io.ReadRoles()
	if err != nil {
		fmt.Println(err, " error reading roles")
		return staffs
	} else {
		for _, role := range roles {
			if role.Role == "staff" {
				roleId = role.Id
				fmt.Println(roleId, " roleId")
			}
		}
		userRoles, err := user_io.ReadUserRoleAllOf(roleId)
		if err != nil {
			fmt.Println(err, " error userRoles")
		} else {
			for _, userRole := range userRoles {
				users, err := user_io.ReadUser(userRole.Email)
				if err != nil {
					fmt.Println(err, " error user")
				} else {
					userImage, err := user_io.ReadUserImageWithEmail(users.Email)
					if err != nil {
						fmt.Println(err, " error userImage")
					} else {
						imageObject, erri := image_io.ReadImage(userImage.ImageId)
						history, err := history_io.ReadHistorie(userImage.HistoryId)
						if err != nil && erri != nil {
							fmt.Println(err, " error reading image and history")
						} else {
							imageHelper := image.ImagesHelper{imageObject.Id, misc.ConvertingToString(imageObject.Image), imageObject.Description, userImage.Id}
							historyHelper := history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}

							staffs = append(staffs, StaffData{users, imageHelper, historyHelper})
						}
					}
				}
			}
		}
	}
	return staffs

}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var bannerImage string
		pageBanner, err := pageData_io.ReadPageBannerWIthPageName("event-page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = misc.GetBannerImage(pageBanner.BannerId)
		}
		type PageData struct {
			Groups        []GroupData
			Staffs        []StaffData
			PageSections  AboutUsPageSection
			AboutUsBanner string
			GalleryImages []misc.GroupGalleryImages
		}

		data := PageData{GetGroupData(), getStaff(), getPageData("aboutUs"), bannerImage, misc.GetAllGroupGallery()}
		files := []string{
			app.Path + "about_us/about_us.html",
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

type AboutUsPageSection struct {
	Introduction     string
	StaffIntroTitle  string
	MemberIntroTitle string
}

func getPageData(pageName string) AboutUsPageSection {
	var introduction string
	var staffIntroTitle string
	var memberIntroTitle string
	//var exibitionintro string

	page, err := pageData_io.ReadPageDataWIthName(pageName)
	if err != nil {
		fmt.Println(err, " error reading page, this may not exist")
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
				if pageSection.SectionName == "Introduction" {
					//fmt.Println(" Introduction",pageSection)
					introduction = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "StaffIntroTitle" {
					//fmt.Println(" StaffIntroTitle",pageSection)
					staffIntroTitle = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "MemberIntroTitle" {
					//fmt.Println(" MemberIntroTitle",pageSection)
					memberIntroTitle = misc.ConvertingToString(pageDateSection.Content)
				}

			}
		}
	}
	return AboutUsPageSection{introduction, staffIntroTitle, memberIntroTitle}
}
