package peoples

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/constates"
	"ostmfe/controller/generic"
	"ostmfe/controller/misc"
	people3 "ostmfe/controller/people"
	"ostmfe/domain/comment"
	"ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	people2 "ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
)

func PeopleHome(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", PeopleHandler(app))
	r.Get("/new", NewPeopleHandler(app))
	r.Get("/new_stp2/{peopleId}", NewPeoplestp2Handler(app))
	r.Get("/edit/{peopleId}", EditPeopleHandler(app))
	r.Get("/delete/{peopleId}", DeletePeopleHandler(app))
	r.Get("/delete-category/{category}", DeletePeopleCategoryHandler(app))

	r.Post("/create_stp1", CreatePeopleHandler(app))
	//r.Post("/create_stp2", CreatePeopleStp2Handler(app))
	r.Post("/create_image", CreatePeopleImageHandler(app))

	r.Post("/update_image", UpdatePeopleImageHandler(app))
	r.Post("/update_details", UpdatePeopleDetailHandler(app))
	r.Post("/update_history", UpdatePeopleHistoryHandler(app))
	r.Post("/create_history", CreatePeopleHistoryHandler(app))
	r.Post("/add_pictures", AddPeopleImageHandler(app))

	r.Post("/add_place", AddPlaceHandler(app))
	r.Post("/add_event", AddEventHandler(app))
	//Gallery
	r.Post("/create-gallery", createPeopleGaller(app))
	r.Get("/delete-gallery/{pictureId}/{peopleId}/{peopleGalleryPictureId}", DeleteGalleryHandler(app))
	r.Get("/activate_comment/{commentId}/{peopleId}", ActivateCommentHandler(app))
	r.Get("/delete_people/{eventPeopleId}/{peopleId}", DeleteEventPeopleHandler(app))

	r.Get("/people_category/new", NewPeopleCategoryHandler(app))
	r.Post("/people_category/create", CreatePeopleCategoryHandler(app))

	r.Post("/create-descriptive-Image", createDescriptiveImage(app))

	return r
}

