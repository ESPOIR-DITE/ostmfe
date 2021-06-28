package collection

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	"ostmfe/domain/collection"
	"ostmfe/io/collection_io"
	"ostmfe/io/image_io"
)

func CollectionHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", CollectionHandler(app))
	r.Get("/collection_type", CollectionTypeHandler(app))
	r.Get("/new", NewCollectionHandler(app))
	r.Get("/collection_type/new", NewCollectionTypeHandler(app))
	r.Get("/new_stp/{collectionId}", NewCollection2Handler(app))
	r.Get("/edit/{collectionId}", EditCollectionHandler(app))
	r.Post("/create_stp1", CreateCollection1(app))
	r.Post("/create_stp2", CreateCollection2(app))
	r.Post("/collection_type/create", CreateCollectionType(app))

	r.Post("/update_image", UpdateCollectionImageType(app))
	r.Post("/create_history", UpdateCollectionHistoryHandler(app))
	r.Post("/update_details", UpdateCollectionDetailsHandler(app))
	r.Post("/create_image", CreateCollectionImageType(app))

	r.Get("/delete_picture/{imageId}/{collectionId}", DeletePictureHandler(app))
	r.Post("/collectionType/update", UpdateColelctionTypeHandler(app))
	r.Get("/delete/collectionType/{collectionId}", DeleteColelctionTypeHandler(app))
	r.Get("/delete/collection/{collectionId}", DeleteCollectionHandler(app))
	return r
}

func DeleteCollectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collectionId := chi.URLParam(r, "collectionId")
		fmt.Println(collectionId, " <<<collectionId")
		collection, err := collection_io.ReadCollection(collectionId)
		if err != nil {
			fmt.Println(err, " error reading collection")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection", 301)
			return
		}
		collectionType, err := collection_io.ReadWithCollectionId(collection.Id)
		if err != nil {
			fmt.Println(err, " error reading collection type")
		} else {
			_, err := collection_io.DeleteCollection_Type(collectionType.Id)
			if err != nil {
				fmt.Println(err, " error delete collection type")
			}
		}
		_, err = collection_io.DeleteCollection(collection.Id)
		if err != nil {
			fmt.Println(err, " error delete collection")
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Collection Type : "+collection.Name)
		http.Redirect(w, r, "/admin_user/collection/", 301)
		return
	}
}

func UpdateColelctionTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		collectionId := r.PostFormValue("collectionTypeId")
		collectionName := r.PostFormValue("collectionName")
		description := r.PostFormValue("description")

		collectionType, err := collection_io.ReadCollectionTyupe(collectionId)
		if err != nil {
			fmt.Println(err, " error reading collection type")
			fmt.Println(err, " error reading collection")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection/collection_type", 301)
			return
		} else if collectionName != "" {
			if description == "" { // Just checking if the update data was also updating description
				description = collectionType.Description
			}
			collectionTypeObejct := collection.CollectionTypes{collectionType.Id, collectionName, description}
			_, err := collection_io.UpdateCollectionTyupe(collectionTypeObejct)
			if err != nil {
				fmt.Println(err, " error delete collection type")
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Collection Type : "+collectionType.Name)
			http.Redirect(w, r, "/admin_user/collection/collection_type", 301)
			return
		}
		fmt.Println(" Missing CollectionName")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/collection/collection_type", 301)
		return

	}
}

func DeleteColelctionTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collectionTypeId := chi.URLParam(r, "collectionId")
		fmt.Println(" collectionTypeId: " + collectionTypeId)

		collectionType, err := collection_io.ReadCollectionTyupe(collectionTypeId)
		if err != nil {
			fmt.Println(err, " error reading collection type")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection/collection_type", 301)
			return
		} else {
			_, err = collection_io.DeleteCollectionTyupe(collectionTypeId)
			if err != nil {
				fmt.Println(err, " error delete collection type")
			}
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Collection Type : "+collectionType.Name)
		http.Redirect(w, r, "/admin_user/collection/collection_type", 301)
		return
	}
}

