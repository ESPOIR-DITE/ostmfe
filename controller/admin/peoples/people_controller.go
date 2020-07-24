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
	r.Get("/new_stp2/{peopleId}", NewPeoplestp2Handler(app))
	r.Get("/edit/{peopleId}", EditPeopleHandler(app))
	r.Get("/delete/{peopleId}", DeletePeopleHandler(app))
	r.Post("/create_stp1", CreatePeopleHandler(app))
	r.Post("/create_stp2", CreatePeopleStp2Handler(app))
	r.Post("/people_category/create", CreatePeopleCategoryHandler(app))
	r.Post("/update_image", UpdatePeopleImageHandler(app))
	r.Post("/update/details", UpdatePeopleDetailHandler(app))
	r.Post("/update/history", UpdatePeopleHistoryHandler(app))
	r.Post("/create_image", CreatePeopleImageHandler(app))
	return r
}

/****
This method is requested when a people was created without an image now this method will help to create one.
*/
func CreatePeopleImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		peopleId := r.PostFormValue("peopleId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		if peopleId != "" {

			peopleImageObject := people2.People_image{"", peopleId, "", ""}
			peopleImageHelper := people2.PeopleImageHelper{peopleImageObject, filesByteArray, ""}

			_, errx := people_io.CreatePeopleImage(peopleImageHelper)
			if errx != nil {
				fmt.Println(errx, " error creating PeopleImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
		return
	}

}

func UpdatePeopleHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		historyId := r.PostFormValue("historyId")
		peopleId := r.PostFormValue("projectId")
		myArea := r.PostFormValue("myArea")
		if myArea != "" && historyId != "" && peopleId != "" {
			historieObject := history2.Histories{historyId, misc.ConvertToByteArray(myArea)}
			_, err := history_io.UpdateHistorie(historieObject)
			if err != nil {
				fmt.Println(err, "updating people histories")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People History of : ")
			http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
		return
	}
}

func UpdatePeopleDetailHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("name")
		peopleId := r.PostFormValue("peopleId")
		surname := r.PostFormValue("surname")
		b_date, _ := time.Parse(misc.YYYMMDDTIME_FORMAT, r.PostFormValue("b_date"))
		d_date, _ := time.Parse(misc.YYYMMDDTIME_FORMAT, r.PostFormValue("d_date"))
		profession := r.PostFormValue("profession")
		origin := r.PostFormValue("origin")
		if name != "" && peopleId != "" && surname != "" && profession != "" && origin != "" {
			//TODO need to learn how to check if a data time if nil or empty....
			peopleObejct := people2.People{peopleId, name, surname, b_date, d_date, origin, profession}
			people, err := people_io.CreatePeople(peopleObejct)
			if err != nil {
				fmt.Println(err, "updating people details")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Details of : "+people.Name)
			http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
		return
	}
}

func UpdatePeopleImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		imageId := r.PostFormValue("imageId")
		peopleId := r.PostFormValue("peopleId")
		peopleImageId := r.PostFormValue("peopleImageId")
		imageType := r.PostFormValue("imageType")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file}
		filesByteArray := misc.CheckFiles(filesArray)

		if peopleId != "" && imageId != "" && imageType != "" && peopleImageId != "" {
			peopleImage := people2.People_image{peopleImageId, peopleId, imageId, imageType}
			peopleImageHeper := people2.PeopleImageHelper{peopleImage, filesByteArray, imageId}
			_, err := people_io.UpdatePeopleImage(peopleImageHeper)
			if err != nil {
				fmt.Println(err, " error updating peopleImageHeper")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
		return
	}
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
				http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new People Type : "+peopleCategory.Category)
			http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
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

		peoples := misc.GetPeopleWithStringdate()
		type PageData struct {
			People        people2.People
			Backend_error string
			Unknown_error string
			Peoples       []misc.PeopleWithStringdate
		}
		data := PageData{people, backend_error, unknown_error, peoples}
		files := []string{
			app.Path + "admin/people/people_step2.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
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
		var historynew history2.Histories
		r.ParseForm()
		//fileslist := r.Form["file"]

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		history := r.PostFormValue("history")
		peopleId := r.PostFormValue("peopleId")
		//latlng := r.PostFormValue("latlng")
		//title := r.PostFormValue("title")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		if history != "" && peopleId != "" {
			fmt.Println(history, " History || PeopleId >>>>", peopleId)
			historyObejct := history2.Histories{"", misc.ConvertToByteArray(history)}
			historynew, err = history_io.CreateHistorie(historyObejct)
			if err != nil {
				fmt.Println(err, " error creating a new history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
				return
			}
			// I am checking if the this person is relative to a particular place
			//if latlng != "" {
			//	latitude, longitude := misc.SeparateLatLng(latlng)
			//	placeObject := place2.Place{"", title, latitude, longitude, ""}
			//	place, err = place_io.CreatePlace(placeObject)
			//	if err != nil {
			//		fmt.Println(err, " error creating a new Place")
			//		_, err := history_io.DeleteHistory(historynew.Id)
			//		if err != nil {
			//			fmt.Println(err, " error could not delete history")
			//		}
			//		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			//			app.Session.Remove(r.Context(), "user-create-error")
			//		}
			//		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			//		http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
			//		return
			//	}
			//}

			peopleHistory := people2.PeopleHistory{"", peopleId, historynew.Id}

			peopleHistoryNew, err := people_io.CreatePeopleHistory(peopleHistory)
			if err != nil {
				fmt.Println(err, " error creating a new people history")
				_, err := history_io.DeleteHistorie(historynew.Id)
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
			peopleImage := people2.PeopleImageHelper{peopleImageObject, filesByteArray, ""}
			_, errr := people_io.CreatePeopleImage(peopleImage)
			/***
			RolBack
			*/
			if errr != nil {
				if historynew.Id != "" {
					fmt.Println(" error could not create people Image....")
					_, err := history_io.DeleteHistorie(historynew.Id)
					if err != nil {
						fmt.Println(err, " error could not delete history....")
					} else {
						fmt.Println(err, " Deleted")
					}
				}

				//we need to make sure that the place was created
				if place.Id != "" {
					fmt.Println(err, " Deleting Place of this event....")
					_, err := place_io.DeletePlace(place.Id)
					if err != nil {
						fmt.Println(err, " error could not delete history")
					} else {
						fmt.Println(err, " Deleted")
					}
				}
				//Now deleting People History
				if peopleHistoryNew.Id != "" {
					_, errx := people_io.DeletePeopleHistory(peopleHistoryNew.Id)
					if errx != nil {
						fmt.Println(err, " error could not delete people History")
					} else {
						fmt.Println(err, " Deleted")
					}
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
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
		placeId := r.PostFormValue("placeId")
		name := r.PostFormValue("name")
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
			//Creating people place
			if placeId != "" {
				peoplePlaceObject := people2.PeoplePlace{"", placeId, people.Id}
				_, err := people_io.CreatePeoplePlace(peoplePlaceObject)
				if err != nil {
					fmt.Println("error creating peoplePlace")
				}
			}
			if people.Id != "" {
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+people.Name)
				http.Redirect(w, r, "/admin_user/people/new_stp2/"+people.Id, 301)
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
		peopleEditable := GetPeopleEditable(people.Id)
		type PageDate struct {
			PeopleDetails people2.People
			People        PeopleEditable
		}
		data := PageDate{people, peopleEditable}
		files := []string{
			app.Path + "admin/people/new_edite_people.html",
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

		peoples := misc.GetPeopleWithStringdate()
		places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println("error reading Places")
		}
		type PagePage struct {
			Places        []place2.Place
			Backend_error string
			Unknown_error string
			Peoples       []misc.PeopleWithStringdate
		}
		data := PagePage{places, backend_error, unknown_error, peoples}
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
