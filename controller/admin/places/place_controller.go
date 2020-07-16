package places

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	image2 "ostmfe/domain/image"
	place2 "ostmfe/domain/place"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/place_io"
)

func PlaceHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", PlacesHandler(app))
	r.Get("/new", NewPlacesHandler(app))
	r.Get("/edit/{placeId}", EditPlacesHandler(app))
	r.Get("/delete/{placeId}", DeletePlacesHandler(app))
	r.Post("/create_stp1", CreateStp1Handler(app))
	r.Post("/create_stp2", CreatePlaceStp2Handler(app))
	r.Get("/new_stp2/{placeId}", NewPlaceStp2Handler(app))
	return r
}

func DeletePlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		placeId := chi.URLParam(r, "placeId")
		fmt.Println(placeId)
		//result,report := deletePlaceData(placeId)
	}
}

func EditPlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		placeId := chi.URLParam(r, "placeId")
		var historyHelper history2.HistoryHelper
		placeDate := getPlaceData(placeId)

		history, err := history_io.ReadHistory(placeDate.History.Id)
		if err != nil {
			fmt.Println("error reading History")
		} else {
			historyHelper = history2.HistoryHelper{history.Id, history.Title, history.Description, misc.ConvertingToString(history.Content), history.Date}
		}

		type PageData struct {
			PlaceData PlaceData
			History   history2.HistoryHelper
		}
		data := PageData{placeDate, historyHelper}
		files := []string{
			app.Path + "admin/place/edite_place.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "base_templates/footer.html",
			app.Path + "admin/template/cards.html",
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
		places, err := place_io.ReadPlaces()
		if err != nil {
			app.InfoLog.Println("error reading Places: ", err)
		}
		type PageData struct {
			Places []place2.Place
		}
		data := PageData{places}
		files := []string{
			app.Path + "admin/place/places.html",
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

//Agreggeted Place Data
type PlaceData struct {
	Place   place2.Place
	Images  []image2.Images
	History history2.History
}

func getPlaceData(placeId string) PlaceData {
	var placeDate PlaceData
	var images []image2.Images
	var history history2.History
	place, err := place_io.ReadPlace(placeId)
	if err != nil {
		fmt.Println("error reading place data")
		return placeDate
	}
	placeImages, err := place_io.ReadPlaceImageAllOf(placeId)
	if err != nil {
		fmt.Println("error reading placeImages")
	} else {
		for _, placeImage := range placeImages {
			image, err := image_io.ReadImage(placeImage.ImageId)
			if err != nil {
				fmt.Println("error reading Images of the following placeId", placeImage.PlaceId)
			} else {
				images = append(images, image)
			}
		}
	}
	placeHistory, err := place_io.ReadPlaceHistporyOf(placeId)
	if err != nil {
		fmt.Println("error reading PlaceHistory")
	} else {
		history, err = history_io.ReadHistory(placeHistory.HistoryId)
		if err != nil {
			fmt.Println("error reading History of the following place", placeId)
		}
	}
	placeDate = PlaceData{place, images, history}
	return placeDate
}

//Todo The following method need to be double checked
func deletePlaceData(placeId string) (bool, string) {
	var stringToreturn string
	_, err := place_io.ReadPlace(placeId)
	if err != nil {
		fmt.Println("error reading place data")
		return false, "Error Place can not be find"
	} else {
		_, err := place_io.DeletePlace(placeId)
		if err != nil {
			fmt.Println("error reading place data")
			return false, "Error Place can not be Deleted"
		}
	}
	placeImages, err := place_io.ReadPlaceImageAllOf(placeId)
	if err != nil {
		fmt.Println("error reading placeImages")
	} else {
		for _, placeImage := range placeImages {
			_, err := image_io.DeleteImage(placeImage.ImageId)
			if err != nil {
				fmt.Println("error Deleting Image of the following placeId", placeImage.PlaceId)
			} else {
				_, err := place_io.DeletePlaceImage(placeImage.Id)
				if err != nil {
					fmt.Println("error Deleting PlaceImage of the following placeId", placeImage.PlaceId)
				}
			}
		}
	}
	placeHistory, err := place_io.ReadPlaceHistporyOf(placeId)
	if err != nil {
		fmt.Println("error reading PlaceHistory")
	} else {
		_, err := history_io.DeleteHistory(placeHistory.HistoryId)
		if err != nil {
			fmt.Println("error reading History of the following place", placeId)
		} else {
			_, err := place_io.DeletePlaceHistpory(placeHistory.Id)
			if err != nil {
				fmt.Println("error delete PlaceHistory of the following place", placeId)
			}
		}
	}
	return true, stringToreturn
}
