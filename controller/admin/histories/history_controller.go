package histories

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	"ostmfe/io/history_io"
	"time"
)

func HistoryHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", HistoryHandler(app))
	r.Get("/new", NewHistoryHandler(app))
	r.Get("/edit", EditHistoryHandler(app))
	r.Post("/create", CreateHistpory(app))
	return r
}

func CreateHistpory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		mytextarea := r.PostFormValue("mytextarea")
		fmt.Println("Title: ", title,
			"Date: ", date,
			"description: ", description,
			"mytextArea: ", mytextarea)
		if title != "" && mytextarea != "" {
			history := history2.History{"", title, description, misc.ConvertToByteArray(mytextarea), date}
			createdHistory, err := history_io.CreateHistory(history)
			if err != nil {
				fmt.Println("error: ", err)
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/history/new", 301)
				return
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully created an new HistoryId : "+createdHistory.Title)
			http.Redirect(w, r, "/admin_user", 301)
			return
		}
	}
}
func EditHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/edit_history.html",
			//app.Path + "admin/template/navbar.html",
			//app.Path + "base_templates/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}

}

func NewHistoryHandler(app *config.Env) http.HandlerFunc {
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
		type PagePage struct {
			Backend_error string
			Unknown_error string
		}
		data := PagePage{backend_error, unknown_error}
		files := []string{
			app.Path + "admin/history/new_history.html",
			app.Path + "admin/template/navbar.html",
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

func HistoryHandler(app *config.Env) http.HandlerFunc {
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
		histories, err := history_io.ReadHistorys()
		if err != nil {
			fmt.Println(err, "Error reading Histories")
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			Histories     history2.History
		}
		data := PagePage{backend_error, unknown_error, histories}
		files := []string{
			app.Path + "admin/history/history.html",
			app.Path + "admin/template/navbar.html",
			//app.Path + "base_templates/footer.html",
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
