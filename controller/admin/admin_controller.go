package admin

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/collection"
	event2 "ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	partner2 "ostmfe/domain/partner"
	people2 "ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	project2 "ostmfe/domain/project"
	user2 "ostmfe/domain/user"
	"ostmfe/io/collection_io"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
	"ostmfe/io/project_io"
	"ostmfe/io/user_io"
	"time"
)

/***
- user-create-error : This is session message reporting an error occurred when there is an error when creating a new USER.
- creation-successful : This is session message reporting an successful creation.
*/

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))

	r.Get("/users", UserHandler(app))
	r.Get("/users/new", NewUserHandler(app))
	r.Get("/users/edit", EditUserHandler(app))
	r.Post("/users/create", CreateUserHandler(app))

	r.Get("/event", EventsHandler(app))
	r.Get("/event/new", NewEventsHandler(app))
	r.Get("/event/edite", EditEventsHandler(app))
	r.Post("/event/create", CreateEventHandler(app))

	r.Get("/project", ProjectsHandler(app))
	r.Get("/project/new", NewProjectsHandler(app))
	r.Get("/project/new_history/{projectId}", NewProjectHistoryHandler(app))
	r.Get("/project/edit/{projectId}", EditeProjectsHandler(app))
	r.Post("/project/create", CreateProjectHandler(app))
	r.Post("/project/create_project_history", CreateProjectHistoryHandler(app))
	//r.Get("/projects/delete",DeleteProjectsHandler(app))

	r.Get("/partner", PartenersHandler(app))
	r.Get("/partner/new", NewPartenersHandler(app))
	r.Get("/partner/edit", EditePartenersHandler(app))
	r.Post("/partner/create", CreatePartenersHandler(app))

	r.Get("/place", PlacesHandler(app))
	r.Get("/place/new", NewPlacesHandler(app))
	r.Get("/place/edit", EditPlacesHandler(app))
	r.Post("/place/create_stp1", CreateStp1Handler(app))
	r.Post("/place/create_stp2", CreatePlaceStp2Handler(app))
	r.Get("/place/new_stp2/{placeId}", NewPlaceStp2Handler(app))

	r.Get("/collection", CollectionHandler(app))
	r.Get("/collection/new", NewCollectionHandler(app))
	r.Get("/collection_type/new", NewCollectionTypeHandler(app))
	r.Get("/collection/new_stp/{collectionId}", NewCollection2Handler(app))
	r.Get("/collection/edit", EditCollectionHandler(app))
	r.Post("/collection/create_stp1", CreateCollection1(app))
	r.Post("/collection/create_stp2", CreateCollection2(app))
	r.Post("/collection_type/create", CreateCollectionType(app))

	r.Get("/history", HistoryHandler(app))
	r.Get("/history/new", NewHistoryHandler(app))
	r.Get("/history/edit", EditHistoryHandler(app))

	r.Get("/people", PeopleHandler(app))
	r.Get("/people_category/new", NewPeopleCategoryHandler(app))
	r.Get("/people/new", NewPeopleHandler(app))
	r.Get("/people/new-stp2/{peopleId}", NewPeoplestp2Handler(app))
	r.Get("/people/edit", EditPeopleHandler(app))
	r.Post("/people/create_stp1", CreatePeopleHandler(app))
	r.Post("/people/create_stp2", CreatePeopleStp2Handler(app))
	r.Post("/people_category/create", CreatePeopleCategoryHandler(app))

	return r
}

func CreateProjectHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		projectId := r.PostFormValue("projectId")
		history := r.PostFormValue("history")
		description := r.PostFormValue("description")

		if projectId != "" && history != "" {
			project, err := project_io.ReadProject(projectId)
			if err != nil {
				fmt.Println(err, " error reading Project")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new_history/"+projectId, 301)
				return
			}
			historyByteArray := []byte(history)
			historyObject := history2.History{"", description, historyByteArray, time.Now()}
			history, err := history_io.CreateHistory(historyObject)
			if err != nil {
				fmt.Println(err, " error creating History")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new_history/"+projectId, 301)
				return
			}
			projectHistoryObject := project2.ProjectHistory{"", projectId, history.Id}
			_, errr := project_io.CreateProjectHistory(projectHistoryObject)
			if errr != nil {
				_, err := history_io.DeleteHistory(history.Id)
				if err != nil {
					fmt.Println(err, " error Delete History")
				}
				fmt.Println(err, " error creating History")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new_history/"+projectId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new People Type : "+project.Title)
			http.Redirect(w, r, "/admin_user", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/project/new_history/"+projectId, 301)
		return
	}
}

