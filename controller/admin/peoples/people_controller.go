package peoples

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	people2 "ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	"ostmfe/io/history_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
	"time"
)

func PeopleHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", PeopleHandler(app))
	r.Get("/people_category/new", NewPeopleCategoryHandler(app))
	r.Get("/new", NewPeopleHandler(app))
	r.Get("/new-stp2/{peopleId}", NewPeoplestp2Handler(app))
	r.Get("/edit/{peopleId}", EditPeopleHandler(app))
	r.Get("/delete/{peopleId}", DeletePeopleHandler(app))
	r.Post("/create_stp1", CreatePeopleHandler(app))
	r.Post("/create_stp2", CreatePeopleStp2Handler(app))
	r.Post("/people_category/create", CreatePeopleCategoryHandler(app))
	return r
}
func CreatePeopleCategoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		category := r.PostFormValue("category")
		if category != "" {
			people := people2.Category{"", category}
			peopleCategory, err := people_io.CreateCategory(people)

			if err != nil {
				fmt.Println(err, " error creating people Category")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people_category/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new People Type : "+peopleCategory.Category)
			http.Redirect(w, r, "/admin_user/people_category/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people_category/new", 301)
		return
	}
}

func NewPeopleCategoryHandler(app *config.Env) http.HandlerFunc {
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

		peoples, err := people_io.ReadPeopleCategorys()
		if err != nil {
			fmt.Println(err, " There is an error when reading all the people category")
		}
		type PageData struct {
			Peoples       []people2.PeopleCategory
			Backend_error string
			Unknown_error string
		}
		data := PageData{peoples, backend_error, unknown_error}
		files := []string{
			app.Path + "admin/people/peopleType_tables.html",
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

func NewPeoplestp2Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peopleId := chi.URLParam(r, "peopleId")
		people, err := people_io.ReadPeople(peopleId)
		if err != nil {
			fmt.Println(err, " error reading the people")
			if app.Session.GetString(r.Context(), "user-read-error") != "" {
				app.Session.Remove(r.Context(), "user-read-error")
			}
			app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/people/new", 301)
			return
		}
		type PageData struct {
			People people2.People
		}
		data := PageData{People: people}
		files := []string{
			app.Path + "admin/people/image_people.html",
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

func CreatePeopleStp2Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var place place2.Place
		r.ParseForm()
		//fileslist := r.Form["file"]

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		history := r.PostFormValue("history")
		date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		peopleId := r.PostFormValue("peopleId")
		description := r.PostFormValue("description")
		latlng := r.PostFormValue("latlng")
		title := r.PostFormValue("title")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		if history != "" && peopleId != "" && title != "" && description != "" {
			historyByteArray := []byte(history)
			historyObejct := history2.History{"", title, description, historyByteArray, date}
			historynew, Err := history_io.CreateHistory(historyObejct)
			if Err != nil {
				fmt.Println(err, " error creating a new history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
				return
			}
			// I am checking if the this person is relative to a particular place
			if latlng != "" {
				latitude, longitude := misc.SeparateLatLng(latlng)
				placeObject := place2.Place{"", title, latitude, longitude, ""}
				place, err = place_io.CreatePlace(placeObject)
				if err != nil {
					fmt.Println(err, " error creating a new Place")
					_, err := history_io.DeleteHistory(historynew.Id)
					if err != nil {
						fmt.Println(err, " error could not delete history")
					}
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
					return
				}
			}

			peopleHistory := people2.PeopleHistory{"", peopleId, history}

			peopleHistoryNew, err := people_io.CreatePeopleHistory(peopleHistory)
			if err != nil {
				fmt.Println(err, " error creating a new people history")
				_, err := history_io.DeleteHistory(historynew.Id)
				if err != nil {
					fmt.Println(err, " error could not delete history")
				}
				//we need to make sure that the place was created
				if place.Id != "" {
					_, err := place_io.DeletePlace(place.Id)
					if err != nil {
						fmt.Println(err, " error could not delete history")
					}
				}

				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
				return
			}
			peopleImageObject := people2.People_image{"", peopleId, "", ""}
			peopleImage := people2.PlaceImageHelper{peopleImageObject, filesByteArray}
			_, errr := people_io.CreatePeopleImage(peopleImage)
			/***
			RolBack
			*/
			if errr != nil {
				_, err := history_io.DeleteHistory(historynew.Id)
				if err != nil {
					fmt.Println(err, " error could not delete history")
				}
				//we need to make sure that the place was created
				if place.Id != "" {
					_, err := place_io.DeletePlace(place.Id)
					if err != nil {
						fmt.Println(err, " error could not delete history")
					}
				}
				//Now deleting People History
				_, errx := people_io.DeletePeopleHistory(peopleHistoryNew.Id)
				if errx != nil {
					fmt.Println(err, " error could not delete people History")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/new_stp2/"+peopleId, 301)
				return
			}

			newPeople, errr := people_io.ReadPeople(peopleId)
			if errr != nil {
				fmt.Println(err, " error reading Place Line: 121")
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new person : "+newPeople.Name)
			http.Redirect(w, r, "/admin_user", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
		return
	}
}

func CreatePeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("partner_name")
		surname := r.PostFormValue("surname")
		profession := r.PostFormValue("profession")
		b_date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("b_date"))
		d_date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("d_date"))
		origin := r.PostFormValue("origin")
		if name != "" && surname != "" && profession != "" && origin != "" {
			peopleObject := people2.People{"", name, surname, b_date, d_date, origin, profession}
			people, err := people_io.CreatePeople(peopleObject)
			if err != nil {
				fmt.Println(err, " error creating a new people")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/new", 301)
				return
			}
			if people.Id != "" {
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+people.Name)
				http.Redirect(w, r, "/admin_user/people/new-stp2/"+people.Id, 301)
				return
			}
		}
		fmt.Println("One of the field is missing or newPlace.Id is nil")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/new", 301)
		return
	}
}

func EditPeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peopleId := chi.URLParam(r, "peopleId")
		people, err := people_io.ReadPeople(peopleId)
		if err != nil {
			fmt.Println(err, "error reading people for the following people id: ", peopleId)
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/people", 301)
			return
		}
		type PageDate struct {
			People people2.People
		}
		data := PageDate{people}
		files := []string{
			app.Path + "admin/collection/edit_people.html",
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

func NewPeopleHandler(app *config.Env) http.HandlerFunc {
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
		peoples, err := people_io.ReadPeoples()
		if err != nil {
			fmt.Println(err, "error reading peoples")
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			Peoples       []people2.People
		}
		data := PagePage{backend_error, unknown_error, peoples}
		files := []string{
			app.Path + "admin/people/new_people.html",
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

func PeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peoples := misc.GetPeopleWithStringdate()
		type PagePage struct {
			Peoples []misc.PeopleWithStringdate
		}
		data := PagePage{peoples}
		files := []string{
			app.Path + "admin/people/people.html",
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

func DeletePeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peopleId := chi.URLParam(r, "peopleId")
		people, err := people_io.DeletePeople(peopleId)
		if err != nil {
			fmt.Println(err, "error reading people for the following people id: ", peopleId)
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/people", 301)
			return
		}
		type PageDate struct {
			People people2.People
		}
		data := PageDate{people}
		files := []string{
			app.Path + "admin/collection/edit_people.html",
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