func createDescriptiveImage(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		peopleId := r.PostFormValue("peopleId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if peopleId != "" && description != "" {
			ImageObject := image.Images{"", content, description}
			imageObj, err := image_io.CreateImage(ImageObject)
			if err != nil {
				fmt.Println(err, " error creating image")
			} else {
				ImageType, err := image_io.ReadImageTypeWithName(constates.DESCRIPTIVE)
				if err != nil {
					fmt.Println(err, " error Reading image Type")
				}
				peopleImageObject := people2.PeopleImage{"", peopleId, imageObj.Id, ImageType.Id}
				_, errx := people_io.CreatePeopleImageHere(peopleImageObject)
				if errx != nil {
					fmt.Println(err, " error creating GroupGallery")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
		return
	}
}
func DeleteEventPeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventPeopleId := chi.URLParam(r, "eventPeopleId")
		peopleId := chi.URLParam(r, "peopleId")
		_, err := event_io.DeleteEventPeople(eventPeopleId)
		if err != nil {
			app.ErrorLog.Println(err, " error when deleting event people.")
		}
		http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
		return
	}
}
func ActivateCommentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentId := chi.URLParam(r, "commentId")
		peopleId := chi.URLParam(r, "peopleId")
		result := misc.ActivateComment(commentId)
		fmt.Print("Activation Result: ", result)
		http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
		return
	}
}
func CreatePeopleHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		peopleId := r.PostFormValue("peopleId")
		myArea := r.PostFormValue("myArea")
		if myArea != "" && peopleId != "" {
			historieObject := history2.Histories{"", misc.ConvertToByteArray(myArea)}
			historie, err := history_io.CreateHistorie(historieObject)
			if err != nil {
				fmt.Println(err, "updating histories")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			people, err := people_io.ReadPeople(peopleId)
			if err != nil {
				fmt.Println(err, " error reading people")
				fmt.Println("Deleting history ....")
				_, err = history_io.DeleteHistorie(historie.Id)
				if err != nil {
					fmt.Println(err, " can't delete peopleHistory")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}

			updatedPeople := people2.People{people.Id, people.Name, people.Surname, people.BirthDate, people.DeathDate, people.Origin, people.Profession, people.Brief, historie.Id}
			_, err = people_io.UpdatePeople(updatedPeople)
			if err != nil {
				fmt.Println(err, " error updating people")
				fmt.Println("Deleting history ....")
				_, err = history_io.DeleteHistorie(historie.Id)
				if err != nil {
					fmt.Println(err, " can't delete peopleHistory")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			fmt.Println("update done successfully.")
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People HistoryId of : ")
			http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people", 301)
		return
	}
}
func AddEventHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		peopleId := r.PostFormValue("peopleId")
		eventId := r.PostFormValue("eventId")
		if eventId != "" && peopleId != "" {
			peoplePlaceObject := event.EventPeople{"", eventId, peopleId}
			_, err := event_io.CreateEventPeople(peoplePlaceObject)
			if err != nil {
				fmt.Println(err, " error creating Event People")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people", 301)
	}
}
func DeletePeopleCategoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := chi.URLParam(r, "category")
		//fmt.Println("categoryId: ",category)

		_, err := people_io.DeleteCategory(category)
		if err != nil {
			fmt.Println("error deleting Category")
		}
		http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
		return
	}
}
func DeleteGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pictureId := chi.URLParam(r, "pictureId")
		peopleId := chi.URLParam(r, "peopleId")
		peopleGalleryPictureId := chi.URLParam(r, "peopleGalleryPictureId")

		//Deleting project
		gallery, err := image_io.DeleteGalery(pictureId)
		if err != nil {
			fmt.Println("error deleting gallery")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
			return
		} else {
			_, err := people_io.DeletePeopleGalery(peopleGalleryPictureId)
			if err != nil {
				fmt.Println("error deleting people gallery")
				fmt.Println("ROLLING BACK!!!")
				_, err := image_io.UpdateGallery(gallery)
				if err != nil {
					fmt.Println("error updating gallery")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: people Gallery. ")
		http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
		return
	}
}
func createPeopleGaller(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		peopleId := r.PostFormValue("peopleId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if peopleId != "" && description != "" {
			galery := image.Gallery{"", content, description}
			galleryObject, err := image_io.CreateGalery(galery)
			if err != nil {
				fmt.Println(err, " error creating gallery")
			} else {
				peopleGallery := people2.PeopleGalery{"", peopleId, galleryObject.Id}
				_, err := people_io.CreatePeopleGalery(peopleGallery)
				if err != nil {
					fmt.Println(err, " error creating GroupGallery")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
		return
	}
}
func AddPlaceHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		peopleId := r.PostFormValue("peopleId")
		placeId := r.PostFormValue("placeId")
		if placeId != "" && peopleId != "" {
			peoplePlaceObejct := people2.PeoplePlace{"", placeId, peopleId}
			_, err := people_io.CreatePeoplePlace(peoplePlaceObejct)
			if err != nil {
				fmt.Println(err, " error creating PeoplePlace")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people", 301)
	}
}

func AddPeopleImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		peopleId := r.PostFormValue("peopleId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		reader := bufio.NewReader(file)
		content, _ := ioutil.ReadAll(reader)

		if peopleId != "" {

			image := image.Images{"", content, ""}
			peopleImageObject := people2.PeopleImage{"", peopleId, image.Id, ""}
			_, err := people_io.CreatePeopleImage(peopleImageObject)
			if err != nil {
				fmt.Println(err, " error creating PeopleImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			//http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
			http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		//http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
		http.Redirect(w, r, "/admin_user/people", 301)
		return
	}
}

//TODO to be removed
/****
This method is requested when a people was created without an image now this method will help to create one.
*/
func CreatePeopleImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		peopleId := r.PostFormValue("peopleId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>")
		}
		reader := bufio.NewReader(file)
		content, _ := ioutil.ReadAll(reader)

		if peopleId != "" {
			image := image.Images{"", content, ""}
			imageObject, err := image_io.CreateImage(image)
			if err != nil {
				fmt.Println(err, "<<<<<< error reading file>>>>")
			} else {
				peopleImageObject := people2.PeopleImage{"", peopleId, imageObject.Id, ""}
				_, err := people_io.CreatePeopleImage(peopleImageObject)
				if err != nil {
					fmt.Println(err, " error creating PeopleImage")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
			}

			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/people/", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people", 301)
		return
	}

}

func UpdatePeopleHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		historyId := r.PostFormValue("historyId")
		peopleId := r.PostFormValue("peopleId")
		myArea := r.PostFormValue("myArea")
		if myArea != "" && historyId != "" && peopleId != "" {
			historieObject := history2.Histories{historyId, misc.ConvertToByteArray(myArea)}
			_, err := history_io.UpdateHistorie(historieObject)
			if err != nil {
				fmt.Println(err, "updating people histories")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People HistoryId of : ")
			http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people", 301)
		return
	}
}

func UpdatePeopleDetailHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("name")
		peopleId := r.PostFormValue("peopleId")
		surname := r.PostFormValue("surname")
		b_date := r.PostFormValue("b_date")
		d_date := r.PostFormValue("d_date")
		profession := r.PostFormValue("profession")
		origin := r.PostFormValue("origin")
		brief := r.PostFormValue("brief")
		categoryId := r.PostFormValue("categoryId")
		historyId := r.PostFormValue("historyId")
		if name != "" && peopleId != "" && surname != "" {
			//TODO need to learn how to check if a data time if nil or empty....
			peopleObejct := people2.People{peopleId, name, surname, b_date, d_date, origin, profession, brief, historyId}
			people, err := people_io.UpdatePeople(peopleObejct)
			if err != nil {
				fmt.Println(err, "updating people details")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if categoryId != "" {
				peopleCategory := people2.PeopleCategory{"", categoryId, peopleId, ""}
				_, err := people_io.CreatePeopleCategory(peopleCategory)
				if err != nil {
					fmt.Println(err, " error creating people category")
				}
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Details of : "+people.Name)
			http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
			return
		}

		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
		return
	}
}

func UpdatePeopleImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		var content []byte
		file, _, err := r.FormFile("file")
		imageId := r.PostFormValue("imageId")
		peopleId := r.PostFormValue("peopleId")
		peopleImageId := r.PostFormValue("peopleImageId")
		imageType := r.PostFormValue("imageType")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}

		if peopleId != "" && imageId != "" && imageType != "" && peopleImageId != "" {
			imageObject := image.Images{imageId, content, ""}
			_, err := image_io.UpdateImage(imageObject)
			if err != nil {
				fmt.Println(err, " error updating peopleImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated a People Picture : ")
			http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/edit/"+peopleId, 301)
		return
	}
}
func CreatePeopleCategoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		category := r.PostFormValue("category")
		fmt.Println(category)
		if category != "" {
			people := people2.Category{"", category}
			peopleCategory, err := people_io.CreateCategory(people)
			if err != nil {
				fmt.Println(err, " error creating people Category")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
				return
			}

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new People Type : "+peopleCategory.Category)
			http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people/people_category/new", 301)
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

		peoples, err := people_io.ReadCategories()
		if err != nil {
			fmt.Println(err, " There is an error when reading all the people category")
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PageData struct {
			Peoples       []people2.Category
			Backend_error string
			Unknown_error string
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}
		data := PageData{peoples, backend_error, unknown_error,
			misc.GetSideBarData("people", "people_category"), adminName, adminImage}
		files := []string{
			app.Path + "admin/people/peopleCategory.html",
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

		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		peoples := misc.GetPeopleWithStringdate()
		type PageData struct {
			People        people2.People
			Backend_error string
			Unknown_error string
			Peoples       []misc.PeopleWithStringdate
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}
		data := PageData{people, backend_error, unknown_error, peoples,
			misc.GetSideBarData("people", "people"), adminName, adminImage}
		files := []string{
			app.Path + "admin/people/people_step2.html",
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

func CreatePeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		var content []byte
		file, _, err := r.FormFile("file")
		placeId := r.PostFormValue("placeId")
		name := r.PostFormValue("name")
		surname := r.PostFormValue("surname")
		profession := r.PostFormValue("profession")
		b_date := r.PostFormValue("b_date")
		d_date := r.PostFormValue("d_date")
		description := r.PostFormValue("description")
		origin := r.PostFormValue("origin")
		brief := r.PostFormValue("brief")
		history := r.PostFormValue("history")
		categoryId := r.PostFormValue("categoryId")

		var newObjectHistory history2.Histories

		if err != nil {
			fmt.Println(err, "<<<error reading file>>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}

		if history != "" {
			//History
			historyObejct := history2.Histories{"", misc.ConvertToByteArray(history)}
			newObjectHistory, err = history_io.CreateHistorie(historyObejct)
			if err != nil {
				fmt.Println(err, " error creating a new history")
			}
		}

		fmt.Println("death date: ", d_date)
		//If the people is already dead, his date of birth and his death date are different otherwise both date will be save with the same values.
		var deathDate string
		if d_date == "" {
			deathDate = b_date
		} else {
			deathDate = d_date
		}
		if name != "" && surname != "" {
			peopleObject := people2.People{"", name, surname, b_date, deathDate, origin, profession, brief, newObjectHistory.Id}
			people, err := people_io.CreatePeople(peopleObject)
			if err != nil {
				fmt.Println(err, " error creating a new people")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/people", 301)
				return
			}
			//Creating people place
			if placeId != "" {
				peoplePlaceObject := people2.PeoplePlace{"", placeId, people.Id}
				_, err := people_io.CreatePeoplePlace(peoplePlaceObject)
				if err != nil {
					fmt.Println("error creating peoplePlace")
				}
			}
			//People Category
			if categoryId != "" {
				peopleCategoryObject := people2.PeopleCategory{"", categoryId, people.Id, ""}
				_, err := people_io.CreatePeopleCategory(peopleCategoryObject)
				if err != nil {
					fmt.Println("error creating people Category")
				}
			}

			//Image
			imageObject := image.Images{"", content, description}
			imageObjectNew, err := image_io.CreateImage(imageObject)
			if err != nil {
				fmt.Println(err, " error creating a new image")
			}
			peopleImageObject := people2.PeopleImage{"", people.Id, imageObjectNew.Id, generic.GetImageTypeId(constates.PROFILE)}
			_, errx := people_io.CreatePeopleImageX(peopleImageObject)
			if errx != nil {
				fmt.Println(errx, " error creating a new placeImage")
			}
			if people.Id != "" {
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+people.Name)
				http.Redirect(w, r, "/admin_user/people", 301)
				return
			}
		}

		fmt.Println("One of the field is missing or newPlace.Id is nil")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/people", 301)
		return
	}
}

func EditPeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		peopleId := chi.URLParam(r, "peopleId")
		people, err := people_io.ReadPeople(peopleId)
		if err != nil {
			fmt.Println(err, "error reading people for the following people id: ", peopleId)
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/people", 301)
			return
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
		places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println(err, " error reading places")
		}
		peoplePlaces := people3.GetPeoplePlace(peopleId)

		// Getting this method from client people_controller. it gives me a list of places that are linked to a people.
		peopleEditable, err := people_io.GetAggregatedPeople(people.Id)
		if err != nil {
			fmt.Println(err, " error reading events")
		}
		//app.InfoLog.Println(peopleEditable.ProfileImage)

		events, err := event_io.ReadEvents()
		if err != nil {
			fmt.Println(err, " error reading events")
		}
		category, err := GetPeopleCategory(peopleId)
		if err != nil {
			fmt.Println(err, " error reading category")
		}

		Categories, err := people_io.ReadCategories()
		if err != nil {
			fmt.Println(err, " error reading categories")
		}
		commentNumber, pendingcomments, activeComments := peopleCommentCalculation(peopleId)
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PageDate struct {
			PeopleDetails   people2.People
			People          people2.PeopleAggregate
			SidebarData     misc.SidebarData
			PeoplePlace     []place2.Place
			Places          []place2.Place
			Events          []event.Event
			Comments        []comment.CommentHelper2
			Gallery         []misc.PeopleGalleryImages
			CommentNumber   int64
			PendingComments int64
			ActiveComments  int64
			Backend_error   string
			Unknown_error   string
			EventPeople     []EventPeopleData
			Categories      []people2.Category
			Category        people2.Category
			AdminName       string
			AdminImage      string
		}
		data := PageDate{people,
			peopleEditable,
			misc.GetSideBarData("people", "people"),
			peoplePlaces,
			places,
			events,
			misc.GetPeopleComments(peopleId),
			misc.GetPeopleGallery(peopleId),
			commentNumber,
			pendingcomments,
			activeComments,
			backend_error,
			unknown_error,
			GetPeopleEvents(peopleId),
			Categories,
			category,
			adminName,
			adminImage,
		}

		files := []string{
			app.Path + "admin/people/edit_people.html",
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
		peoples, err := people_io.ReadPeoples()
		if err != nil {
			fmt.Println(err, "error reading peoples")
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			Peoples       []people2.People
		}
		data := PagePage{backend_error, unknown_error, peoples}
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

		peoples := misc.GetPeopleWithStringdate()
		places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println("error reading Places")
		}
		categories, err := people_io.ReadCategories()
		if err != nil {
			fmt.Println("error reading Categories")
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PagePage struct {
			Places        []place2.Place
			Categories    []people2.Category
			Backend_error string
			Unknown_error string
			Peoples       []misc.PeopleWithStringdate
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}
		data := PagePage{places, categories, backend_error, unknown_error,
			peoples, misc.GetSideBarData("people", "people"), adminName, adminImage}
		files := []string{
			app.Path + "admin/people/people.html",
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

func DeletePeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peopleId := chi.URLParam(r, "peopleId")
		people, err := people_io.ReadPeople(peopleId)
		if err != nil {
			fmt.Println(err, "error reading people for the following people id: ", peopleId)
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/people", 301)
			return
		}
		//Deleting people
		_, err = people_io.DeletePeople(peopleId)
		if err != nil {
			fmt.Println(err, "error reading people for the following people id: ", peopleId)
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/people", 301)
			return
		}
		//HistoryId
		peopleHistory, err := people_io.ReadPeopleHistoryWithPplId(peopleId)
		if err != nil {
			fmt.Println(err, " error reading peopleHistory")
		} else {
			_, err := people_io.DeletePeopleHistory(peopleHistory.Id)
			if err != nil {
				fmt.Println(err, " error deleting peopleHistory")
			}
			_, errx := history_io.DeleteHistorie(peopleHistory.HistoryId)
			if errx != nil {
				fmt.Println(errx, " error deleting peopleHistory")
			}
		}
		//Image
		peopleImages, err := people_io.ReadPeopleImageWithPeopleId(peopleId)
		if err != nil {
			fmt.Println(err, " error reading peopleImage")
		} else {
			_, err := people_io.DeletePeopleImage(peopleImages.Id)
			if err != nil {
				fmt.Println(err, " error deleting peopleImage")
			}
			_, errs := image_io.DeleteImage(peopleImages.ImageId)
			if errs != nil {
				fmt.Println(errs, " error deleting peopleImage")
			}
		}
		if people.Id != "" {
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted : "+people.Name)
			http.Redirect(w, r, "/admin_user/people", 301)
			return
		}
		//app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted : "+people.Name)
		http.Redirect(w, r, "/admin_user/people", 301)
		return

	}
}
