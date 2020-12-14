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
	r.Get("/", homeHandler(app))
	r.Get("/{peopleId}", PeopleHanler(app))

	return r
}

func PeopleHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peopleId := chi.URLParam(r, "peopleId")
		peopleDataHistory := GetPeopleDataHistory(peopleId)

		//We are checking if the previous method returns nothing, we should redirect people home page
		//TODO we need to implement error reporter on People Home Page
		if peopleDataHistory.History.Id == "" {
			//app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/people", 301)
			return
		}
		type PageData struct {
			PeopleDataHistory PeopleDataHistory
			GalleryString     []string
		}

		data := PageData{peopleDataHistory, GetpeopleGallery(peopleId)}
		files := []string{
			app.Path + "people/people_single.html",
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

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peoples, err := people_io.ReadCategories()
		if err != nil {
			fmt.Println(err, " There is an error when reading all the category")
		}
		peopleData := GetPeopleBriefData()
		type PageData struct {
			Peoples    []people.Category
			PeopleData []PeopleBriefData
		}

		data := PageData{peoples, peopleData}
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
