package collection

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/domain/collection"
	image2 "ostmfe/domain/image"
	"ostmfe/io/collection_io"
	"ostmfe/io/image_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
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
			app.Path + "collection.html",
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
