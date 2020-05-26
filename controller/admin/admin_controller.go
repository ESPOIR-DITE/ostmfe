package admin

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	event2 "ostmfe/domain/event"
	place2 "ostmfe/domain/place"
	project2 "ostmfe/domain/project"
	user2 "ostmfe/domain/user"
	"ostmfe/io/event_io"
	"ostmfe/io/place_io"
	"ostmfe/io/project_io"
	"ostmfe/io/user_io"
	"time"
)

/***
- user-create-error : This is session message reporting an error occurred when there is an error when creating a new USER.
- creation-successful : This is session message reporting an successful creation.
*/

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))

	r.Get("/users", UserHandler(app))
	r.Get("/users/new", NewUserHandler(app))
	r.Get("/users/edit", EditUserHandler(app))
	r.Post("/users/create", CreateUserHandler(app))

	r.Get("/event", EventsHandler(app))
	r.Get("/event/new", NewEventsHandler(app))
	r.Get("/event/edite", EditEventsHandler(app))
	r.Post("/event/create", CreateEventHandler(app))

	r.Get("/project", ProjectsHandler(app))
	r.Get("/project/new", NewProjectsHandler(app))
	r.Get("/project/edit", EditeProjectsHandler(app))
	r.Post("/project/create", CreateProjectHandler(app))
	//r.Get("/projects/delete",DeleteProjectsHandler(app))

	r.Get("/place", PlacesHandler(app))
	r.Get("/place/new", NewPlacesHandler(app))
	r.Get("/place/edit", EditPlacesHandler(app))

	r.Get("/collection", CollectionHandler(app))
	r.Get("/collection/new", NewCollectionHandler(app))
	r.Get("/collection/edit", EditCollectionHandler(app))

	r.Get("/history", HistoryHandler(app))
	r.Get("/history/new", NewHistoryHandler(app))
	r.Get("/history/edit", EditHistoryHandler(app))

	return r
}
func contains(slice []string) [][]byte {
	var set [][]byte
	for _, s := range slice {
		v := []byte(s)
		set = append(set, v)
	}
	return set
}
func CreateProjectHandler(app *config.Env) http.HandlerFunc {
	/***
	Here we create a new project

	*/
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//fileslist := r.Form["file"]

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		project_name := r.PostFormValue("project_name")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>>>>")
		}

		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)
		//reader := bufio.NewReader(file)
		//reader2 := bufio.NewReader(file2)
		//reader3 := bufio.NewReader(file3)
		//reader4 := bufio.NewReader(file4)
		//reader5 := bufio.NewReader(file5)
		//reader6 := bufio.NewReader(file6)
		//
		//content, _ := ioutil.ReadAll(reader)
		//content2, _ := ioutil.ReadAll(reader2)
		//content3, _ := ioutil.ReadAll(reader3)
		//content4, _ := ioutil.ReadAll(reader4)
		//content5, _ := ioutil.ReadAll(reader5)
		//content6, _ := ioutil.ReadAll(reader6)
		fmt.Println(project_name, "<<<Project Name|| description>>>", description)

		if project_name != "" && description != "" {
			project := project2.Project{"", project_name, description}
			new_project, err := project_io.CreateProject(project)
			if err != nil {
				fmt.Println(err, " could not create project Line: 190")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new", 301)
				return
			}
			//COnverting file(byte arrays) into Array of byte
			//sliceOfImage := [][]byte{content, content2, content3, content4, content5, content6}

			projectImage := project2.ProjectImage{"", new_project.Id, "", ""}
			helper := project2.ProjectImageHelper{filesByteArray, projectImage}
			_, errr := project_io.CreateProjectImage(helper)
			if errr != nil {
				fmt.Println(err, " error creating projectImage")
				_, err := project_io.DeleteProject(new_project.Id)
				if err != nil {
					fmt.Println(err, " error deleting project")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+project_name)
			http.Redirect(w, r, "/admin_user", 301)
			return
			//event_name := r.PostFormValue("event_name")
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/project/new", 301)
		return

	}

}

