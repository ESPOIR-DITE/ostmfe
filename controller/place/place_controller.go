package place

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/event"
	"ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/people"
	"ostmfe/domain/place"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/pageData_io"
	"ostmfe/io/people_io"
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
			Places        []place.Place
			PlaceData     PlaceSingleData
			GalleryString []string
		}
		data := PageData{places, getPlaceSingleData(placeId), getPlaceGallery(placeId)}
		files := []string{
			app.Path + "place/place_single.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/comments.html",
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
func getPlaceImage(PlaceId string) string {
	placeImage, err := place_io.ReadPlaceImageByPlaceId(PlaceId)
	if err != nil {
		fmt.Println(err, "Error reading places Image")
		return ""
	}
	image, err := image_io.ReadImage(placeImage.ImageId)
	if err != nil {
		fmt.Println(err, "Error reading Image")
		return ""
	}
	return misc.ConvertingToString(image.Image)
}
func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println(err, "Error reading places")
		}
		var bannerImage string
		pageBanner, err := pageData_io.ReadPageBannerWIthPageName("place-page")
		if err != nil {
			fmt.Println(err, " There is an error when reading place pageBanner")
		} else {
			bannerImage = misc.GetBannerImage(pageBanner.BannerId)
		}
		type PageData struct {
			Places      []place.Place
			PlaceBanner string
			GetImage    func(placeId string) string
			PlaceData   []PlaceAggregatedDate
		}

		data := PageData{places,
			bannerImage,
			func(placeId string) string {
				return getPlaceImage(placeId)
			},
			getPlaceAggregatedData(),
		}
		files := []string{
			//app.Path + "place/places_page.html",
			//app.Path + "place/places_page_mapBox.html",
			app.Path + "place/mapBoxWithPics.html",
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

func getPlaceGallery(placeId string) []string {
	var picture []string
	placeGallerys, err := place_io.ReadAllByPlaceGallery(placeId)
	if err != nil {
		fmt.Println(err, " error peopleGalleries.")
	} else {
		for _, placeGallery := range placeGallerys {
			gallery, err := image_io.ReadGallery(placeGallery.GalleryId)
			if err != nil {
				fmt.Println(err, " error gallery")
			} else {
				picture = append(picture, misc.ConvertingToString(gallery.Image))
			}
		}
	}
	return picture
}

type PlaceAggregatedDate struct {
	Place   place.Place
	Event   event.Event
	Gallery []string
	People  []misc.PeopleData
	Image   string
}

func getPlaceAggregatedData() []PlaceAggregatedDate {
	var placeAggregated []PlaceAggregatedDate
	var events event.Event
	var gallery []string
	var peoples []misc.PeopleData
	var image string

	places, err := place_io.ReadPlaces()
	if err != nil {
		fmt.Println(err, " error places")
		return placeAggregated
	}
	for _, place := range places {
		//Event Place
		eventPlace, err := event_io.ReadEventPlaceOf(place.Id)
		if err != nil {
			fmt.Println(err, " error Event place")
		} else {
			events = getEvent(eventPlace.EventId)
		}

		//Gallery
		PlaceGallery, err := place_io.ReadAllByPlaceGallery(place.Id)
		if err != nil {
			fmt.Println(err, " error Gallery place")
		} else {
			gallery = getGallery(PlaceGallery)
		}

		//People
		peoplePlaceObject, err := people_io.ReadPeoplePlaceAllByPlaceId(place.Id)
		if err != nil {
			fmt.Println(err, " error people place")
		} else {
			peoples = getPeopleWithPeoplePlace(peoplePlaceObject)
		}
		//Image
		placeImage, err := place_io.ReadPlaceImageByPlaceId(place.Id)
		if err != nil {
			fmt.Println(err, " error image place")
		} else {
			image = getImage(placeImage.ImageId).Id
		}
		placeAggregated = append(placeAggregated, PlaceAggregatedDate{place, events, gallery, peoples, image})
	}
	return placeAggregated
}

func getPeopleWithPeoplePlace(peoplePlaces []people.PeoplePlace) []misc.PeopleData {
	var peoples []misc.PeopleData
	for _, peoplePlace := range peoplePlaces {
		peoples = append(peoples, misc.GetPeopleData(peoplePlace.PeopleId))
	}
	return peoples
}

func getEvent(eventId string) event.Event {
	var eventObject event.Event
	newEventObject, err := event_io.ReadEvent(eventId)
	if err != nil {
		fmt.Println(err, " error places")
		return eventObject
	}
	return newEventObject
}
func getGallery(placeGalleries []place.PlaceGallery) []string {
	var imageHelper []string

	for _, placeGallerie := range placeGalleries {
		imageobject, err := image_io.ReadGallery(placeGallerie.GalleryId)
		if err != nil {
			fmt.Println(err, " error gallery")
		} else {
			imageHelper = append(imageHelper, misc.ConvertingToString(imageobject.Image))
		}
	}

	return imageHelper
}
func getPeople(pepleId string) people.People {
	var people people.People
	peopleObject, err := people_io.ReadPeople(pepleId)
	if err != nil {
		fmt.Println(err, " error gallery")
		return people
	}
	return peopleObject
}
func getImage(imageId string) image3.Images {
	var image image3.Images
	imageObject, err := image_io.ReadImage(imageId)
	if err != nil {
		fmt.Println(err, " error gallery")
		return image
	}
	return imageObject
}
