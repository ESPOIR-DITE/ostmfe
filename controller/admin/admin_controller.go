package admin

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	collection2 "ostmfe/controller/admin/collection"
	"ostmfe/controller/admin/event"
	"ostmfe/controller/admin/group"
	"ostmfe/controller/admin/histories"
	"ostmfe/controller/admin/parteners"
	"ostmfe/controller/admin/peoples"
	"ostmfe/controller/admin/places"
	project3 "ostmfe/controller/admin/project"
	"ostmfe/controller/admin/users"
)

/***
- user-create-error : This is session message reporting an error occurred when there is an error when creating a new USER.
- creation-successful : This is session message reporting an successful creation.
*/

func Home(app *config.Env) http.Handler {
	mux := chi.NewMux()

	mux.Handle("/", homeHanler(app))

	mux.Mount("/users", users.UserController(app))

	mux.Mount("/role", users.RoleController(app))

	mux.Mount("/event", event.EventHome(app))

	mux.Mount("/project", project3.ProjectHome(app))

	mux.Mount("/partner", parteners.PartnerHome(app))

	mux.Mount("/place", places.PlaceHome(app))

	mux.Mount("/collection", collection2.CollectionHome(app))

	mux.Mount("/history", histories.HistoryHome(app))

	mux.Mount("/people", peoples.PeopleHome(app))

	mux.Mount("/group", group.GroupHome(app))

	return mux
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var success_notice string
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			success_notice = app.Session.GetString(r.Context(), "creation-successful")
			app.Session.Remove(r.Context(), "creation-successful")
		}
		type PageData struct {
			Success_notice string
		}
		data := PageData{success_notice}
		files := []string{
			app.Path + "admin/admin.html",
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