func CollectionTypeHandler(app *config.Env) http.HandlerFunc {
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

		collectionType, err := collection_io.ReadCollectionTyupes()
		if err != nil {
			fmt.Println(err, " error reading collections")
		}

		type PagePage struct {
			Backend_error  string
			Unknown_error  string
			Collections    []misc.CollectionBridge
			CollectionType []collection.CollectionTypes
			SidebarData    misc.SidebarData
		}
		data := PagePage{backend_error, unknown_error, collections, collectionType, misc.GetSideBarData("collection", "collectio_type")}
		files := []string{
			app.Path + "admin/collection/collectionType_tables.html",
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

func UpdateCollectionHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		collectionId := r.PostFormValue("collectionId")
		myArea := r.PostFormValue("myArea")

		collectionObject, err := collection_io.ReadCollection(collectionId)
		if err != nil {
			fmt.Println(err, " could not read collection")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection", 301)
			return
		}
		if myArea != "" {
			collectionObject := collection.Collection{collectionId, collectionObject.Name, collectionObject.ProfileDescription, misc.ConvertToByteArray(myArea)}
			_, errr := collection_io.UpdateCollection(collectionObject)
			if errr != nil {
				fmt.Println(errr, " error updating collection")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following Colletion Image : "+collectionObject.Name)
			http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
			return
		}
		fmt.Println("one field is empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
		return
	}
}

func UpdateCollectionDetailsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		collectionId := r.PostFormValue("collectionId")
		collectionName := r.PostFormValue("collectionName")
		description := r.PostFormValue("description")

		collectionObject, err := collection_io.ReadCollection(collectionId)
		if err != nil {
			fmt.Println(err, " could not read collection")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection", 301)
			return
		}
		if collectionName != "" && description != "" {
			collectionObject := collection.Collection{collectionId, collectionName, description, collectionObject.History}
			_, errr := collection_io.UpdateCollection(collectionObject)
			if errr != nil {
				fmt.Println(errr, " error updating collection")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following Colletion Image : "+collectionObject.Name)
			http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
			return
		}
		fmt.Println("one field is empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
		return

	}
}

func CreateCollectionImageType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		collectionId := r.PostFormValue("collectionId")

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
			http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
			return
		}
		collectionImage := collection.Collection_image{"", collectionObject.Id, collectionId, ""}
		CollectionImageHelper := collection.CollectionImageHelper{collectionImage, filesByteArray}

		_, errr := collection_io.CreateCollectionImg(CollectionImageHelper)
		if errr != nil {
			fmt.Println(errr, " error creating CollectionImage")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
			return
		}

		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+collectionObject.Name)
		http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
		return
	}
}

func UpdateCollectionImageType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		file, _, err := r.FormFile("file")
		collectionId := r.PostFormValue("collectionId")
		imageId := r.PostFormValue("imageId")
		collectionImageId := r.PostFormValue("collectionImageId")
		imageType := r.PostFormValue("imageType")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		collectionObject, err := collection_io.ReadCollection(collectionId)
		if err != nil {
			fmt.Println(err, " could not read collection")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection", 301)
			return
		}
		if imageId != "" && imageType != "" && collectionImageId != "" {
			filesArray := []io.Reader{file}
			filesByteArray := misc.CheckFiles(filesArray)
			collectionImage := collection.Collection_image{collectionImageId, imageId, collectionObject.Id, imageType}

			helper := collection.CollectionImageHelper{collectionImage, filesByteArray}
			_, errr := collection_io.UpdateCollectionImg(helper)
			if errr != nil {
				fmt.Println(errr, " error updating collectionImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following Colletion Image : "+collectionObject.Name)
			http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
			return
		}
		fmt.Println("one field is empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
		return

	}
}

func DeletePictureHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageId := chi.URLParam(r, "imageId")
		collectionId := chi.URLParam(r, "collectionId")

		collection, err := collection_io.ReadCollection(collectionId)
		if err != nil {
			fmt.Println(err, " error reading collection")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection", 301)
			return
		}
		//if we can read collection. than we can delete it image.
		_, err = image_io.ReadImage(imageId)
		if err != nil {
			fmt.Println(err, " error reading image")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection", 301)
			return
		}
		//we need to read collection Image so that we can delete
		collectionImage, err := collection_io.ReadCollectionImgWithCollectionId(collectionId)
		if err != nil {
			fmt.Println(err, " error reading collection image")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/collection", 301)
			return
		}
		_, err = image_io.DeleteImage(imageId)
		if err != nil {
			fmt.Println(err, " error reading image")
		}
		_, err = collection_io.DeleteCollectionImg(collectionImage.Id)
		if err != nil {
			fmt.Println(err, " error reading collection image")
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Collection Type : "+collection.Name)
		http.Redirect(w, r, "/admin_user/collection/edit/"+collectionId, 301)
		return
	}
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
			SidebarData    misc.SidebarData
		}
		data := PagePage{backend_error, unknown_error, collectionTypes, misc.GetSideBarData("collection", "collection-type")}
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
		collectionImage := collection.Collection_image{"", collectionObject.Id, collectionId, ""}
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
		http.Redirect(w, r, "/admin_user/place", 301)
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
		collectionId := chi.URLParam(r, "collectionId")

		type DataPage struct {
			CollectionData CollectionData
			SidebarData    misc.SidebarData
		}
		data := DataPage{GetCollectionData(collectionId), misc.GetSideBarData("collection", "collection-type")}
		files := []string{
			app.Path + "admin/collection/edit_collection.html",
			app.Path + "admin/template/navbar.html",
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

		collectionType, err := collection_io.ReadCollectionTyupes()
		if err != nil {
			fmt.Println(err, " error reading collections")
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}

		type PagePage struct {
			Backend_error  string
			Unknown_error  string
			Collections    []misc.CollectionBridge
			CollectionType []collection.CollectionTypes
			SidebarData    misc.SidebarData
			AdminName      string
			AdminImage     string
		}
		data := PagePage{backend_error,
			unknown_error, collections,
			collectionType,
			misc.GetSideBarData("collection", "collection"),
			adminName, adminImage}
		files := []string{
			app.Path + "admin/collection/collections.html",
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
