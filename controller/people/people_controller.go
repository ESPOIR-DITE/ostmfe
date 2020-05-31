package people

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/domain/people"
	"ostmfe/io/people_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peoples, err := people_io.ReadPeopleCategorys()
		if err != nil {
			fmt.Println(err, " There is an error when reading all the people category")
		}
		type PageData struct {
			Peoples []people.PeopleCategory
		}
		data := PageData{peoples}
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
