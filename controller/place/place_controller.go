package place

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/domain/event"
	"ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/people"
	"ostmfe/domain/place"
	"ostmfe/io/comment_io"
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
		count, err := comment_io.CountCommentPlace(placeId)
		if err != nil {
			fmt.Println(err, "Error counting place comments")
		}
		pageFlow, err := place_io.ReadAllPlacePageFlowByPlaceId(placeId)
		if err != nil {
			fmt.Println(err, "Error counting place Page FLow")
		}
		type PageData struct {
			Places        []place.Place
			PlaceData     PlaceSingleData
			GalleryString []image3.GaleryHelper
			CommentNumber int64
			Comments      []comment.CommentStack
			PageFlows     []place.PlacePageFlow
		}
		data := PageData{places, getPlaceSingleData(placeId), getPlaceGallery(placeId), count, GetPlaceComments(placeId), pageFlow}
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
			if placeImage.Id != "" {
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

func getPlaceGallery(placeId string) []image3.GaleryHelper {
	var picture []image3.GaleryHelper
	placeGallerys, err := place_io.ReadAllByPlaceGallery(placeId)
	if err != nil {
		fmt.Println(err, " error peopleGalleries.")
	} else {
		for _, placeGallery := range placeGallerys {
			gallery, err := image_io.ReadGallery(placeGallery.GalleryId)
			if err != nil {
				fmt.Println(err, " error gallery")
			} else {
				picture = append(picture, image3.GaleryHelper{gallery.Id, misc.ConvertingToString(gallery.Image), gallery.Description})
			}
		}
	}
	return picture
}

type PlaceAggregatedDate struct {
	Place         place.Place
	Events        []event.Event
	Gallery       []string
	People        []PeopleDataPlace
	Image         string
	PlaceCategory string
	Category      place.PlaceCategory
}

//Get people data
type PeopleDataPlace struct {
	People        people.People
	Image         image3.Images
	Profession    []people.Profession
	History       history.History
	PlaceCategory []people.Category
}

func getPlaceAggregatedData() []PlaceAggregatedDate {
	var placeAggregated []PlaceAggregatedDate
	var events []event.Event
	var gallery []string
	var peoples []PeopleDataPlace
	var image string
	var PlaceCategory string
	var category place.PlaceCategory

	places, err := place_io.ReadPlaces()
	if err != nil {
		fmt.Println(err, " error places")
		return placeAggregated
	}
	for _, place := range places {
		PlaceCategory = getPlaceCategory(place.Id)
		//Event Place
		eventPlace, err := event_io.ReadEventFindByPlaceId(place.Id)
		if err != nil {
			fmt.Println(err, " error Event place")
		} else {
			events = getEvents(eventPlace)
		}
		//Category
		category, err = misc.GetPlaceCategory(place.Id)
		if err != nil {
			fmt.Println(err, " error getting place category")
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
		placeAggregated = append(placeAggregated, PlaceAggregatedDate{place,
			events, gallery, peoples, image, PlaceCategory, category})
	}
	return placeAggregated
}

func getPlaceCategory(placeId string) string {
	var placeCategoryObject string
	PlaceType, err := place_io.ReadPlaceTypeOf(placeId)
	if err != nil {
		fmt.Println(err, " error reading Place Type")
		return placeCategoryObject
	}
	placeCategory, err := place_io.ReadPlaceCategory(PlaceType.PlaceCategoryId)
	if err != nil {
		fmt.Println(err, " error reading Place category")
	} else {
		return placeCategory.Name
	}
	return placeCategoryObject
}

//this method works as a bridge to allow extract
//id list from eventPlace ids to form a list of EventIds.
func getEvents(eventPlaces []event.EventPlace) []event.Event {
	var eventIds []string
	for _, eventPlace := range eventPlaces {
		eventIds = append(eventIds, eventPlace.EventId)
	}
	return misc.GetEventListOfEventIdList(eventIds)
}

//GET A PEOPLE WITH ALL HIS DATA
func GetPeopleDataPlace(peopleId string) PeopleDataPlace {
	var peopleData PeopleDataPlace
	var professions []people.Profession
	var categoryList []people.Category
	var imageList image3.Images
	var history history.History
	people, err := people_io.ReadPeople(peopleId)
	if err != nil {
		fmt.Println(err, "error reading peoples")
		return peopleData
	} else {
		//getting the Images od this person
		peopleImages, err := people_io.ReadPeopleImagewithPeopleId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading peopleImage for peopleId: ", people.Id)
		} else {
			for _, peopleImage := range peopleImages {
				images, err := image_io.ReadImage(peopleImage.ImageId)
				if err != nil {
					fmt.Println(err, "error reading peopleImage for peopleImageId: ", peopleImage.Id)
				} else {
					imageList = images
				}
			}
		}

		//Getting People proffesions
		peopleProfession, err := people_io.ReadPeopleProfessionWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading peopleProfession for peopleId: ", people.Id)
		} else {
			for _, peopleProfession := range peopleProfession {
				profession, err := people_io.ReadProfession(peopleProfession.Profession)
				if err != nil {
					fmt.Println(err, "error reading Profession for peopleId: ", people.Id)
				} else {
					professions = append(professions, profession)
				}
			}
		}

		//Getting People history
		peopleHistory, err := people_io.ReadPeopleHistoryWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading people history for peopleId: ", people.Id)
		} else {
			history, err = history_io.ReadHistory(peopleHistory.HistoryId)
			if err != nil {
				fmt.Println(err, "error reading history: ", people.Id)
			}
		}

		//Getting People category
		peoplecategories, err := people_io.ReadPeopleCategoryWithCategoryId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading people category for peopleId: ", people.Id)
		} else {
			for _, peoplecategory := range peoplecategories {
				category, err := people_io.ReadCategory(peoplecategory.Id)
				if err != nil {
					fmt.Println(err, "error reading category for peopleId: ", peoplecategory.Id)
				} else {
					categoryList = append(categoryList, category)
				}
			}
		}
		peopleDataObject := PeopleDataPlace{people, imageList, professions, history, categoryList}
		return peopleDataObject
	}
	return peopleData
}

