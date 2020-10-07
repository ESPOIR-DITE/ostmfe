package collection

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/collection"
	history2 "ostmfe/domain/history"
	image2 "ostmfe/domain/image"
	"ostmfe/io/collection_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
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
		}
		data := PageData{collections, collectionTypes}
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