func CreateEventHandler(app *config.Env) http.HandlerFunc {
	/***
	Here we create a new event

	*/
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		event_name := r.PostFormValue("event_name")
		date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		project := r.PostFormValue("project")
		partner := r.PostFormValue("partner")
		latlng := r.PostFormValue("latlng")
		place := r.PostFormValue("place")

		if event_name != "" {
			eventObject := event2.Event{"", event_name, date}
			newEvent, err := event_io.CreateEvent(eventObject)
			if err != nil {
				fmt.Println(err, " error when creating a new event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			}
			eventPartner := event2.EventPartener{"", partner, newEvent.Id, ""}
			_, err = event_io.CreateEventPartener(eventPartner)
			if err != nil {
				fmt.Println(err, " error when creating event partner")
				/**
				Rolling back
				*/
				_, err := event_io.DeleteEvent(newEvent.Id)
				if err != nil {
					fmt.Println(err, " error when deleting event in rolling back action")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			}

			eventProject := event2.EventProject{project, newEvent.Id, ""}
			_, err = event_io.CreateEventProject(eventProject)
			if err != nil {
				fmt.Println(err, " error when creating event project")
				/**
				Rolling back
				*/
				_, err := event_io.DeleteEvent(newEvent.Id)
				if err != nil {
					fmt.Println(err, " error when deleting event in rolling back action")
				}
				_, errr := event_io.DeleteEventPartener(eventPartner.PartenerId)
				if errr != nil {
					fmt.Println(errr, " error when deleting event partner in rolling back action")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred when, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			}
			latitude, longitude := misc.SeparateLatLng(latlng)
			place := place2.Place{"", place, latitude, longitude}
			newPlace, err := place_io.CreatePlace(place)
			if err != nil {
				fmt.Println(err, " error when creating a new place")

				_, errr := event_io.DeleteEventPartener(eventPartner.PartenerId)
				if errr != nil {
					fmt.Println(errr, " error when deleting event partner in rolling back action")
				}
				_, err := event_io.DeleteEvent(newEvent.Id)
				if err != nil {
					fmt.Println(err, " error when deleting event in rolling back action")
				}
				_, errrr := event_io.DeleteEventProject(eventProject.EventId)
				if errrr != nil {
					fmt.Println(err, " error when deleting Project in rolling back action")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return

			} else {
				eventPlace := event2.EventPlace{newPlace.Id, eventObject.Id, ""}
				_, err := event_io.CreateEventPlace(eventPlace)
				if err != nil {
					fmt.Println(err, " error when creating Event place")
					_, errr := event_io.DeleteEventPartener(eventPartner.PartenerId)
					if errr != nil {
						fmt.Println(errr, " error when deleting event partner in rolling back action")
					}
					_, err := event_io.DeleteEvent(newEvent.Id)
					if err != nil {
						fmt.Println(err, " error when deleting event in rolling back action")
					}
					_, errrr := event_io.DeleteEventProject(eventProject.EventId)
					if errrr != nil {
						fmt.Println(err, " error when deleting event Project in rolling back action")
					}
					_, errrrr := place_io.DeletePlace(newPlace.Id)
					if errrrr != nil {
						fmt.Println(err, " error when deleting place in rolling back action")
					}
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/users/new", 301)
					return
				}
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Event : "+event_name)
			http.Redirect(w, r, "/admin_user", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/users/new", 301)
		return
	}

}

func CreateUserHandler(app *config.Env) http.HandlerFunc {
	/***
	we are getting the form from html
	we grab all the fields corresponding to the name assigned to them
	we create an object with the records collected from the html
	we then send the object to the backend, if an error occurs we will redirect back to new user html file to try again.
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("name")
		surname := r.PostFormValue("surname")
		email := r.PostFormValue("email")
		fmt.Println(name, "<<name  surname>>", surname, "  email>>", email)
		if name != "" && surname != "" && email != "" {
			user := user2.Users{email, name, surname}
			newUser, err := user_io.CreateUser(user)
			if err != nil {
				fmt.Println(err, "error creating new user line: 57")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			} else {
				fmt.Println(err, "Creation of a new user successful")
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
					return
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new user : "+newUser.Name)
				http.Redirect(w, r, "/admin_user", 301)
				return
			}
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/users/new", 301)
		return
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
		files := []string{
			app.Path + "admin/collection/new_history.html",
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

func HistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/history.html",
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

func EditCollectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/edit_collection.html",
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

func NewCollectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/new_collection.html",
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

func CollectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/collections.html",
			app.Path + "admin/template/navbar.html",
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

func EditPlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/place/edit_places.html",
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

func NewPlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/place/new_places.html",
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

func PlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/place/places.html",
			app.Path + "admin/template/navbar.html",
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

func DeleteProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/edit_events.html",
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

func EditeProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/project/project_tables.html",
			app.Path + "admin/template/navbar.html",
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

func NewProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/project/new_project.html",
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

func ProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/project/projects.html",
			app.Path + "admin/template/navbar.html",
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

func EditEventsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/edit_events.html",
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

func NewEventsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/new_event.html",
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

func EventsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/events.html",
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

func EditUserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/edit_user.html",
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

func UserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/users.html",
			app.Path + "admin/template/navbar.html",
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

func NewUserHandler(app *config.Env) http.HandlerFunc {
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
			app.Path + "admin/new_user.html",
			//app.Path + "admin/template/navbar.html",
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
