package about_us

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/domain/group"
	"ostmfe/domain/image"
	"ostmfe/io/group_io"
	"ostmfe/io/image_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/single/{groupId}", GroupHanler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))
	return r
}

func GroupHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupId := chi.URLParam(r, "groupId")
		groupDataHistory := GetGroupDataHistory(groupId)

		//We are checking if the previous method returns nothing, we should redirect people home page
		//TODO we need to implement error reporter on People Home Page
		if groupDataHistory.History.Id == "" {
			//app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/group", 301)
			return

		}
		type PageData struct {
			GroupDataHistory GroupDataHistory
		}
		data := PageData{groupDataHistory}
		files := []string{
			app.Path + "about_us/groups_single.html",
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

type GroupData struct {
	Group group.Groups
	Image image.Images
}

func getGroupData() []GroupData {
	var goupDatas []GroupData
	groups, err := group_io.ReadGroups()
	if err != nil {
		fmt.Println(err, " error reading groups")
		return goupDatas
	}
	for _, group := range groups {
		groupImage, err := group_io.ReadGroupImageWithGroupId(group.Id)
		if err != nil {
			fmt.Println(err, " error reading groups image")
		} else {
			image, err := image_io.ReadImage(groupImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading groups image")
			} else {
				groupDataObject := GroupData{group, image}
				goupDatas = append(goupDatas, groupDataObject)
				groupDataObject = GroupData{}
			}
		}
	}
	return goupDatas
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			Groups []GroupData
		}
		data := PageData{getGroupData()}
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
