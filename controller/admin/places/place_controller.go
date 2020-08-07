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
	r.Get("/delete_image/{imageId}/{placeId}", DeleteImageHandler(app))

	r.Post("/create_image", CreatePlaceImage(app))
	r.Post("/create_history", CreateHistoryHandler(app))

	r.Post("/update_pictures", UpdatePictureHandler(app))
	r.Post("/update_details", UpdateDetailsHandler(app))
	r.Post("/update_history", UpdateHistoryHandler(app))

	return r
}

func DeleteImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageId := chi.URLParam(r, "imageId")
		placeId := chi.URLParam(r, "placeId")
		fmt.Println(imageId, " imageId||placeId ", placeId)
		_, err := image_io.ReadImage(imageId)
		if err != nil {
			fmt.Println(err, " error reading image")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		//If we can read the image, than we can delete it.
		_, err = image_io.DeleteImage(imageId)
		if err != nil {
			fmt.Println(err, " error deleting image")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		//if image is deleted now we can delete place image
		placeImage, err := place_io.ReadPlaceImageWithImageId(imageId)
		if err != nil {
			fmt.Println(err, " error reading Place image")
		} else {
			_, err := place_io.DeletePlaceImage(placeImage.Id)
			if err != nil {
				fmt.Println(err, " error reading Place image")
			}
		}

		fmt.Println(" successfully deleted")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an image")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return

	}
}

func UpdateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		placeId := r.PostFormValue("placeId")
		historyId := r.PostFormValue("historyId")
		//checking if the projectHistory exists
		_, err := history_io.ReadHistorie(historyId)
		fmt.Println(historyContent)
		if err != nil {
			fmt.Println(err, " something went wrong! could not read history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		histories := history2.Histories{historyId, misc.ConvertToByteArray(historyContent)}

		_, errr := history_io.UpdateHistorie(histories)
		if errr != nil {
			fmt.Println(err, " something went wrong! could not update history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		place, errx := place_io.ReadPlace(placeId)
		if errx != nil {
			fmt.Println("error reading project")
		}
		fmt.Println(" successfully updated")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: "+place.Title+"  project. ")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func CreateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		PlaceId := r.PostFormValue("PlaceId")
		//checking if there is contents in the variables
		if historyContent != "" && PlaceId != "" {
			history := history2.Histories{"", misc.ConvertToByteArray(historyContent)}

			newHistory, err := history_io.CreateHistorie(history)
			if err != nil {
				fmt.Println(err, " something went wrong! could not create history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
				return
			}
			fmt.Println("HistoryId created successfully ..")
			fmt.Println(" proceeding into creation of a place_history.....")
			placeHistory := place2.PlaceHistory{"", PlaceId, newHistory.Id}
			_, errr := place_io.CreatePlaceHistpory(placeHistory)
			if errr != nil {
				fmt.Println(err, " could not create ProjectHistory")
				fmt.Println("RollBack ...")
				fmt.Println("deleting history ...")
				_, err := history_io.DeleteHistorie(newHistory.Id)
				if err != nil {
					fmt.Println("Error deleting history ...")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
				return
			}
			fmt.Println(" successfully created")
			http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
			return
		}
		fmt.Println("one or more fields missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
		return

	}
}

func UpdateDetailsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		placeTitle := r.PostFormValue("placeTitle")
		placeId := r.PostFormValue("placeId")
		description := r.PostFormValue("description")

		place, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println("error reading place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}

		//TODO latutude and longitude should come from the fields
		if placeTitle != "" && description != "" {
			placeObject := place2.Place{placeId, placeTitle, place.Latitude, place.Longitude, description}
			_, errs := place_io.UpdatePlace(placeObject)
			if errs != nil {
				fmt.Println("error updating place")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
				return
			}

			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: Project Details. ")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		fmt.Println(" error creating project One field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return

	}
}

func UpdatePictureHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		file, _, err := r.FormFile("file")
		PlaceId := r.PostFormValue("PlaceId")
		imageId := r.PostFormValue("imageId")
		placeImageId := r.PostFormValue("placeImageId")
		imageType := r.PostFormValue("imageType")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		place, err := place_io.ReadPlace(PlaceId)
		if err != nil {
			fmt.Println(err, " could not read place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place", 301)
			return
		}
		if imageId != "" && imageType != "" && placeImageId != "" {
			filesArray := []io.Reader{file}
			filesByteArray := misc.CheckFiles(filesArray)
			placeImage := place2.PlaceImage{placeImageId, place.Id, imageId, imageType}

			helper := place2.PlaceImageHelper{placeImage, filesByteArray}
			_, errr := place_io.UpdatePlaceImage(helper)
			if errr != nil {
				fmt.Println(errr, " error creating placeImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following Place : "+place.Title)
			http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
			return
		}
		fmt.Println("one field is empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
		return

	}
}

func CreatePlaceImage(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		placeId := r.PostFormValue("placeId")
		imageType := r.PostFormValue("imageType")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		place, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println(err, " could not read project Line: 113")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}

		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)
		placeImage := place2.PlaceImage{"", placeId, "", imageType}

		helper := place2.PlaceImageHelper{placeImage, filesByteArray}
		_, errr := place_io.CreatePlaceImage(helper)
		if errr != nil {
			fmt.Println(errr, " error creating PlaceImage")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully created image(s) for the following Place  : "+place.Title)
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func DeletePlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		placeId := chi.URLParam(r, "placeId")

		//Reading the project
		place, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println("error reading place")
			fmt.Println(err, " error creating place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/place", 301)
			return
		}
		//Deleting project
		_, err = place_io.DeletePlace(place.Id)
		if err != nil {
			fmt.Println("error deleting place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/place", 301)
			return
		}

		//Reading Project Image
		placeImages, err := place_io.ReadPlaceImageAllOf(placeId)
		if err != nil {
			fmt.Println("error Read Place Image. this project may not have an image")
		} else {
			for _, placeImage := range placeImages {
				_, err = place_io.DeletePlaceImage(placeImage.Id)
				if err != nil {
					fmt.Println("error deleting place Image")
				} else {
					_, err = image_io.DeleteImage(placeImage.ImageId)
					if err != nil {
						fmt.Println("error deleting image")
					}
				}
			}

		}

		//Reading History
		placeHistory, err := place_io.ReadPlaceHistporyOf(placeId)
		if err != nil {
			fmt.Println("error reading placeHistory. this place may not have History")
		} else {
			_, err = history_io.DeleteHistorie(placeHistory.HistoryId)
			if err != nil {
				fmt.Println("error deleting history. this project may not have an history")
			}
			_, err := place_io.DeletePlaceHistpory(placeHistory.Id)
			if err != nil {
				fmt.Println("error deleting projectHistory")
			}
		}

		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: Project Details. ")
		http.Redirect(w, r, "/admin_user/place", 301)
		return
	}
}

func EditPlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		placeId := chi.URLParam(r, "placeId")
		placeDate := GetPlaceEditable(placeId)

		type PageData struct {
			PlaceData PlaceDataEditable
		}
		data := PageData{placeDate}
		files := []string{
			app.Path + "admin/place/new_edit_place.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "base_templates/footer.html",
			app.Path + "admin/template/cards.html",
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
		places, err := place_io.ReadPlaces()
		if err != nil {
			app.InfoLog.Println("error reading Places: ", err)
		}
		type PageData struct {
			Backend_error string
			Unknown_error string
			Places        []place2.Place
		}
		data := PageData{backend_error, unknown_error, places}
		files := []string{
			app.Path + "admin/place/places.html",
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
				fmt.Println(err, " error deleting Place HistoryId")
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
			fmt.Println("error reading HistoryId of the following place", placeId)
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
			fmt.Println("error reading HistoryId of the following place", placeId)
		} else {
			_, err := place_io.DeletePlaceHistpory(placeHistory.Id)
			if err != nil {
				fmt.Println("error delete PlaceHistory of the following place", placeId)
			}
		}
	}
	return true, stringToreturn
}
