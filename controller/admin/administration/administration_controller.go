package administration

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/io/login"
)

func AdministrationController(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", LoginHandler(app))
	r.Post("/submint", SubmitLoginDetails(app))
	r.Get("/logout", AdminLogout(app))
	return r
}

func AdminLogout(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Session.Destroy(r.Context())
		http.Redirect(w, r, "/", 301)
	}
}

func SubmitLoginDetails(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		if email != "" && password != "" {
			result := login.AdminLogin(email, password)
			if result == true {
				app.Session.Put(r.Context(), "email", email)
				adminHelper.PutAdminDataInSession(app, r, email)
				http.Redirect(w, r, "/admin_user/", 301)
				return
			} else {
				http.Redirect(w, r, "/administration", 301)
				return
			}
		}
		http.Redirect(w, r, "/administration", 301)
		return
	}
}

func LoginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/login/login.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
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
