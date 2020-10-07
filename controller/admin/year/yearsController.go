package year

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	museum "ostmfe/domain"
	"ostmfe/io"
	"strconv"
)

func YearHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", HomeHandler(app))
	r.Get("/delete/{yearId}", DeleteHandler(app))
	r.Post("/create", CreateYear(app))
	return r
}

func DeleteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		yearId := chi.URLParam(r, "yearId")

		_, err := io.DeleteYear(yearId)
		if err != nil {
			fmt.Println(err, " something went wrong! could not delete a year.")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/years", 301)
			return
		}
		fmt.Println(" successfully deleted")
		app.Session.Put(r.Context(), "creation-successful", "You have successful optional")
		http.Redirect(w, r, "/admin_user/years", 301)
		return
	}
}

func CreateYear(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		year, _ := strconv.Atoi(r.PostFormValue("year"))

		if year != 0 {
			yearObject := museum.Years{"", year}
			_, err := io.CreateYear(yearObject)
			if err != nil {
				fmt.Println(err, " something went wrong! could not update history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/years", 301)
				return
			}
			fmt.Println(" successfully created")
			app.Session.Put(r.Context(), "creation-successful", "You have successful optional")
			http.Redirect(w, r, "/admin_user/years", 301)
			return
		}
	}
}

func HomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var unknown_error string
		var backend_error string
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			unknown_error = app.Session.GetString(r.Context(), "creation-unknown-error")
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			backend_error = app.Session.GetString(r.Context(), "user-create-error")
			app.Session.Remove(r.Context(), "user-create-error")
		}

		years, err := io.ReadYears()
		if err != nil {
			fmt.Println(err, " error reading years")
		}

		type PagePage struct {
			Backend_error string
			Unknown_error string
			Years         []museum.Years
			SidebarData   misc.SidebarData
		}

		data := PagePage{backend_error, unknown_error, years, misc.GetSideBarData("event", "year")}
		files := []string{
			app.Path + "admin/event/years-page.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
			app.Path + "base_templates/footer.html",
			app.Path + "admin/template/cards.html",
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
