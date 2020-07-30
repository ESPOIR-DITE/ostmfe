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
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

type GroupData struct {
	Group group.Group
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
