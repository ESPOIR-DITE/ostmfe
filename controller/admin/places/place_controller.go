package places

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/domain/contribution"
	"ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	image2 "ostmfe/domain/image"
	"ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	"ostmfe/io/comment_io"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
	"time"
)

func PlaceHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", PlacesHandler(app))
	r.Get("/new", NewPlacesHandler(app))
	r.Get("/edit/{placeId}", EditPlacesHandler(app))
	r.Get("/delete/{placeId}", DeletePlacesHandler(app))
	r.Post("/create_stp1", CreateStp1Handler(app))
	//r.Post("/create_stp2", CreatePlaceStp2Handler(app))
	//r.Get("/new_stp2/{placeId}", NewPlaceStp2Handler(app))
	r.Get("/delete_image/{imageId}/{placeId}", DeleteImageHandler(app))

	//r.Post("/create_image", CreatePlaceImage(app))
	r.Post("/create_history", CreateHistoryHandler(app))
	r.Post("/create_image", CreateImageHandler(app))

	r.Post("/update_pictures", UpdatePictureHandler(app))
	r.Post("/update_details", UpdateDetailsHandler(app))
	r.Post("/update_history", UpdateHistoryHandler(app))
	r.Post("/create-gallery", CreatePlaceGalleryHandler(app))
	r.Post("/create-people", CreatePlacePeopleHandler(app))
	r.Post("/create-page-flow", CreatePageFlowHandler(app))

	r.Get("/delete-gallery/{pictureId}/{placeId}/{PlaceGalleryId}", DeleteGalleryHandler(app))
	r.Get("/delete-pageFlow/{placePageFlowId}/{placeId}", DeletePageFlowHandler(app))
	r.Get("/delete-people/{peoplePlaceId}/{placeId}", DeletePeopleHandler(app))

	r.Get("/activate_comment/{commentId}/{placeId}", ActivateCommentHandler(app))

	return r
}

func DeletePageFlowHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		placePageFlowId := chi.URLParam(r, "placePageFlowId")
		placeId := chi.URLParam(r, "placeId")

		//Deleting project
		_, err := place_io.DeletePlacePageFlow(placePageFlowId)
		if err != nil {
			fmt.Println("error deleting Place Page FLow")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: Project Gallery. ")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func CreateImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		PlaceId := r.PostFormValue("placeId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		if PlaceId != "" {
			reader := bufio.NewReader(file)
			content, _ := ioutil.ReadAll(reader)

			imageObject := image2.Images{"", content, "created on: " + misc.FormatDateTime(time.Now())}
			imageRetun, errr := image_io.CreateImage(imageObject)
			if errr != nil {
				fmt.Println(errr, " error updating placeImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
				return
			}
			placeTypeObject, err := image_io.ReadImageTypeWithName("Profile")
			if err != nil {
				fmt.Println(err, " error reading ImageType by Name")
			}

			//fmt.Println("image Object :",imageRetun.Id," description :",imageRetun.Description)
			placeImageObject := place2.PlaceImage{"", PlaceId, imageRetun.Id, placeTypeObject.Id}
			_, err = place_io.CreatePlaceImage(placeImageObject)
			if err != nil {
				fmt.Println(err, " error creating place Image")
			}

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully created a Place ")
			http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
			return
		}
		fmt.Println("one field is empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
		return

	}
}

func CreatePageFlowHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		placeId := r.PostFormValue("placeId")
		pageFlowTitle := r.PostFormValue("pageFlowTitle")
		scr := r.PostFormValue("scr")

		if scr != "" && placeId != "" && pageFlowTitle != "" {
			_, err := place_io.CreatePlacePageFlow(place2.PlacePageFlow{"", placeId, pageFlowTitle, scr})
			if err != nil {
				fmt.Println(err, " error creating page flow!")
			}
		}
		fmt.Print("")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func ActivateCommentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentId := chi.URLParam(r, "commentId")
		placeId := chi.URLParam(r, "placeId")
		result := misc.ActivateComment(commentId)
		fmt.Print("Activation Result: ", result)
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func DeletePeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peoplePlaceId := chi.URLParam(r, "peoplePlaceId")
		placeId := chi.URLParam(r, "placeId")

		if peoplePlaceId != "" {
			_, err := people_io.DeletePeoplePlace(peoplePlaceId)
			if err != nil {
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		fmt.Println("missing fields!")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func CreatePlacePeopleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		peopleId := r.PostFormValue("peopleId")
		placeId := r.PostFormValue("placeId")

		fmt.Println("peopleId: ", peopleId, "  placeId: ", placeId)
		if peopleId != "" && placeId != "" {
			placePeopleObject := people.PeoplePlace{"", placeId, peopleId}
			_, err := people_io.CreatePeoplePlace(placePeopleObject)
			if err != nil {
				fmt.Println(err, " error creating placePeople!")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		fmt.Println("missing fields!")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func DeleteGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pictureId := chi.URLParam(r, "pictureId")
		placeId := chi.URLParam(r, "placeId")
		PlaceGalleryId := chi.URLParam(r, "PlaceGalleryId")

		//Deleting project
		gallery, err := image_io.DeleteGalery(pictureId)
		if err != nil {
			fmt.Println("error deleting gallery")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		} else {
			_, err := place_io.DeletePlaceGalery(PlaceGalleryId)
			if err != nil {

				fmt.Println("error deleting place gallery")
				fmt.Println("ROLLING BACK!!!")
				_, err := image_io.UpdateGallery(gallery)
				if err != nil {
					fmt.Println("error updating gallery")
				}

				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
				return
			}
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: Project Gallery. ")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func CreatePlaceGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		placeId := r.PostFormValue("placeId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if placeId != "" && description != "" {
			galery := image2.Gallery{"", content, description}
			galleryObject, err := image_io.CreateGalery(galery)
			if err != nil {
				fmt.Println(err, " error creating gallery")
			} else {
				placeGallery := place2.PlaceGallery{"", placeId, galleryObject.Id}
				_, err := place_io.CreatePlaceGalery(placeGallery)
				if err != nil {
					fmt.Println(err, " error creating GroupGallery")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func DeleteImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageId := chi.URLParam(r, "imageId")
		placeId := chi.URLParam(r, "placeId")
		fmt.Println(imageId, " imageId||placeId ", placeId)
		_, err := image_io.ReadImage(imageId)
		if err != nil {
			fmt.Println(err, " error reading image")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		//If we can read the image, than we can delete it.
		_, err = image_io.DeleteImage(imageId)
		if err != nil {
			fmt.Println(err, " error deleting image")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		//if image is deleted now we can delete place image
		placeImage, err := place_io.ReadPlaceImageWithImageId(imageId)
		if err != nil {
			fmt.Println(err, " error reading Place image")
		} else {
			_, err := place_io.DeletePlaceImage(placeImage.Id)
			if err != nil {
				fmt.Println(err, " error reading Place image")
			}
		}

		fmt.Println(" successfully deleted")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an image")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return

	}
}

func UpdateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		placeId := r.PostFormValue("placeId")
		historyId := r.PostFormValue("historyId")
		//checking if the projectHistory exists
		_, err := history_io.ReadHistorie(historyId)
		fmt.Println(historyContent)
		if err != nil {
			fmt.Println(err, " something went wrong! could not read history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
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
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		place, errx := place_io.ReadPlace(placeId)
		if errx != nil {
			fmt.Println("error reading project")
		}
		fmt.Println(" successfully updated")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: "+place.Title+"  project. ")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return
	}
}

func CreateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		PlaceId := r.PostFormValue("PlaceId")
		//checking if there is contents in the variables
		if historyContent != "" && PlaceId != "" {
			history := history2.Histories{"", misc.ConvertToByteArray(historyContent)}

			newHistory, err := history_io.CreateHistorie(history)
			if err != nil {
				fmt.Println(err, " something went wrong! could not create history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
				return
			}
			fmt.Println("HistoryId created successfully ..")
			fmt.Println(" proceeding into creation of a place_history.....")
			placeHistory := place2.PlaceHistories{"", PlaceId, newHistory.Id}
			_, errr := place_io.CreatePlaceHistpory(placeHistory)
			if errr != nil {
				fmt.Println(err, " could not create ProjectHistory")
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
				http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
				return
			}
			fmt.Println(" successfully created")
			http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
			return
		}
		fmt.Println("one or more fields missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
		return

	}
}

func UpdateDetailsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		placeTitle := r.PostFormValue("placeTitle")
		placeId := r.PostFormValue("placeId")
		placeCategory := r.PostFormValue("category")
		description := r.PostFormValue("description")

		place, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println("error reading place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		app.InfoLog.Println("category: ", placeCategory)
		if placeCategory != "" {
			placeTypeObject := place2.PlaceType{placeId, placeCategory}
			_, err = place_io.CreatePlaceType(placeTypeObject)
			if err != nil {
				fmt.Println(err, " error when creating placeType")
			}
		}

		//TODO latutude and longitude should come from the fields
		if placeTitle != "" && description != "" {
			placeObject := place2.Place{placeId, placeTitle, place.Latitude, place.Longitude, description}
			_, errs := place_io.UpdatePlace(placeObject)
			if errs != nil {
				fmt.Println("error updating place")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
				return
			}

			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: Project Details. ")
			http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
			return
		}
		fmt.Println(" error creating project One field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+placeId, 301)
		return

	}
}

func UpdatePictureHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		PlaceId := r.PostFormValue("PlaceId")
		imageId := r.PostFormValue("imageId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		place, err := place_io.ReadPlace(PlaceId)
		if err != nil {
			fmt.Println(err, " could not read place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/place", 301)
			return
		}
		if imageId != "" {
			reader := bufio.NewReader(file)
			content, _ := ioutil.ReadAll(reader)

			imageObject := image2.Images{imageId, content, "updated on: " + misc.FormatDateTime(time.Now())}
			_, errr := image_io.UpdateImage(imageObject)
			if errr != nil {
				fmt.Println(errr, " error updating placeImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following Place : "+place.Title)
			http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
			return
		}
		fmt.Println("one field is empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/place/edit/"+PlaceId, 301)
		return

	}
}

func DeletePlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		placeId := chi.URLParam(r, "placeId")

		//Reading the project
		place, err := place_io.ReadPlace(placeId)
		if err != nil {
			fmt.Println("error reading place")
			fmt.Println(err, " error creating place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/place", 301)
			return
		}
		//Deleting project
		_, err = place_io.DeletePlace(place.Id)
		if err != nil {
			fmt.Println("error deleting place")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/place", 301)
			return
		}

		//Reading Project Image
		placeImages, err := place_io.ReadPlaceImageAllOf(placeId)
		if err != nil {
			fmt.Println("error Read Place Image. this project may not have an image")
		} else {
			for _, placeImage := range placeImages {
				_, err = place_io.DeletePlaceImage(placeImage.Id)
				if err != nil {
					fmt.Println("error deleting place Image")
				} else {
					_, err = image_io.DeleteImage(placeImage.ImageId)
					if err != nil {
						fmt.Println("error deleting image")
					}
				}
			}

		}

		//Reading History
		placeHistory, err := place_io.ReadPlaceHistporyOf(placeId)
		if err != nil {
			fmt.Println("error reading placeHistory. this place may not have History")
		} else {
			_, err = history_io.DeleteHistorie(placeHistory.HistoryId)
			if err != nil {
				fmt.Println("error deleting history. this project may not have an history")
			}
			_, err := place_io.DeletePlaceHistpory(placeHistory.Id)
			if err != nil {
				fmt.Println("error deleting projectHistory")
			}
		}

		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: Project Details. ")
		http.Redirect(w, r, "/admin_user/place", 301)
		return
	}
}

func EditPlacesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		var placeCategory place2.PlaceCategory
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
		placeId := chi.URLParam(r, "placeId")
		placeDate := GetPlaceEditable(placeId)

		peoples, err := people_io.ReadPeoples()
		if err != nil {
			fmt.Println(err, " error reading peoples")
		}

		Events, err := event_io.ReadEvents()
		if err != nil {
			fmt.Println(err, " error reading events")
		}
		pageFlows, err := place_io.ReadAllPlacePageFlowByPlaceId(placeId)
		if err != nil {
			fmt.Println(err, " error reading placePageFlow")
		}
		placeCategory, err = misc.GetPlaceCategory(placeId)
		if err != nil {
			fmt.Println(err, " error reading placeCategory")
		}

		placeCategories, err := place_io.ReadPlaceCategories()
		if err != nil {
			app.InfoLog.Println("error reading Places Category: ", err)
		}
		commentNumber, pendingcomments, activeComments := placeCommentCalculation(placeId)
		type PageData struct {
			PlaceData       PlaceDataEditable
			SidebarData     misc.SidebarData
			Comments        []comment.CommentHelper2
			Gallery         []misc.PlaceGalleryImages
			Peoples         []people.People
			AllPeople       []people.People
			Events          []event.Event
			AllEvents       []event.Event
			CommentNumber   int64
			PendingComments int64
			ActiveComments  int64
			PageFlows       []place2.PlacePageFlow
			Backend_error   string
			Unknown_error   string
			PlaceCategories []place2.PlaceCategory
			PlaceCategory   place2.PlaceCategory
		}
		data := PageData{placeDate, misc.GetSideBarData("place", ""),
			GetPlaceCommentsWithEventId(placeId), misc.GetPlaceGallery(placeId),
			GetAllPeoplePlace(placeId), peoples, Events, GetEventPlace(placeId),
			commentNumber,
			pendingcomments,
			activeComments,
			pageFlows,
			backend_error,
			unknown_error,
			placeCategories,
			placeCategory,
		}
		files := []string{
			app.Path + "admin/place/new_edit_place.html",
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
		err = ts.Execute(w, data)
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

//With peopleId, you get the commentNumber, pending, active.
func placeCommentCalculation(placeId string) (commentNumber int64, pending int64, active int64) {
	var commentNumbers int64 = 0
	var pendings int64 = 0
	var actives int64 = 0
	placeComments, err := comment_io.ReadAllCommentPlace(placeId)
	if err != nil {
		fmt.Println(err, " error reading People comment")
		return commentNumbers, pendings, actives
	} else {
		for _, placeComment := range placeComments {
			comments, err := comment_io.ReadComment(placeComment.CommentId)
			if err != nil {
				fmt.Println(err, " error reading comment")
			} else {
				if comments.Stat == true {
					actives++
				} else {
					pending++
				}
				commentNumber++
			}
		}
	}
	return commentNumbers, pendings, actives
}

func PlacesHandler(app *config.Env) http.HandlerFunc {
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
		places, err := place_io.ReadPlaces()
		if err != nil {
			app.InfoLog.Println("error reading Places: ", err)
		}
		placeCategories, err := place_io.ReadPlaceCategories()
		if err != nil {
			app.InfoLog.Println("error reading Places Category: ", err)
		}
		type PageData struct {
			Backend_error   string
			Unknown_error   string
			Places          []place2.Place
			SidebarData     misc.SidebarData
			PlaceCategories []place2.PlaceCategory
		}
		data := PageData{backend_error,
			unknown_error,
			places,
			misc.GetSideBarData("place", ""),
			placeCategories,
		}
		files := []string{
			app.Path + "admin/place/places.html",
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

func CreateStp1Handler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var content []byte
		file, _, err := r.FormFile("file")
		title := r.PostFormValue("title")
		category := r.PostFormValue("category")
		latlng := r.PostFormValue("latlng")
		description := r.PostFormValue("description")
		history := r.PostFormValue("history")
		src := r.PostFormValue("src")
		pageFlowTitle := r.PostFormValue("pageFlowTitle")

		if err != nil {
			fmt.Println(err, "<<<error reading file>>>>This error may happen if there is no picture selected>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}

		fmt.Println(title, "<<<title|| latlng>>>>", latlng, "  description>>>", description)
		if title != "" && latlng != "" && category != "" {
			latitude, longitude := misc.SeparateLatLng(latlng)
			fmt.Println("latitude: ", latitude, "longitude: ", longitude)
			place := place2.Place{"", title, latitude, longitude, description}
			newPlace, err := place_io.CreatePlace(place)
			if err != nil {
				fmt.Println(err, " error when creating a new Place")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/place", 301)
				return
			}
			//Creation of pageFlow
			if src != "" && pageFlowTitle != "" {
				eventPageFlowObject := contribution.EventPageFlow{"", newPlace.Id, pageFlowTitle, src}
				_, err := event_io.CreateEventPageFlow(eventPageFlowObject)
				if err != nil {
					fmt.Println(err, " error when creating placePageFLow")
				}
			}
			placeTypeObject := place2.PlaceType{newPlace.Id, category}
			_, err = place_io.CreatePlaceType(placeTypeObject)
			if err != nil {
				fmt.Println(err, " error when creating placeType")
			}

			//History
			historiesObject := history2.Histories{"", misc.ConvertToByteArray(history)}
			histories, err := history_io.CreateHistorie(historiesObject)
			if err != nil {
				fmt.Println(err, " error creating a new History")
			} else {
				placeHistoryObject := place2.PlaceHistories{"", newPlace.Id, histories.Id}
				_, err = place_io.CreatePlaceHistpory(placeHistoryObject)
				if err != nil {
					fmt.Println(err, " error creating a new placeHistory")
				}
			}
			imageObject, err := misc.CreateImageHelper(content, description)
			if err != nil {
				fmt.Println(err, " error creating a new image")
			}
			placeImageObject := place2.PlaceImage{"", newPlace.Id, imageObject.Id, description}
			_, errr := place_io.CreatePlaceImage(placeImageObject)
			if errr != nil {
				fmt.Println(errr, " error creating placeImage")
			}

			//Here we are trying to make sure that newPlace.Id is not nil.
			if newPlace.Id != "" {
				fmt.Println("successful")
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Place : "+newPlace.Title)
				http.Redirect(w, r, "/admin_user/place", 301)
				return
			}
		}

		fmt.Println("One of the field is missing or newPlace.Id is nil")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/place", 301)
		return
	}
}

//Agreggeted Place Data
type PlaceData struct {
	Place   place2.Place
	Images  []image2.Images
	History history2.History
}

func getPlaceData(placeId string) PlaceData {
	var placeDate PlaceData
	var images []image2.Images
	var history history2.History
	place, err := place_io.ReadPlace(placeId)
	if err != nil {
		fmt.Println("error reading place data")
		return placeDate
	}
	placeImages, err := place_io.ReadPlaceImageAllOf(placeId)
	if err != nil {
		fmt.Println("error reading placeImages")
	} else {
		for _, placeImage := range placeImages {
			image, err := image_io.ReadImage(placeImage.ImageId)
			if err != nil {
				fmt.Println("error reading Images of the following placeId", placeImage.PlaceId)
			} else {
				images = append(images, image)
			}
		}
	}
	placeHistory, err := place_io.ReadPlaceHistporyOf(placeId)
	if err != nil {
		fmt.Println("error reading PlaceHistories")
	} else {
		history, err = history_io.ReadHistory(placeHistory.HistoryId)
		if err != nil {
			fmt.Println("error reading HistoryId of the following place", placeId)
		}
	}
	placeDate = PlaceData{place, images, history}
	return placeDate
}

//Todo The following method need to be double checked
func deletePlaceData(placeId string) (bool, string) {
	var stringToreturn string
	_, err := place_io.ReadPlace(placeId)
	if err != nil {
		fmt.Println("error reading place data")
		return false, "Error Place can not be find"
	} else {
		_, err := place_io.DeletePlace(placeId)
		if err != nil {
			fmt.Println("error reading place data")
			return false, "Error Place can not be Deleted"
		}
	}
	placeImages, err := place_io.ReadPlaceImageAllOf(placeId)
	if err != nil {
		fmt.Println("error reading placeImages")
	} else {
		for _, placeImage := range placeImages {
			_, err := image_io.DeleteImage(placeImage.ImageId)
			if err != nil {
				fmt.Println("error Deleting Image of the following placeId", placeImage.PlaceId)
			} else {
				_, err := place_io.DeletePlaceImage(placeImage.Id)
				if err != nil {
					fmt.Println("error Deleting PlaceImage of the following placeId", placeImage.PlaceId)
				}
			}
		}
	}
	placeHistory, err := place_io.ReadPlaceHistporyOf(placeId)
	if err != nil {
		fmt.Println("error reading PlaceHistories")
	} else {
		_, err := history_io.DeleteHistory(placeHistory.HistoryId)
		if err != nil {
			fmt.Println("error reading HistoryId of the following place", placeId)
		} else {
			_, err := place_io.DeletePlaceHistpory(placeHistory.Id)
			if err != nil {
				fmt.Println("error delete PlaceHistories of the following place", placeId)
			}
		}
	}
	return true, stringToreturn
}