func NewProjectHistoryHandler(app *config.Env) http.HandlerFunc {
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
		projectId := chi.URLParam(r, "projectId")
		project, err := project_io.ReadProject(projectId)
		if err != nil {
			fmt.Println(" error reading project")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/new", 301)
			return
		}

		type PageData struct {
			Project       project2.Project
			Backend_error string
			Unknown_error string
		}
		data := PageData{project, backend_error, unknown_error}
		files := []string{
			app.Path + "admin/project/new_project_history.html",
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

func CreatePeopleCategoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		category := r.PostFormValue("category")
		if category != "" {
			people := people2.PeopleCategory{"", category}
			peopleCategory, err := people_io.CreatePeopleCategory(people)

			if err != nil {
				fmt.Println(err, " error creating people Category")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people_category/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new People Type : "+peopleCategory.Category)
			http.Redirect(w, r, "/admin_user/people_category/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people_category/new", 301)
		return
	}
}

func NewPeopleCategoryHandler(app *config.Env) http.HandlerFunc {
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

		peoples, err := people_io.ReadPeopleCategorys()
		if err != nil {
			fmt.Println(err, " There is an error when reading all the people category")
		}
		type PageData struct {
			Peoples       []people2.PeopleCategory
			Backend_error string
			Unknown_error string
		}
		data := PageData{peoples, backend_error, unknown_error}
		files := []string{
			app.Path + "admin/people/peopleType_tables.html",
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
				http.Redirect(w, r, "/admin_user/collection_type/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Collection Type : "+collectionType.Name)
			http.Redirect(w, r, "/admin_user/collection_type/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/collection_type/new", 301)
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
		collectionImage := collection.Collection_image{"", collectionObject.Id, ""}
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
			collectionObject := collection.Collection{"", collection_name, brief, history}
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

func NewPeoplestp2Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peopleId := chi.URLParam(r, "peopleId")
		people, err := people_io.ReadPeople(peopleId)
		if err != nil {
			fmt.Println(err, " error reading the people")
			if app.Session.GetString(r.Context(), "user-read-error") != "" {
				app.Session.Remove(r.Context(), "user-read-error")
			}
			app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/people/new", 301)
			return
		}
		type PageData struct {
			People people2.People
		}
		data := PageData{People: people}
		files := []string{
			app.Path + "admin/people/image_people.html",
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

func CreatePeopleStp2Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var place place2.Place
		r.ParseForm()
		//fileslist := r.Form["file"]

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		history := r.PostFormValue("history")
		date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		peopleId := r.PostFormValue("peopleId")
		description := r.PostFormValue("description")
		latlng := r.PostFormValue("latlng")
		title := r.PostFormValue("title")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		if history != "" && peopleId != "" && title != "" && description != "" {
			historyByteArray := []byte(history)
			historyObejct := history2.History{"", description, historyByteArray, date}
			historynew, Err := history_io.CreateHistory(historyObejct)
			if Err != nil {
				fmt.Println(err, " error creating a new history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
				return
			}
			// I am checking if the this person is relative to a particular place
			if latlng != "" {
				latitude, longitude := misc.SeparateLatLng(latlng)
				placeObject := place2.Place{"", title, latitude, longitude, ""}
				place, err = place_io.CreatePlace(placeObject)
				if err != nil {
					fmt.Println(err, " error creating a new Place")
					_, err := history_io.DeleteHistory(historynew.Id)
					if err != nil {
						fmt.Println(err, " error could not delete history")
					}
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
					return
				}
			}

			peopleHistory := people2.PeopleHistory{"", peopleId, history}

			peopleHistoryNew, err := people_io.CreatePeopleHistory(peopleHistory)
			if err != nil {
				fmt.Println(err, " error creating a new people history")
				_, err := history_io.DeleteHistory(historynew.Id)
				if err != nil {
					fmt.Println(err, " error could not delete history")
				}
				//we need to make sure that the place was created
				if place.Id != "" {
					_, err := place_io.DeletePlace(place.Id)
					if err != nil {
						fmt.Println(err, " error could not delete history")
					}
				}

				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
				return
			}
			peopleImageObject := people2.People_image{"", peopleId, "", ""}
			peopleImage := people2.PlaceImageHelper{peopleImageObject, filesByteArray}
			_, errr := people_io.CreatePeopleImage(peopleImage)
			/***
			RolBack
			*/
			if errr != nil {
				_, err := history_io.DeleteHistory(historynew.Id)
				if err != nil {
					fmt.Println(err, " error could not delete history")
				}
				//we need to make sure that the place was created
				if place.Id != "" {
					_, err := place_io.DeletePlace(place.Id)
					if err != nil {
						fmt.Println(err, " error could not delete history")
					}
				}
				//Now deleting People History
				_, errx := people_io.DeletePeopleHistory(peopleHistoryNew.Id)
				if errx != nil {
					fmt.Println(err, " error could not delete people History")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/new_stp2/"+peopleId, 301)
				return
			}

			newPeople, errr := people_io.ReadPeople(peopleId)
			if errr != nil {
				fmt.Println(err, " error reading Place Line: 121")
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new person : "+newPeople.Name)
			http.Redirect(w, r, "/admin_user", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/new_stp2/"+peopleId, 301)
		return
	}
}

func CreatePeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("partner_name")
		surname := r.PostFormValue("surname")
		profession := r.PostFormValue("profession")
		b_date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("b_date"))
		d_date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("d_date"))
		origin := r.PostFormValue("origin")
		if name != "" && surname != "" && profession != "" && origin != "" {
			peopleObject := people2.People{"", name, surname, b_date, d_date, origin, profession}
			people, err := people_io.CreatePeople(peopleObject)
			if err != nil {
				fmt.Println(err, " error creating a new people")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/new", 301)
				return
			}
			if people.Id != "" {
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+people.Name)
				http.Redirect(w, r, "/admin_user/people/new-stp2/"+people.Id, 301)
				return
			}
		}
		fmt.Println("One of the field is missing or newPlace.Id is nil")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/new", 301)
		return
	}
}

func EditPeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/edit_people.html",
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

func NewPeopleHandler(app *config.Env) http.HandlerFunc {
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
			app.Path + "admin/people/new_people.html",
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

func PeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/people.html",
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

func CreatePartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		partner_name := r.PostFormValue("partner_name")
		description := r.PostFormValue("description")
		url := r.PostFormValue("url")

		if url != "" && description != "" && partner_name != "" {
			partner := partner2.Partner{"", partner_name, description, url}
			partnerResult, err := partner_io.CreatePartner(partner)
			if err != nil {
				fmt.Println(err, " error creating a new partner")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/partner/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new partenr : "+partnerResult.Name)
			http.Redirect(w, r, "/admin_user", 301)
			return
		}
	}
}

func EditePartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/edit_partner.html",
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

func PartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/partner/partner.html",
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

func NewPartenersHandler(app *config.Env) http.HandlerFunc {
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
			app.Path + "admin/partner/new_partner.html",
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
func contains(slice []string) [][]byte {
	var set [][]byte
	for _, s := range slice {
		v := []byte(s)
		set = append(set, v)
	}
	return set
}
func CreateProjectHandler(app *config.Env) http.HandlerFunc {
	/***
	Here we create a new project

	*/
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		project_name := r.PostFormValue("project_name")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>>>>")
		}

		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		fmt.Println(project_name, "<<<Project Name|| description>>>", description)

		if project_name != "" && description != "" {
			project := project2.Project{"", project_name, description}
			new_project, err := project_io.CreateProject(project)
			if err != nil {
				fmt.Println(err, " could not create project Line: 190")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new", 301)
				return
			}

			projectImage := project2.ProjectImage{"", new_project.Id, "", ""}
			helper := project2.ProjectImageHelper{filesByteArray, projectImage}
			_, errr := project_io.CreateProjectImage(helper)
			if errr != nil {
				fmt.Println(errr, " error creating projectImage")
				_, err := project_io.DeleteProject(new_project.Id)
				if err != nil {
					fmt.Println(err, " error deleting project")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+project_name)
			http.Redirect(w, r, "/admin_user/project/new_history/"+new_project.Id, 301)
			return
			//event_name := r.PostFormValue("event_name")
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/project/new", 301)
		return

	}

}

func CreateEventHandler(app *config.Env) http.HandlerFunc {
	/***
	Here we create a new event

	*/
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		event_name := r.PostFormValue("event_name")
		date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		project := r.PostFormValue("project")
		partner := r.PostFormValue("partner")
		latlng := r.PostFormValue("latlng")
		place := r.PostFormValue("place")

		if event_name != "" {
			eventObject := event2.Event{"", event_name, date}
			newEvent, err := event_io.CreateEvent(eventObject)
			if err != nil {
				fmt.Println(err, " error when creating a new event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			}
			eventPartner := event2.EventPartener{"", partner, newEvent.Id, ""}
			_, err = event_io.CreateEventPartener(eventPartner)
			if err != nil {
				fmt.Println(err, " error when creating event partner")
				/**
				Rolling back
				*/
				_, err := event_io.DeleteEvent(newEvent.Id)
				if err != nil {
					fmt.Println(err, " error when deleting event in rolling back action")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			}

			eventProject := event2.EventProject{project, newEvent.Id, ""}
			_, err = event_io.CreateEventProject(eventProject)
			if err != nil {
				fmt.Println(err, " error when creating event project")
				/**
				Rolling back
				*/
				_, err := event_io.DeleteEvent(newEvent.Id)
				if err != nil {
					fmt.Println(err, " error when deleting event in rolling back action")
				}
				_, errr := event_io.DeleteEventPartener(eventPartner.PartenerId)
				if errr != nil {
					fmt.Println(errr, " error when deleting event partner in rolling back action")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred when, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			}
			latitude, longitude := misc.SeparateLatLng(latlng)
			place := place2.Place{"", place, latitude, longitude, ""}
			newPlace, err := place_io.CreatePlace(place)
			if err != nil {
				fmt.Println(err, " error when creating a new place")

				_, errr := event_io.DeleteEventPartener(eventPartner.PartenerId)
				if errr != nil {
					fmt.Println(errr, " error when deleting event partner in rolling back action")
				}
				_, err := event_io.DeleteEvent(newEvent.Id)
				if err != nil {
					fmt.Println(err, " error when deleting event in rolling back action")
				}
				_, errrr := event_io.DeleteEventProject(eventProject.EventId)
				if errrr != nil {
					fmt.Println(err, " error when deleting Project in rolling back action")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return

			} else {
				eventPlace := event2.EventPlace{newPlace.Id, eventObject.Id, ""}
				_, err := event_io.CreateEventPlace(eventPlace)
				if err != nil {
					fmt.Println(err, " error when creating Event place")
					_, errr := event_io.DeleteEventPartener(eventPartner.PartenerId)
					if errr != nil {
						fmt.Println(errr, " error when deleting event partner in rolling back action")
					}
					_, err := event_io.DeleteEvent(newEvent.Id)
					if err != nil {
						fmt.Println(err, " error when deleting event in rolling back action")
					}
					_, errrr := event_io.DeleteEventProject(eventProject.EventId)
					if errrr != nil {
						fmt.Println(err, " error when deleting event Project in rolling back action")
					}
					_, errrrr := place_io.DeletePlace(newPlace.Id)
					if errrrr != nil {
						fmt.Println(err, " error when deleting place in rolling back action")
					}
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/users/new", 301)
					return
				}
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Event : "+event_name)
			http.Redirect(w, r, "/admin_user", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/users/new", 301)
		return
	}

}

func CreateUserHandler(app *config.Env) http.HandlerFunc {
	/***
	we are getting the form from html
	we grab all the fields corresponding to the name assigned to them
	we create an object with the records collected from the html
	we then send the object to the backend, if an error occurs we will redirect back to new user html file to try again.
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("name")
		surname := r.PostFormValue("surname")
		email := r.PostFormValue("email")
		fmt.Println(name, "<<name  surname>>", surname, "  email>>", email)
		if name != "" && surname != "" && email != "" {
			user := user2.Users{email, name, surname}
			newUser, err := user_io.CreateUser(user)
			if err != nil {
				fmt.Println(err, "error creating new user line: 57")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			} else {
				fmt.Println(err, "Creation of a new user successful")
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
					return
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new user : "+newUser.Name)
				http.Redirect(w, r, "/admin_user", 301)
				return
			}
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/users/new", 301)
		return
	}
}

func EditHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/edit_history.html",
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

func NewHistoryHandler(app *config.Env) http.HandlerFunc {
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
			app.Path + "admin/collection/new_history.html",
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

func HistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/collection/history.html",
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
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}

}

func EditPlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/place/edit_places.html",
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
		files := []string{
			app.Path + "admin/place/places.html",
			app.Path + "admin/template/navbar.html",
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

func DeleteProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/edit_events.html",
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

func EditeProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")
		selectedProjest := misc.GetProjectEditable(projectId)
		type PageData struct {
			Project misc.ProjectEditable
		}
		data := PageData{selectedProjest}
		files := []string{
			app.Path + "admin/project/edite_project.html",
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

func NewProjectsHandler(app *config.Env) http.HandlerFunc {
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
			app.Path + "admin/project/new_project.html",
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

func ProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/project/projects.html",
			app.Path + "admin/template/navbar.html",
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

func EditEventsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/edit_events.html",
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

func NewEventsHandler(app *config.Env) http.HandlerFunc {
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
		//Reading all the Projects
		projects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading all the projects")
		}
		partners, err := partner_io.ReadPartners()
		if err != nil {
			fmt.Println(err, " error reading all the partners")
		}
		type PageData struct {
			Projects      []project2.Project
			Partners      []partner2.Partner
			Backend_error string
			Unknown_error string
		}
		data := PageData{projects, partners, backend_error, unknown_error}
		files := []string{
			app.Path + "admin/new_event.html",
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

func EventsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/events.html",
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

func EditUserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/edit_user.html",
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

func UserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "admin/users.html",
			app.Path + "admin/template/navbar.html",
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

func NewUserHandler(app *config.Env) http.HandlerFunc {
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
			app.Path + "admin/new_user.html",
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

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var success_notice string
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			success_notice = app.Session.GetString(r.Context(), "creation-successful")
			app.Session.Remove(r.Context(), "creation-successful")
		}
		type PageData struct {
			Success_notice string
		}
		data := PageData{success_notice}
		files := []string{
			app.Path + "admin/admin.html",
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
