package contribution

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	"ostmfe/domain/contribution"
	"ostmfe/io/contribution_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Post("/createFileType", CreateFileType(app))
	return r
}

func CreateFileType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}

		r.ParseForm()
		fileType := r.PostFormValue("fileType")

		if fileType != "" {
			contributionFileType := contribution.ContributionType{"", fileType}
			_, err := contribution_io.CreateContributionFileType(contributionFileType)
			if err != nil {
				fmt.Println("error creating file type")
				http.Redirect(w, r, "/admin_user/contribution/", 301)
			}
		}
		http.Redirect(w, r, "/admin_user/contribution/", 301)
	}
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var success_notice string
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			success_notice = app.Session.GetString(r.Context(), "creation-successful")
			app.Session.Remove(r.Context(), "creation-successful")
		}
		contributions, err := contribution_io.ReadContributions()
		if err != nil {
			fmt.Println("error reading contributions")
		}
		contributionFileTypes, err := contribution_io.ReadContributionFileTypes()
		if err != nil {
			fmt.Println("error reading contributions types")
		}
		type PageData struct {
			Success_notice    string
			SidebarData       misc.SidebarData
			Contribution      []contribution.Contribution
			ContributionTypes []contribution.ContributionType
		}
		data := PageData{success_notice, misc.GetSideBarData("contributor", ""), contributions, contributionFileTypes}
		files := []string{
			app.Path + "admin/settings/admin-settings.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
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
