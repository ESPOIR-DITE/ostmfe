package visit

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Post("/book", BookHanler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func BookHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		//var newEvent event2.Event
		//event_name := r.PostFormValue("event_name")
		//date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		//project := r.PostFormValue("project")
		//description := r.PostFormValue("description")
		//partner := r.PostFormValue("partner")
		//latlng := r.PostFormValue("latlng")
		//place := r.PostFormValue("place")

		http.Redirect(w, r, "/visit", 301)
		return
	}
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "visit.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
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