func getPeopleWithPeoplePlace(peoplePlaces []people.PeoplePlace) []PeopleDataPlace {
	var peoples []PeopleDataPlace
	for _, peoplePlace := range peoplePlaces {
		peoples = append(peoples, GetPeopleDataPlace(peoplePlace.PeopleId))
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

func GetPlaceComments(placeId string) []comment.CommentStack {
	var parentCommentStack []comment.CommentStack
	var subCommentStack []comment.CommentHelper

	for _, commentObject := range getComments(placeId) {
		myComment, err := comment_io.ReadComment(commentObject.Id)
		if err != nil {
			fmt.Println("error reading Comment")
		}
		if myComment.ParentCommentId != "" {
			subCommentStack = getSubComment(commentObject.Id)
		}
		parentCommentStack = append(parentCommentStack, comment.CommentStack{commentObject, subCommentStack})
	}

	fmt.Println("parentStack ", parentCommentStack)

	return parentCommentStack
}

func getSubComment(parentComment string) []comment.CommentHelper {
	var myComments []comment.CommentHelper
	subComments, err := comment_io.ReadCommentWithParentId(parentComment)
	if err != nil {
		return myComments
	}
	for _, eventComment := range subComments {
		if eventComment.ParentCommentId == parentComment && eventComment.Comment != nil {
			commentHelper := comment.CommentHelper{eventComment.Id, eventComment.Email, eventComment.Name, misc.FormatDateMonth(eventComment.Date), misc.ConvertingToString(eventComment.Comment), eventComment.ParentCommentId, eventComment.Stat}
			myComments = append(myComments, commentHelper)
		}
	}
	return myComments
}

//This method returns a list of either parent or subComment
func getComments(placeId string) []comment.CommentHelper {
	var myCommentObject []comment.CommentHelper
	placeComments, err := comment_io.ReadAllByPlaceId(placeId)
	if err != nil {
		fmt.Println("error reading place Comment")
		return myCommentObject
	}
	for _, eventComment := range placeComments {
		myComment, err := comment_io.ReadComment(eventComment.CommentId)
		if err != nil {
			fmt.Println("error reading Comment")
		} else {
			commentHelper := comment.CommentHelper{myComment.Id, myComment.Email, myComment.Name, misc.FormatDateMonth(myComment.Date), misc.ConvertingToString(myComment.Comment), myComment.ParentCommentId, myComment.Stat}
			myCommentObject = append(myCommentObject, commentHelper)
		}
	}
	return myCommentObject
}
