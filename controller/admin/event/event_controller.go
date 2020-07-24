package event

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	event2 "ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	partner2 "ostmfe/domain/partner"
	place2 "ostmfe/domain/place"
	project2 "ostmfe/domain/project"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/place_io"
	"ostmfe/io/project_io"
	"time"
)

func EventHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", EventsHandler(app))
	r.Get("/new", NewEventsHandler(app))
	r.Get("/edite/{eventId}", EditEventsHandler(app))
	r.Post("/create", CreateEventHandler(app))
	r.Post("/create-history", CreateEventHistoryEventHandler(app))
	r.Post("/update", UpdateEventHandler(app))
	r.Get("/picture/{eventId}", EventPicture(app))
	return r
}

func CreateEventHistoryEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		var histories history2.Histories
		var eventHistory event2.EventHistory

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		mytextarea := r.PostFormValue("mytextarea")
		eventId := r.PostFormValue("eventId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		//Creating EventHistory and History
		fmt.Println("eventIed: ", eventId, " test>>>>", mytextarea)
		if eventId != "" && mytextarea != "" {
			//Creating Histories Object
			historyObject := history2.Histories{"", misc.ConvertToByteArray(mytextarea)}
			histories, err = history_io.CreateHistorie(historyObject)
			if err != nil {
				fmt.Println("could not create history and wont create Event history")
			} else {
				//creating Event History
				evenHistory := event2.EventHistory{"", histories.Id, eventId}
				eventHistory, err = event_io.CreateEventHistory(evenHistory)
				if err != nil {
					fmt.Println("could not create event history")
					_, err := history_io.DeleteHistorie(histories.Id)
					if err != nil {
						fmt.Println("error deleting history")
					}
				}
			}

		}

		//creating EVentImage
		eventImageObject := event2.EventImage{"", eventId, ""}
		eventImageHelperObject := event2.EventImageHelper{eventImageObject, filesByteArray}
		_, errx := event_io.CreateEventImg(eventImageHelperObject)
		/**
		Rolling back
		*/
		if errx != nil {
			fmt.Println(err, " error could not create eventImage Proceeding into rol back.....")
			if histories.Id != "" {
				fmt.Println(err, " Deleting histories of this event....")
				_, err := history_io.DeleteHistorie(histories.Id)
				if err != nil {
					fmt.Println(err, " !!!!!error could not delete history")
				} else {
					fmt.Println(err, " Deleted")
				}
			}
			if eventHistory.Id != "" {
				fmt.Println(err, " Deleting Event histories of this event....")
				_, err := event_io.DeleteEventHistory(eventHistory.Id)
				if err != nil {
					fmt.Println(err, " !!!!!error could not delete history")
				} else {
					fmt.Println(err, " Deleted")
				}
			}
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event/picture/"+eventId, 301)
			return
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Event : ")
		http.Redirect(w, r, "/admin_user", 301)
		return
	}
}

