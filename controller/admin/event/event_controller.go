package event

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	event3 "ostmfe/controller/event"
	"ostmfe/controller/misc"
	museum "ostmfe/domain"
	"ostmfe/domain/comment"
	"ostmfe/domain/contribution"
	event2 "ostmfe/domain/event"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	partner2 "ostmfe/domain/partner"
	"ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	project2 "ostmfe/domain/project"
	io2 "ostmfe/io"
	"ostmfe/io/comment_io"
	"ostmfe/io/event_io"
	"ostmfe/io/group_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
	"ostmfe/io/project_io"
)

func EventHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", EventsHandler(app))
	r.Get("/new", NewEventsHandler(app))
	r.Get("/edit/{eventId}", EditEventsHandler(app))
	r.Post("/create", CreateEventHandler(app))
	r.Post("/create-history", CreateEventHistoryEventHandler(app))
	r.Post("/create-pictures", CreateEventPictureEventHandler(app))
	r.Post("/update", UpdateEventHandler(app))
	r.Get("/picture/{eventId}", EventPicture(app))
	r.Get("/delete_image/{imageId}/{eventId}/{EventImageId}", DeleteImagehandler(app))

	r.Post("/update_history", UpdateHistoryHandler(app))
	r.Post("/update_place", UpdatePlaceHandler(app))
	r.Post("/update_details", UpdateDetailsHandler(app))
	r.Post("/update_pictures", UpdatePicturesHandler(app))
	r.Post("/add_pictures", AddPictureHandler(app))
	r.Post("/add_people", AddPeopleHandler(app))
	r.Post("/add_group", AddGroupHandler(app))
	r.Post("/update_event_place", UpdateEventPlaceHandler(app))
	r.Post("/create-page-flow", CreatePageFlowHandler(app))

	r.Get("/delete-page-fLow/{pageFlowId}/{eventId}", DeletePageFlowHandler(app))
	r.Get("/delete/{eventId}", DeleteEventHandler(app))
	r.Get("/delete_people/{peopleId}/{eventId}", DeletepeopleEventHandler(app))
	r.Get("/delete_group/{groupId}/{eventId}", DeleteGroupEventHandler(app))
	r.Get("/delete_comment/{commentId}/{eventCommentId}", DeleteCommentEventHandler(app))
	r.Get("/activate_comment/{commentId}/{eventId}", ActivateCommentHandler(app))

	//Gallery
	r.Post("/create-gallery", CreateEventGalleryHandler(app))
	r.Get("/delete-gallery/{pictureId}/{eventId}/{eventGalleryId}", DeleteGalleryHandler(app))

	return r
}

func DeletePageFlowHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		pageFlowId := chi.URLParam(r, "pageFlowId")
		eventId := chi.URLParam(r, "eventId")

		_, err := event_io.DeleteEventPageFlow(pageFlowId)
		if err != nil {
			fmt.Println(err, " error deleting page flow!")
		}
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
	}
}

func CreatePageFlowHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		eventId := r.PostFormValue("eventId")
		pageFlowTitle := r.PostFormValue("pageFlowTitle")
		scr := r.PostFormValue("scr")
		fmt.Println(scr, "  src", pageFlowTitle, "  page flow", eventId, " eventId")
		if scr != "" && eventId != "" && pageFlowTitle != "" {
			_, err := event_io.CreateEventPageFlow(contribution.EventPageFlow{"", eventId, pageFlowTitle, scr})
			if err != nil {
				fmt.Println(err, " error creating page flow!")
			}
		}
		fmt.Print("")
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
		return
	}
}

func ActivateCommentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentId := chi.URLParam(r, "commentId")
		eventId := chi.URLParam(r, "eventId")
		result := misc.ActivateComment(commentId)
		fmt.Print("Activation Result: ", result)
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
		return
	}
}

func DeleteGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pictureId := chi.URLParam(r, "pictureId")
		eventId := chi.URLParam(r, "eventId")
		eventGalleryId := chi.URLParam(r, "eventGalleryId")

		//Deleting project
		gallery, err := image_io.DeleteGalery(pictureId)
		if err != nil {
			fmt.Println("error deleting gallery")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		} else {
			_, err := event_io.DeleteEventGalery(eventGalleryId)
			if err != nil {
				fmt.Println("error deleting group gallery")
				fmt.Println("ROLLING BACK!!!")
				_, err := image_io.UpdateGallery(gallery)
				if err != nil {
					fmt.Println("error updating gallery")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/eventId/edit/"+eventId, 301)
				return
			}
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: group Gallery. ")
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
		return
	}
}

func CreateEventGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		eventId := r.PostFormValue("eventId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if eventId != "" && description != "" {
			galery := image.Galery{"", content, description}
			galleryObject, err := image_io.CreateGalery(galery)
			if err != nil {
				fmt.Println(err, " error creating gallery")
			} else {
				eventGalery := event2.EventGalery{"", eventId, galleryObject.Id}
				_, err := event_io.CreateEventGalery(eventGalery)
				if err != nil {
					fmt.Println(err, " error creating EventGallery")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
		return
	}
}

func DeleteCommentEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		commentId := chi.URLParam(r, "commentId")
		eventCommentId := chi.URLParam(r, "eventCommentId")

		if eventCommentId != "" && commentId != "" {
			eventComment, err := comment_io.ReadCommentEvent(eventCommentId)
			if err != nil {
				fmt.Println(err, " error deleting event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event", 301)
				return
			}
			_, err2 := comment_io.DeleteComment(eventComment.CommentId)
			if err2 != nil {
				fmt.Println(err, " error deleting comment")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event", 301)
				return
			}
			_, err1 := comment_io.DeleteCommentEvent(eventComment.Id)
			if err1 != nil {
				fmt.Println(err, " error deleting event comment")
			}

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventComment.EventId, 301)
			return
		}
		fmt.Println(" error field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/event", 301)
		return

	}
}

func DeleteGroupEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		groupId := chi.URLParam(r, "groupId")
		eventId := chi.URLParam(r, "eventId")

		if eventId != "" && groupId != "" {
			eventGroup, err := event_io.ReadEventGroupWithBoth(eventId, groupId)
			if err != nil {
				fmt.Println(err, " error reading eventGroup")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
			_, err = event_io.DeleteEventGroup(eventGroup.Id)
			if err != nil {
				fmt.Println(err, " error deleting eventGroup")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		fmt.Println(" error field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/event", 301)
		return

	}
}

func AddGroupHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		EventId := r.PostFormValue("EventId")
		groupId := r.PostFormValue("groupId")

		if EventId != "" && groupId != "" {

			event, err := event_io.ReadEvent(EventId)
			if err != nil {
				fmt.Println(err, " error reading event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event", 301)
				return
			}
			eventGroupObject := event2.EventGroup{"", EventId, groupId}
			_, err = event_io.CreateEventGroup(eventGroupObject)
			if err != nil {
				fmt.Println(err, " error creating  eventGroup")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+EventId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following Event: "+event.Name)
			http.Redirect(w, r, "/admin_user/event/edit/"+EventId, 301)
			return
		}

		fmt.Println("field empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/event", 301)
		return

	}
}

func DeletepeopleEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		peopleId := chi.URLParam(r, "peopleId")
		eventId := chi.URLParam(r, "eventId")

		if eventId != "" && peopleId != "" {
			eventPeople, err := event_io.ReadEventPeopleWithBoth(eventId, peopleId)
			if err != nil {
				fmt.Println(err, " error reading eventPeople")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
			_, err = event_io.DeleteEventPeople(eventPeople.Id)
			if err != nil {
				fmt.Println(err, " error deleting eventPeople")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted ")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		fmt.Println(" error field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/event", 301)
		return

	}
}

func UpdateEventPlaceHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		EventId := r.PostFormValue("EventId")
		placeId := r.PostFormValue("placeId")

		//Checking if place exist
		place, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println(err, " error reading place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+EventId, 301)
			return
		}

		if EventId != "" && placeId != "" {
			eventPlace, err := event_io.ReadEventPlaceOf(EventId)
			if err != nil {
				fmt.Println(err, " error reading event Place")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+EventId, 301)
				return
			} else {
				eventPlaceObject := event2.EventPlace{eventPlace.Id, placeId, EventId, eventPlace.Description}
				_, err := event_io.UpdateEventPlace(eventPlaceObject)
				if err != nil {
					fmt.Println(err, " error updating EventPlace")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/event/edit/"+EventId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully added the following Place: "+place.Title)
				http.Redirect(w, r, "/admin_user/event/edit/"+EventId, 301)
				return
			}
		}
		fmt.Println("field empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/event", 301)
		return

	}
}

func AddPeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		EventId := r.PostFormValue("EventId")
		peopleId := r.PostFormValue("peopleId")

		if EventId != "" && peopleId != "" {

			event, err := event_io.ReadEvent(EventId)
			if err != nil {
				fmt.Println(err, " error reading event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event", 301)
				return
			}
			eventPeopleObject := event2.EventPeople{"", EventId, peopleId}
			_, err = event_io.CreateEventPeople(eventPeopleObject)
			if err != nil {
				fmt.Println(err, " error reading event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+EventId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted the following Event: "+event.Name)
			http.Redirect(w, r, "/admin_user/event/edit/"+EventId, 301)
			return
		}

		fmt.Println("field empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/event", 301)
		return

	}
}

func DeleteEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		eventId := chi.URLParam(r, "eventId")

		fmt.Println("deleting: ", eventId)
		//checking if we have that event in our database
		event, err := event_io.ReadEvent(eventId)
		if err != nil {
			fmt.Println(err, " error reading event")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event", 301)
			return
		}
		//Deleting event
		_, erra := event_io.DeleteEvent(event.Id)
		if erra != nil {
			fmt.Println(err, " error deleting event")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event", 301)
			return
		}

		//Deleting images
		eventImages, err := event_io.ReadEventImgOf(eventId)
		if err != nil {
			fmt.Println(err, " error reading event, this event may not have eventImages")
		} else {
			for _, eventImage := range eventImages {
				_, err := image_io.DeleteImage(eventImage.ImageId)
				if err != nil {
					fmt.Println(err, " error deleting image")
				} else {
					_, err := event_io.DeleteEventImg(eventImage.Id)
					if err != nil {
						fmt.Println(err, " error deleting image")
					}
				}
			}
		}

		//Deleting History
		eventHistory, err := event_io.ReadEventHistoryWithEventId(eventId)
		if err != nil {
			fmt.Println(err, " error reading event History, this event may not have eventHistory")
		} else {
			_, err := history_io.DeleteHistorie(eventHistory.HistoryId)
			if err != nil {
				fmt.Println(err, " error deleting histories")
			} else {
				_, err := event_io.DeleteEventHistory(eventHistory.Id)
				if err != nil {
					fmt.Println(err, " error deleting EventHistories")
				}
			}
		}
		//Deleting EVent Place
		eventPlace, err := event_io.ReadEventPlaceOf(eventId)
		if err != nil {
			fmt.Println(err, " error reading eventPlace, this event may not have a place")
		} else {
			_, err := place_io.DeletePlace(eventPlace.Id)
			if err != nil {
				fmt.Println(err, " error deleting Place")
			} else {
				_, err := event_io.DeleteEventPlace(eventPlace.Id)
				if err != nil {
					fmt.Println(err, " error deleting Event Place")
				}
			}
		}

		//Event partner
		eventPartners, err := event_io.ReadEventPartenerOf(eventId)
		if err != nil {
			fmt.Println(err, " error reading Event partner")
		} else {
			for _, eventPartner := range eventPartners {
				_, err := event_io.DeleteEventPartener(eventPartner.Id)
				if err != nil {
					fmt.Println(err, " error deleting Event partner")
				}
			}
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted the following Event: "+event.Name)
		http.Redirect(w, r, "/admin_user/event", 301)
		return

	}
}

func DeleteImagehandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageId := chi.URLParam(r, "imageId")
		eventId := chi.URLParam(r, "eventId")
		EventImageId := chi.URLParam(r, "EventImageId")

		//check eventImage before deleting
		eventImage, err := event_io.ReadEventImg(EventImageId)
		if err != nil {
			fmt.Println(err, " error event can not be find")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		//Check if the image exists
		image, err := image_io.ReadImage(imageId)
		if err != nil {
			fmt.Println(err, " error image can not be find")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}

		//Now we are deleting
		_, errx := event_io.DeleteEventImg(eventImage.Id)
		if errx != nil {
			fmt.Println(errx, " error Deleting event Image")
		} else {
			_, err := image_io.DeleteImage(image.Id)
			if err != nil {
				fmt.Println(err, " error image can not be deleted")
			}
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
		return
	}

}

func CreateEventPictureEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		eventId := r.PostFormValue("eventId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		if eventId != "" {

			eventImageObject := event2.EventImage{"", "", eventId, ""}
			eventImageHelper := event2.EventImageHelper{eventImageObject, filesByteArray}

			_, errx := event_io.CreateEventImg(eventImageHelper)
			if errx != nil {
				fmt.Println(errx, " error creating Event Image")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/event/", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
		return
	}
}

func UpdatePicturesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		file, _, err := r.FormFile("file")
		imageId := r.PostFormValue("imageId")
		eventId := r.PostFormValue("eventId")
		eventImageId := r.PostFormValue("eventImageId")
		imageType := r.PostFormValue("imageType")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file}
		filesByteArray := misc.CheckFiles(filesArray)

		//Checking fields contents
		if eventId != "" && imageId != "" && imageType != "" && eventImageId != "" {
			eventImage := event2.EventImage{eventImageId, imageId, eventId, imageType}
			eventImageHelper := event2.EventImageHelper{eventImage, filesByteArray}
			_, err := event_io.UpdateEventImg(eventImageHelper)
			if err != nil {
				fmt.Println(err, " error updating eventImage Helper")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
		return
	}
}

func AddPictureHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		eventId := r.PostFormValue("eventId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		if eventId != "" {
			eventImageObject := event2.EventImage{"", "", eventId, ""}
			eventImageHelper := event2.EventImageHelper{eventImageObject, filesByteArray}

			_, errx := event_io.CreateEventImg(eventImageHelper)
			if errx != nil {
				fmt.Println(errx, " error creating PeopleImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/event", 301)
		return
	}
}

func UpdateDetailsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		event_name := r.PostFormValue("event_name")
		eventId := r.PostFormValue("eventId")
		date := r.PostFormValue("date")
		description := r.PostFormValue("description")
		projectId := r.PostFormValue("projectId")
		eventStatus := r.PostFormValue("eventStatus")
		yearId := r.PostFormValue("year")

		_, err := event_io.ReadEvent(eventId)
		if err != nil {
			fmt.Println("error reading Event")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		//updating event Year
		if yearId != "" {
			eventYearRead, err := event_io.ReadEventYearWithEventId(eventId)
			if err != nil {
				eventYearNewObject := event2.EventYear{"", eventId, yearId}
				_, err := event_io.CreateEventYear(eventYearNewObject)
				if err != nil {
					fmt.Println("could not create a new eventYear")
				}
			} else {
				eventYearNewObject := event2.EventYear{eventYearRead.Id, eventId, yearId}
				_, err := event_io.CreateEventYear(eventYearNewObject)
				if err != nil {
					fmt.Println("could not update a new eventYear")
				}
			}
		}

		if projectId != "" {
			eventProject, err := event_io.ReadEventProjectWithEventId(eventId)
			if err != nil {
				fmt.Println("error reading Event project, this event may not had a project yet. proceeding into creating a project")
				eventProjectObject := event2.EventProject{"", projectId, eventId, ""}
				_, err := event_io.CreateEventProject(eventProjectObject)
				if err != nil {
					fmt.Println("error Creating Event project")
				}
			} else {
				eventProjectObject := event2.EventProject{eventProject.Id, projectId, eventId, eventProject.Description}
				_, err := event_io.UpdateEventProject(eventProjectObject)
				if err != nil {
					fmt.Println("error updating Event project")
				}
			}
		}

		//If the event already exist, this time we need to update.
		eventObject := event2.Event{eventId, event_name, date, eventStatus, description}
		_, errs := event_io.UpdateEvent(eventObject)
		if errs != nil {
			fmt.Println("error update Event")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}

		fmt.Println(" successfully updated")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: Event Details. ")
		http.Redirect(w, r, "/admin_user/event", 301)
		return
	}
}

func UpdatePlaceHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		description := r.PostFormValue("description")
		eventId := r.PostFormValue("eventId")
		placeId := r.PostFormValue("PlaceId")

		//Check the placeId
		_, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println("error reading Place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		if description != "" && eventId != "" {
			eventPlaceObject := event2.EventPlace{"", placeId, eventId, description}
			_, err := event_io.CreateEventPlace(eventPlaceObject)
			if err != nil {
				fmt.Println(err, " error creating event place")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updating:  Event Place. ")
			http.Redirect(w, r, "/admin_user/event", 301)
			return
		}
		fmt.Println(" successfully updated")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: Event Place. ")
		http.Redirect(w, r, "/admin_user/event", 301)
		return
	}
}

func UpdateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		eventId := r.PostFormValue("eventId")
		historyId := r.PostFormValue("historyId")

		//checking if the EventtHistory exists
		_, err := history_io.ReadHistorie(historyId)
		//fmt.Println(historyContent)
		if err != nil {
			fmt.Println(err, " could not read history")
			fmt.Println(" proceeding into creation of a history.....")
			history := history2.Histories{"", misc.ConvertToByteArray(historyContent)}

			//fmt.Println("history Object: ", history)

			newHistory, err := history_io.CreateHistorie(history)
			if err != nil {
				fmt.Println(err, " something went wrong! could not create history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
			fmt.Println("HistoryId created successfully ..")
			fmt.Println(" proceeding into creation of a event_history.....")
			eventHistory := event2.EventHistory{"", eventId, newHistory.Id}
			_, errr := event_io.CreateEventHistory(eventHistory)
			if errr != nil {
				fmt.Println(err, " could not create eventHistory")
				fmt.Println("RollBack ...")
				fmt.Println("deleting history ...")
				_, err := history_io.DeleteHistorie(newHistory.Id)
				if err != nil {
					fmt.Println("Error deleting history ...")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
				return
			}
			fmt.Println(" successfully created")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		histories := history2.Histories{historyId, misc.ConvertToByteArray(historyContent)}

		_, errr := history_io.UpdateHistorie(histories)
		if errr != nil {
			fmt.Println(err, " something went wrong! could not update history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
			return
		}
		event, errx := event_io.ReadEvent(eventId)
		if errx != nil {
			fmt.Println("error reading project")
		}
		fmt.Println(" successfully updated")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: "+event.Name+"  Event. ")
		//http.Redirect(w, r, "/admin_user/event", 301)
		http.Redirect(w, r, "/admin_user/event/edit/"+eventId, 301)
		return
	}
}

func CreateEventHistoryEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
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
		//eventId := r.PostForm["groupId"]
		groupIds := r.Form["groupId"]
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		//Creating EventHistory and HistoryId
		//fmt.Println("eventIed: ", eventId, " test>>>>", mytextarea)
		if eventId != "" && mytextarea != "" {
			//Creating Histories Object
			historyObject := history2.Histories{"", misc.ConvertToByteArray(mytextarea)}
			histories, err = history_io.CreateHistorie(historyObject)
			if err != nil {
				fmt.Println("could not create history and wont create Event history")
			} else {
				//creating Event HistoryId
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

			if groupIds != nil {
				//Creating peopleEvent
				for _, groupId := range groupIds {
					eventGroupObject := event2.EventGroup{"", eventId, groupId}
					_, err := event_io.CreateEventGroup(eventGroupObject)
					if err != nil {
						fmt.Println(err, " error when creating event group")
					}
				}

			}

		}

		//creating EVentImage
		eventImageObject := event2.EventImage{"", "", eventId, ""}
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
		http.Redirect(w, r, "/admin_user/event", 301)
		return
	}
}

func EventPicture(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
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
		groups, err := group_io.ReadGroups()
		if err != nil {
			fmt.Println(err, " error reading all the groups")
		}

		type PageData struct {
			Projects      []project2.Project
			Partners      []partner2.Partner
			Event         event2.Event
			Groups        []group.Groupes
			Backend_error string
			Unknown_error string
		}
		data := PageData{projects, partners, event, groups, backend_error, unknown_error}
		files := []string{
			//app.Path + "admin/event/new_event_picture.html",
			app.Path + "admin/event/event_history_picture.html",
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
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
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
		projects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading projects")
		}
		partners, err := partner_io.ReadPartners()
		if err != nil {
			fmt.Println(err, " error reading partener")
		}
		peoples, err := people_io.ReadPeoples()
		if err != nil {
			fmt.Println(err, " error reading peoples")
		}
		places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println(err, " error reading peoples")
		}
		years, err := io2.ReadYears()
		if err != nil {
			fmt.Println(err, " error reading all the years")
		}
		eventData := misc.GetEventDate(eventId)

		groups, err := group_io.ReadGroups()
		if err != nil {
			fmt.Println(err, " error reading all the groups")
		}
		pageFlows, err := event_io.ReadAllEventPageFlowByEventId(eventId)
		if err != nil {
			fmt.Println(err, " error reading all the pageflow")
		}
		type PageData struct {
			Event       event2.Event
			EventData   misc.EventData
			Projects    []project2.Project
			Partners    []partner2.Partner
			SidebarData misc.SidebarData
			Peoples     []people.People
			Places      []place2.Place
			Years       []museum.Years
			GroupData   []event3.GroupData
			Groups      []group.Groupes
			Comments    []comment.CommentHelper2
			Gallery     []misc.EventGalleryImages
			PageFLow    []contribution.EventPageFlow
		}
		date := PageData{event,
			eventData,
			projects,
			partners,
			misc.GetSideBarData("event", ""),
			peoples,
			places,
			years,
			event3.GetGroupsData(eventId),
			groups,
			GetEventCommentsWithEventId(eventId),
			misc.GetEventGallery(eventId),
			pageFlows}
		files := []string{
			app.Path + "admin/event/edit_event.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "base_templates/footer.html",
			app.Path + "admin/template/cards.html",
			app.Path + "admin/template/topbar.html",
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

//Not in user anymore
func NewEventsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
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

		fmt.Println("Session result: ", adminHelper.CheckAdminInSession(app, r))
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}

		var unknown_error string
		var backend_error string
		var eventList []event2.Event
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
		for _, event := range events {
			eventobject := event2.Event{event.Id, event.Name, misc.FormatDateMonth(event.Date), event.IsPast, event.Description}
			eventList = append(eventList, eventobject)
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
		places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println(err, " error reading all the places")
		}
		peoples, err := people_io.ReadPeoples()
		if err != nil {
			fmt.Println(err, " error reading all the peoples")
		}
		years, err := io2.ReadYears()
		if err != nil {
			fmt.Println(err, " error reading all the years")
		}

		type PageData struct {
			Projects      []project2.Project
			Partners      []partner2.Partner
			Backend_error string
			Unknown_error string
			Events        []event2.Event
			SidebarData   misc.SidebarData
			Places        []place2.Place
			Peoples       []people.People
			Years         []museum.Years
		}
		data := PageData{projects, partners, backend_error, unknown_error, eventList, misc.GetSideBarData("event", "event"), places, peoples, years}
		files := []string{
			app.Path + "admin/event/events.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
			app.Path + "admin/template/cards.html",
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

type YearData struct {
	Year   museum.Years
	Number int64
}

func getYearDate() []YearData {
	var year []YearData
	yearResults, err := io2.ReadYears()
	if err != nil {
		fmt.Println(err, "error reading year")
		return year
	} else {
		//get the year event
		for _, yearResult := range yearResults {
			amount, err := event_io.CountEventYearWithYearId(yearResult.Id)
			if err != nil {
				fmt.Println(err, "error reading year with yearId.")
			} else {
				year = append(year, YearData{yearResult, amount})
			}
		}
	}
	return year
}

/***
Here we create a new event
*/
func CreateEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()

		var newEvent event2.Event
		event_name := r.PostFormValue("event_name")
		date := r.PostFormValue("date")
		project := r.PostFormValue("project")
		description := r.PostFormValue("description")
		partner := r.PostFormValue("partner")
		placeId := r.PostFormValue("placeId")
		yearId := r.PostFormValue("year")
		src := r.PostFormValue("src")
		pageFlowTitle := r.PostFormValue("pageFlowTitle")
		eventStatus := r.PostFormValue("eventStatus")
		peopleIds := r.Form["peopleId"]

		//fmt.Println("peopleId: ",peopleIds)
		if event_name != "" {
			eventObject := event2.Event{"", event_name, date, eventStatus, description}
			errs := errors.New("")
			newEvent, errs = event_io.CreateEvent(eventObject)
			if errs != nil {
				fmt.Println(errs, " error when creating a new event")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/event/new", 301)
				return
			}
			if partner != "" && newEvent.Id != "" {
				eventPartner := event2.EventPartener{"", partner, newEvent.Id, ""}
				_, err := event_io.CreateEventPartener(eventPartner)
				if err != nil {
					fmt.Println(err, " error when creating event partner")
				}
			}
			//Creation of pageFlow
			if src != "" && pageFlowTitle != "" {
				eventPageFlowObject := contribution.EventPageFlow{"", newEvent.Id, pageFlowTitle, src}
				_, err := event_io.CreateEventPageFlow(eventPageFlowObject)
				if err != nil {
					fmt.Println(err, " error when creating eventPageFLow")
				}
			}
			//Creating event year
			if yearId != "" {
				//fmt.Println("eventStatus: ",eventStatus)
				eventyearObject := event2.EventYear{"", newEvent.Id, yearId}
				_, err := event_io.CreateEventYear(eventyearObject)
				if err != nil {
					fmt.Println(err, " error when creating event year")
				}
			}

			//TODO will need to create EventProject description Field in HTML.
			if project != "" && newEvent.Id != "" {
				eventProject := event2.EventProject{"", project, newEvent.Id, ""}
				_, err := event_io.CreateEventProject(eventProject)
				if err != nil {
					fmt.Println(err, " error when creating event project")
				}
			}
			//Creating peopleEvent
			if peopleIds != nil {
				for _, people := range peopleIds {
					eventPeopleObject := event2.EventPeople{"", newEvent.Id, people}
					_, err := event_io.CreateEventPeople(eventPeopleObject)
					if err != nil {
						fmt.Println(err, " error when creating event people")
					}
				}

			}

			//TODO should create place description Field
			eventPlace := event2.EventPlace{"", placeId, newEvent.Id, ""}
			_, err := event_io.CreateEventPlace(eventPlace)
			if err != nil {
				fmt.Println(err, " error when creating Event place")
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
		http.Redirect(w, r, "/admin_user/event/new", 301)
		return
	}
}

func UpdateEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		id := r.PostFormValue("id")
		name := r.PostFormValue("name")
		description := r.PostFormValue("description")
		date := r.PostFormValue("date")
		year := r.PostFormValue("year")
		eventStatus := r.PostFormValue("eventStatus")

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

		//checking if this event has been associated to a year
		enventYear, errx := event_io.ReadEventYearWithEventId(id)
		if errx != nil {
			//Creating event year
			if year != "" {
				//fmt.Println("eventStatus: ",eventStatus)
				eventyearObject := event2.EventYear{"", id, year}
				_, err := event_io.CreateEventYear(eventyearObject)
				if err != nil {
					fmt.Println(err, " error when creating event year")
				}
			}

		} else { // if this event has been already associated to a year now we need to update.
			eventyearObject := event2.EventYear{"", enventYear.EventId, year}
			fmt.Println("updating event year :", eventyearObject)
			_, err := event_io.UpdateEventYear(eventyearObject)
			if err != nil {
				fmt.Println(err, " error when creating event year")
			}
		}

		//we checking if there is a need of updating
		if event.Name != name && event.Id != id && event.Date != date {
			event := event2.Event{id, name, date, eventStatus, description}
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
