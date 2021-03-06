package collection

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	classroom2 "ostmfe/domain/classroom"
	"ostmfe/domain/collection"
	history2 "ostmfe/domain/history"
	image2 "ostmfe/domain/image"
	classroom3 "ostmfe/io/classroom"
	"ostmfe/io/collection_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/pageData_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/single_collection/{collectionId}", SingleCollectionHandler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func SingleCollectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collectionId := chi.URLParam(r, "collectionId")
		collectionDataHistory := getColectionDataHistory(collectionId)
		if collectionDataHistory.History.Id == "" {
			http.RedirectHandler("/collection", 309)
		}
		type PageData struct {
			CollectionData CollectionDataHistory
		}
		data := PageData{collectionDataHistory}
		files := []string{
			app.Path + "collection/collection_single.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
			app.Path + "base_templates/comments.html",
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
		collectionTypes, err := collection_io.ReadCollectionTyupes()
		if err != nil {
			fmt.Println(err, " error reading collection type")
		}
		collections := getCollectionDatas()

		type PageData struct {
			Collections     []CollectionData
			CollectionTypes []collection.CollectionTypes
			CollectionPage  CollectionPage
			Classrooms      []classroom2.ClassroomHelper
		}
		data := PageData{collections, collectionTypes, getPageData(), getClassroom()}
		files := []string{
			app.Path + "collection/collection.html",
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

//help to get formated classroom
func getClassroom() []classroom2.ClassroomHelper {
	var myClassroomList []classroom2.ClassroomHelper

	classrooms, err := classroom3.ReadClassrooms()
	if err != nil {
		fmt.Println(err, " error reading classroom")
	} else {
		for _, classroom := range classrooms {
			myClassroom := classroom2.ClassroomHelper{classroom.Id, classroom.Name, classroom.Description, misc.ConvertingToString(classroom.Details), misc.ConvertingToString(classroom.Icon)}
			myClassroomList = append(myClassroomList, myClassroom)
		}
	}
	return myClassroomList
}

type CollectionData struct {
	Collection collection.Collection
	Image      image2.Images
}

func getCollectionDatas() []CollectionData {
	var CollectionList []CollectionData
	var images image2.Images
	collections, err := collection_io.ReadCollections()
	if err != nil {
		fmt.Println(err, " error reading collections")
		return CollectionList
	}
	for _, collection := range collections {
		collectionImage, err := collection_io.ReadCollectionImgWithCollectionId(collection.Id)
		if err != nil {
			fmt.Println(err, " error reading collectionImage")
		} else {
			images, err = image_io.ReadImage(collectionImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading Image")
			}
		}
		collectionDataObject := CollectionData{collection, images}
		CollectionList = append(CollectionList, collectionDataObject)
	}
	return CollectionList
}

type CollectionDataHistory struct {
	Collection   collection.Collection
	ProfileImage image2.Images
	Images       []image2.Images
	History      history2.HistoriesHelper
}

func getColectionDataHistory(collectionId string) CollectionDataHistory {
	var collectionDataHistory CollectionDataHistory
	var profileImage image2.Images
	var images []image2.Images
	var histories history2.HistoriesHelper
	//check collection
	collection, err := collection_io.ReadCollection(collectionId)
	if err != nil {
		fmt.Println(err, " error reading collection")
		return collectionDataHistory
	}

	//Images
	collectionImages, err := collection_io.ReadCollectionImgsWithCollectionId(collectionId)
	if err != nil {
		fmt.Println(err, " error reading collectionImage")
	} else {
		for _, collectionImage := range collectionImages {
			if collectionImage.Description == "1" || collectionImage.Description == "profile" {
				profileImage, err = image_io.ReadImage(collectionImage.ImageId)
				if err != nil {
					fmt.Println(err, " error reading profile Image")
				}
			}
			image, err := image_io.ReadImage(collectionImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading  Image")
			}
			images = append(images, image)

		}
	}
	////History
	collectionHistorie, err := collection_io.ReadCollectionHistoryWithCollectionId(collectionId)
	if err != nil {
		fmt.Println(err, " error reading Collection History")
	} else {
		history, err := history_io.ReadHistorie(collectionHistorie.HistoryId)
		if err != nil {
			fmt.Println(err, " error reading Historie")
		}
		histories = history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}
	}

	collectionDataHistory = CollectionDataHistory{collection, profileImage, images, histories}
	return collectionDataHistory
}

type CollectionPage struct {
	Banner            string
	SahoContent       string
	ResourceContent   string
	CollectionContent string
}

func getPageData() CollectionPage {
	var banner string
	var sahoContent string
	var resourceContent string
	var collectionContent string

	page, err := pageData_io.ReadPageDataWIthName("collection-page")
	if err != nil {
		fmt.Println(err, " error reading page")
	} else {
		pageDateSectionObject, err := pageData_io.ReadPageSectionAllOf(page.Id)
		if err != nil {
			fmt.Println(err, " error reading page")
		}
		for _, pageDateSection := range pageDateSectionObject {
			pageSection, err := pageData_io.ReadSection(pageDateSection.SectionId)
			if err != nil {
				fmt.Println(err, " error reading page")
			} else {
				if pageSection.SectionName == "sahoContent" {
					fmt.Println(" sahoContent", pageSection)
					sahoContent = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "resourceContent" {
					fmt.Println(" resourceContent", pageSection)
					resourceContent = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "collectionContent" {
					fmt.Println(" collectionContent", pageSection)
					collectionContent = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "banner" {
					fmt.Println(" banner", pageSection)
					banner = misc.ConvertingToString(pageDateSection.Content)
				}
			}
		}
	}
	return CollectionPage{sahoContent, resourceContent, collectionContent, banner}
}