func EventPicture(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventId := chi.URLParam(r, "eventId")
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
		//Checking if the eventiId passed is for an existing event
		event, err := event_io.ReadEvent(eventId)
		if err != nil {
			fmt.Println(err, " error reading the event")
			if app.Session.GetString(r.Context(), "user-read-error") != "" {
				app.Session.Remove(r.Context(), "user-read-error")
			}
			app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event", 301)
			return
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
			Event         event2.Event
			Backend_error string
			Unknown_error string
		}
		data := PageData{projects, partners, event, backend_error, unknown_error}
		files := []string{
			app.Path + "admin/event/new_event_picture.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
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
func EditEventsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventId := chi.URLParam(r, "eventId")
		event, err := event_io.ReadEvent(eventId)
		if err != nil {
			fmt.Println(err, " error reading an event")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event", 301)
			return
		}
		eventData := misc.GetEventDate(eventId)
		type PageData struct {
			Event     event2.Event
			EventData misc.EventData
		}
		date := PageData{event, eventData}
		files := []string{
			app.Path + "admin/event/edite_event.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "base_templates/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, date)
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
			app.Path + "admin/event/new_event.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
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

func EventsHandler(app *config.Env) http.HandlerFunc {
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
		events, err := event_io.ReadEvents()
		if err != nil {
			fmt.Println(err, " error reading Users")
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
			Events        []event2.Event
		}
		data := PageData{projects, partners, backend_error, unknown_error, events}
		files := []string{
			app.Path + "admin/event/events.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
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

func CreateEventHandler(app *config.Env) http.HandlerFunc {
	/***
	Here we create a new event

	*/
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		var newEvent event2.Event
		event_name := r.PostFormValue("event_name")
		date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		project := r.PostFormValue("project")
		description := r.PostFormValue("description")
		partner := r.PostFormValue("partner")
		latlng := r.PostFormValue("latlng")
		place := r.PostFormValue("place")

		if event_name != "" {
			eventObject := event2.Event{"", event_name, date, description}
			errs := errors.New("")
			newEvent, errs = event_io.CreateEvent(eventObject)
			if errs != nil {
				fmt.Println(errs, " error when creating a new event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/new", 301)
				return
			}
			if partner != "" && newEvent.Id != "" {
				eventPartner := event2.EventPartener{"", partner, newEvent.Id, ""}
				_, err := event_io.CreateEventPartener(eventPartner)
				if err != nil {
					fmt.Println(err, " error when creating event partner")
				}
			}

			if project != "" && newEvent.Id != "" {
				eventProject := event2.EventProject{project, newEvent.Id, ""}
				_, err := event_io.CreateEventProject(eventProject)
				if err != nil {
					fmt.Println(err, " error when creating event project")
				}
			}

			latitude, longitude := misc.SeparateLatLng(latlng)
			place := place2.Place{"", place, latitude, longitude, ""}
			newPlace, err := place_io.CreatePlace(place)
			if err != nil {
				fmt.Println(err, " error when creating a new place")
				//_, errr := event_io.DeleteEventPartener(eventPartner.PartenerId)
				//if errr != nil {
				//	fmt.Println(errr, " error when deleting event partner in rolling back action")
				//}
				//_, err := event_io.DeleteEvent(newEvent.Id)
				//if err != nil {
				//	fmt.Println(err, " error when deleting event in rolling back action")
				//}
				//_, errrr := event_io.DeleteEventProject(eventProject.EventId)
				//if errrr != nil {
				//	fmt.Println(err, " error when deleting Project in rolling back action")
				//}
				//if app.Session.GetString(r.Context(), "user-create-error") != "" {
				//	app.Session.Remove(r.Context(), "user-create-error")
				//}
				//app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				//http.Redirect(w, r, "/admin_user/users/new", 301)
				//return
			} else {
				eventPlace := event2.EventPlace{newPlace.Id, eventObject.Id, ""}
				_, err := event_io.CreateEventPlace(eventPlace)
				if err != nil {
					//fmt.Println(err, " error when creating Event place")
					//_, errr := event_io.DeleteEventPartener(eventPartner.PartenerId)
					//if errr != nil {
					//	fmt.Println(errr, " error when deleting event partner in rolling back action")
					//}
					//_, err := event_io.DeleteEvent(newEvent.Id)
					//if err != nil {
					//	fmt.Println(err, " error when deleting event in rolling back action")
					//}
					//_, errrr := event_io.DeleteEventProject(eventProject.EventId)
					//if errrr != nil {
					//	fmt.Println(err, " error when deleting event Project in rolling back action")
					//}
					//_, errrrr := place_io.DeletePlace(newPlace.Id)
					//if errrrr != nil {
					//	fmt.Println(err, " error when deleting place in rolling back action")
					//}
					//if app.Session.GetString(r.Context(), "user-create-error") != "" {
					//	app.Session.Remove(r.Context(), "user-create-error")
					//}
					//app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					//http.Redirect(w, r, "/admin_user/users/new", 301)
					//return
				}
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Event : "+event_name)
			http.Redirect(w, r, "/admin_user/event/picture/"+newEvent.Id, 301)
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

func UpdateEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id := r.PostFormValue("id")
		name := r.PostFormValue("name")
		description := r.PostFormValue("description")
		date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		event, err := event_io.ReadEvent(id)
		if err != nil {
			fmt.Println(err, " could not read event")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+id, 301)
			return
		}
		//we checking if there is a need of updating
		if event.Name != name && event.Id != id && event.Date != date {
			event := event2.Event{id, name, date, description}
			eventAfterUpdate, err := event_io.UpdateEvent(event)
			if err != nil {
				fmt.Println(err, " could not update event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+id, 301)
				return
			} else {
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following User : "+eventAfterUpdate.Name)
				http.Redirect(w, r, "/admin_user/event", 301)
				return
			}
		}
		fmt.Println(err, " No need for Update because you haven't made any change")
		http.Redirect(w, r, "/admin_user/event", 301)

	}
}
