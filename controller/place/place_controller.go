package place

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/place"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/place_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/{placeId}", SinglePlaceHanler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func SinglePlaceHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		placeId := chi.URLParam(r, "placeId")
		places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println(err, "Error reading places")
		}
		type PageData struct {
			Places    []place.Place
			PlaceData PlaceSingleData
		}
		data := PageData{places, getPlaceSingleData(placeId)}
		files := []string{
			app.Path + "place/place_single.html",
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

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println(err, "Error reading places")
		}
		type PageData struct {
			Places []place.Place
		}
		data := PageData{places}
		files := []string{
			app.Path + "place/places_page.html",
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

type PlaceSingleData struct {
	Place        place.Place
	Image        []image3.ImagesHelper
	History      history.HistoriesHelper
	ProfileImage image3.ImagesHelper
}

func getPlaceSingleData(placeId string) PlaceSingleData {
	var placeSingleData PlaceSingleData
	var profileImage image3.ImagesHelper
	var image []image3.ImagesHelper
	var histor history.HistoriesHelper

	place, err := place_io.ReadPlace(placeId)
	if err != nil {
		fmt.Println(err, "Error reading place")
		return placeSingleData
	}

	//Images
	placeImages, err := place_io.ReadPlaceImageAllOf(placeId)
	if err != nil {
		fmt.Println(err, "Error reading imagePlaces")
	} else {
		for _, placeImage := range placeImages {
			if placeImage.Description == "1" || placeImage.Description == "profile" {
				img, err := image_io.ReadImage(placeImage.ImageId)
				if err != nil {
					fmt.Println(err, "Error reading image")
				}
				profileImage = image3.ImagesHelper{img.Id, misc.ConvertingToString(img.Image), placeImage.Id}
			}
			img, err := image_io.ReadImage(placeImage.ImageId)
			if err != nil {
				fmt.Println(err, "Error reading image")
			}
			imageObject := image3.ImagesHelper{img.Id, misc.ConvertingToString(img.Image), placeImage.Id}
			image = append(image, imageObject)
		}

	}
	//History
	placeHistory, err := place_io.ReadPlaceHistporyOf(placeId)
	if err != nil {
		fmt.Println(err, "Error reading place history. This place may not have history")
	} else {
		historyObejct, err := history_io.ReadHistorie(placeHistory.HistoryId)
		if err != nil {
			fmt.Println(err, "Error reading image")
		}
		histor = history.HistoriesHelper{historyObejct.Id, misc.ConvertingToString(historyObejct.History)}
	}
	placeSingleData = PlaceSingleData{place, image, histor, profileImage}
	return placeSingleData
}
