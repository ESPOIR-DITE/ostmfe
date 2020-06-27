package places

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	place2 "ostmfe/domain/place"
	"ostmfe/io/place_io"
)

func PlaceHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", PlacesHandler(app))
	r.Get("/new", NewPlacesHandler(app))
	r.Get("/edit", EditPlacesHandler(app))
	r.Post("/create_stp1", CreateStp1Handler(app))
	r.Post("/create_stp2", CreatePlaceStp2Handler(app))
	r.Get("/new_stp2/{placeId}", NewPlaceStp2Handler(app))
	return r
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
			app.Path + "admin/place/new_place.html",
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

func CreatePlaceStp2Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//fileslist := r.Form["file"]
		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		history := r.PostFormValue("history")
		placeId := r.PostFormValue("placeId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}

		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)
		placeHistoryObject := place2.PlaceHistory{"", placeId, history}
		placeHistory, err := place_io.CreatePlaceHistpory(placeHistoryObject)
		if err != nil {
			fmt.Println(err, " error creating a new placeHistory")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/new_stp2/"+placeId, 301)
			return
		}
		placeImageObejct := place2.PlaceImage{"", placeId, "", description}
		placeImageHelper := place2.PlaceImageHelper{placeImageObejct, filesByteArray}
		_, errr := place_io.CreatePlaceImage(placeImageHelper)
		if errr != nil {
			fmt.Println(errr, " error creating projectImage")
			_, err := place_io.DeletePlaceHistpory(placeHistory.Id)
			if err != nil {
				fmt.Println(err, " error deleting Place History")
			}
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/new_stp2/"+placeId, 301)
			return
		}
		place, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println(err, " error reading Place Line: 121")
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+place.Title)
		http.Redirect(w, r, "/admin_user", 301)
		return
	}
}

func NewPlaceStp2Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		placeId := chi.URLParam(r, "placeId")
		place, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println(err, " error reading the Place")
			if app.Session.GetString(r.Context(), "user-read-error") != "" {
				app.Session.Remove(r.Context(), "user-read-error")
			}
			app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/new", 301)
			return
		}
		type PageData struct {
			Place place2.Place
		}
		data := PageData{place}
		files := []string{
			app.Path + "admin/place/image_place.html",
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

func CreateStp1Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		title := r.PostFormValue("title")
		latlng := r.PostFormValue("latlng")
		description := r.PostFormValue("description")
		fmt.Println(title, "<<<title|| latlng>>>>", latlng, "  description>>>", description)
		if title != "" && latlng != "" {
			latitude, longitude := misc.SeparateLatLng(latlng)
			fmt.Println("latitude: ", latitude, "longitude: ", longitude)
			place := place2.Place{"", title, latitude, longitude, description}
			newPlace, err := place_io.CreatePlace(place)
			if err != nil {
				fmt.Println(err, " error when creating a new Place")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/new", 301)
				return
			}
			//Here are trying to make sure that newPlace.Id is not nil.
			if newPlace.Id != "" {
				fmt.Println("successful")
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Place : "+newPlace.Title)
				http.Redirect(w, r, "/admin_user/place/new_stp2/"+newPlace.Id, 301)
				return
			}

		}
		fmt.Println("One of the field is missing or newPlace.Id is nil")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/place/new", 301)
		return
	}
}
