package collection

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/collection"
	"ostmfe/io/collection_io"
)

func CollectionHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", CollectionHandler(app))
	r.Get("/new", NewCollectionHandler(app))
	r.Get("/collection_type/new", NewCollectionTypeHandler(app))
	r.Get("/new_stp/{collectionId}", NewCollection2Handler(app))
	r.Get("/edit", EditCollectionHandler(app))
	r.Post("/create_stp1", CreateCollection1(app))
	r.Post("/create_stp2", CreateCollection2(app))
	r.Post("/collection_type/create", CreateCollectionType(app))
	return r
}
func CreateCollectionType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		type_name := r.PostFormValue("type_name")
		description := r.PostFormValue("description")

		if description != "" && type_name != "" {
			collectionTypeObject := collection.CollectionTypes{"", type_name, description}
			collectionType, err := collection_io.CreateCollectionTyupe(collectionTypeObject)
			if err != nil {
				fmt.Println(err, " error creating collectionType")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/collection/collection_type/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Collection Type : "+collectionType.Name)
			http.Redirect(w, r, "/admin_user/collection/collection_type/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/collection/collection_type/new", 301)
		return
	}
}

func NewCollectionTypeHandler(app *config.Env) http.HandlerFunc {
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
		collectionTypes, err := collection_io.ReadCollectionTyupes()
		if err != nil {
			fmt.Println(err, " error Creating collection")
		}
		type PagePage struct {
			Backend_error  string
			Unknown_error  string
			CollectionType []collection.CollectionTypes
		}
		data := PagePage{backend_error, unknown_error, collectionTypes}
		files := []string{
			app.Path + "admin/Collection/collectionType_tables.html",
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

func CreateCollection2(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//fileslist := r.Form["file"]

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		collectionId := r.PostFormValue("collectionId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}

		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		fmt.Println(collectionId, " CollectionId")
		collectionObject, err := collection_io.ReadCollection(collectionId)
		if err != nil {
			fmt.Println(err, " error reading Collection")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection/new_stp2/"+collectionId, 301)
			return
		}
		collectionImage := collection.Collection_image{"", collectionObject.Id, collectionId, description}
		CollectionImageHelper := collection.CollectionImageHelper{collectionImage, filesByteArray}

		_, errr := collection_io.CreateCollectionImg(CollectionImageHelper)
		if errr != nil {
			fmt.Println(errr, " error creating CollectionImage")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection/new_stp2/"+collectionId, 301)
			return
		}

		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+collectionObject.Name)
		http.Redirect(w, r, "/admin_user", 301)
		return
	}
}

func NewCollection2Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collectionId := chi.URLParam(r, "collectionId")
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
		collectionObject, err := collection_io.ReadCollection(collectionId)
		if err != nil {
			fmt.Println(err, " error reading collection")
		}

		type PagePage struct {
			Backend_error string
			Unknown_error string
			Collection    collection.Collection
		}
		data := PagePage{backend_error, unknown_error, collectionObject}
		files := []string{
			app.Path + "admin/collection/image_collection.html",
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

func CreateCollection1(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		collection_name := r.PostFormValue("collection_name")
		collectionId := r.PostFormValue("collectionId")
		brief := r.PostFormValue("brief")
		history := r.PostFormValue("history")
		if history != "" && collection_name != "" && collectionId != "" && brief != "" {

			collectionObject := collection.Collection{"", collection_name, brief, misc.ConvertToByteArray(history)}
			newCollection, err := collection_io.CreateCollection(collectionObject)
			if err != nil {
				fmt.Println(err, " error reading the collection")
				if app.Session.GetString(r.Context(), "user-read-error") != "" {
					app.Session.Remove(r.Context(), "user-read-error")
				}
				app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/collection/new", 301)
				return
			}

			collectionTypeObject := collection.Collection_type{"", newCollection.Id, collectionId}
			_, errr := collection_io.CreateCollection_Type(collectionTypeObject)
			if errr != nil {
				_, err := collection_io.DeleteCollection(newCollection.Id)
				if err != nil {
					fmt.Println(err, " error could not delete collection")
				}
				//fmt.Println(err, " error reading the collection")
				if app.Session.GetString(r.Context(), "user-read-error") != "" {
					app.Session.Remove(r.Context(), "user-read-error")
				}
				app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/collection/new", 301)
				return
			}
			//If creation successful
			if newCollection.Id != "" {
				http.Redirect(w, r, "/admin_user/collection/new_stp/"+newCollection.Id, 301)
				return
			}

		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/collection/new", 301)
		return
	}
}

func EditCollectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/edit_collection.html",
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

func NewCollectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var collectionType []collection.CollectionTypes
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
		collectionType, err := collection_io.ReadCollectionTyupes()
		if err != nil {
			fmt.Println(err, " error reading Collections")
		}
		type PagePage struct {
			Backend_error  string
			Unknown_error  string
			CollectionType []collection.CollectionTypes
		}
		data := PagePage{backend_error, unknown_error, collectionType}
		files := []string{
			app.Path + "admin/collection/new_collection.html",
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

func CollectionHandler(app *config.Env) http.HandlerFunc {
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
		collections := misc.GetCollectionBridge()

		type PagePage struct {
			Backend_error string
			Unknown_error string
			Collections   []misc.CollectionBridge
		}
		data := PagePage{backend_error, unknown_error, collections}
		files := []string{
			app.Path + "admin/collection/collections.html",
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
