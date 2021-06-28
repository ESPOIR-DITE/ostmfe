package histories

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/domain/contribution"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	"ostmfe/io/comment_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
)

func HistoryHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", HistoryHandler(app))
	r.Get("/edit/{historyId}", EditHistoryHandler(app))
	r.Post("/create", CreateHistpory(app))
	r.Get("/create_step2/{historyId}", CreateImageHelper(app))

	r.Post("/create_image", CreateHistoryImageHandler(app))
	r.Post("/create_histories", CreateHistoriesImageHandler(app))
	r.Post("/update_pictures", UpdateHistoryImageHandler(app))
	r.Post("/update_details", UpdateHistoryDetailsHandler(app))
	r.Post("/update_histories", UpdateHistoryHistoriessHandler(app))
	r.Get("/delete_image/{imageId}/{historyId}", DeleteHistoryImage(app))
	r.Get("/delete_history/{historyId}", DeleteHistoryHandler(app))

	//gallery
	r.Post("/create-gallery", CreateEventHistoryHandler(app))
	r.Post("/create-page-flow", CreatePageFlowHandler(app))
	r.Get("/delete-pageFlow/{historyPageFlowId}/{historyId}", DeletePageFlowHandler(app))
	r.Get("/delete-gallery/{pictureId}/{historyId}/{historyGalleryId}", DeleteGalleryHandler(app))
	r.Get("/activate_comment/{commentId}/{historyId}", ActivateCommentHandler(app))
	return r
}

func DeletePageFlowHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		historyPageFlowId := chi.URLParam(r, "historyPageFlowId")
		historyId := chi.URLParam(r, "historyId")

		_, err := history_io.DeleteHistoryPageFLow(historyPageFlowId)
		if err != nil {
			fmt.Println("error deleting History Page FLow")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: Project Gallery. ")
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return
	}
}

func CreatePageFlowHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		historyId := r.PostFormValue("historyId")
		pageFlowTitle := r.PostFormValue("pageFlowTitle")
		scr := r.PostFormValue("scr")

		if scr != "" && historyId != "" && pageFlowTitle != "" {
			_, err := history_io.CreateHistoryPageFLow(history2.HistoryPageFlow{"", pageFlowTitle, historyId, scr})
			if err != nil {
				fmt.Println(err, " error creating page flow!")
			}
		} else {
			app.ErrorLog.Print("Error creating HistoryPageFlow")
		}
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return
	}
}

func ActivateCommentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		commentId := chi.URLParam(r, "commentId")
		historyId := chi.URLParam(r, "historyId")
		result := misc.ActivateComment(commentId)
		fmt.Print("Activation Result: ", result)
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return
	}
}

func DeleteGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		pictureId := chi.URLParam(r, "pictureId")
		historyId := chi.URLParam(r, "historyId")
		historyGalleryId := chi.URLParam(r, "historyGalleryId")

		//Deleting project
		gallery, err := image_io.DeleteGalery(pictureId)
		if err != nil {
			fmt.Println("error deleting gallery")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		} else {
			_, err := history_io.DeleteHistoryGallery(historyGalleryId)
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
				http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
				return
			}
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: people Gallery. ")
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return
	}
}

func CreateEventHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		historyId := r.PostFormValue("historyId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if historyId != "" && description != "" {
			galery := image.Gallery{"", content, description}
			galleryObject, err := image_io.CreateGalery(galery)
			if err != nil {
				fmt.Println(err, " error creating gallery")
			} else {
				historyGalery := history2.HistoryGalery{"", historyId, galleryObject.Id}
				_, err := history_io.CreateHistoryGallery(historyGalery)
				if err != nil {
					fmt.Println(err, " error creating GroupGallery")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return
	}
}

func DeleteHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		historyId := chi.URLParam(r, "historyId")

		//Check if the history tobe updated exists
		history, err := history_io.ReadHistory(historyId)
		if err != nil {
			fmt.Println(err, " error reading history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history", 301)
			return
		} else {
			_, err := history_io.DeleteHistory(historyId)
			if err != nil {
				fmt.Println(err, " error reading history")
			}
		}
		//checking and deleting  HistoryImage and image
		historyImage, err := history_io.ReadHistoryImageWithHistoryId(historyId)
		if err != nil {
			fmt.Println(err, " error reading history Image, this history may not have an image")
		} else {
			_, err := history_io.DeleteHistoryImage(historyImage.Id)
			if err != nil {
				fmt.Println(err, " Could not delete History Image")
			} else {
				_, err := image_io.DeleteImage(historyImage.ImageId)
				if err != nil {
					fmt.Println(err, " could not delete image")
				}
			}
		}

		//checking and deleting  HistoryHistories and histories
		histories, err := history_io.ReadHistoryHistoriesWithHistoryId(historyId)
		if err != nil {
			fmt.Println(err, " error reading history Histories, this history may not have an Histories")
		} else {
			_, err := history_io.DeleteHistoryHistories(histories.Id)
			if err != nil {
				fmt.Println(err, " Could not delete History histories")
			} else {
				_, err := history_io.DeleteHistorie(histories.HistoriesId)
				if err != nil {
					fmt.Println(err, " Could not delete histories")
				}
			}
		}

		fmt.Println(err, " Delete successful")
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully Updated Details for the following History  : "+history.Title)
		http.Redirect(w, r, "/admin_user/history", 301)
		return
	}
}

func UpdateHistoryHistoriessHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		historyId := r.PostFormValue("historyId")
		myArea := r.PostFormValue("myArea")
		historiesId := r.PostFormValue("historiesId")

		//Check if the history tobe updated exists
		history, err := history_io.ReadHistory(historyId)
		if err != nil {
			fmt.Println(err, " error reading history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history", 301)
			return
		}
		if myArea != "" {
			historyies := history2.Histories{historiesId, misc.ConvertToByteArray(myArea)}
			_, err := history_io.UpdateHistorie(historyies)
			if err != nil {
				fmt.Println(err, " error create histories")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/history", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully Updated Details for the following History  : "+history.Title)
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}
		fmt.Println(" error updating History Details. One Field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return

	}
}

func CreateHistoriesImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		historyId := r.PostFormValue("historyId")
		myArea := r.PostFormValue("myArea")

		//Check if the history tobe updated exists
		history, err := history_io.ReadHistory(historyId)
		if err != nil {
			fmt.Println(err, " error reading history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history", 301)
			return
		}
		if myArea != "" {
			historyies := history2.Histories{"", misc.ConvertToByteArray(myArea)}
			newHistoryies, err := history_io.CreateHistorie(historyies)
			if err != nil {
				fmt.Println(err, " error create histories")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/history", 301)
				return
			}
			historyHistoriesObejct := history2.HistoryHistories{"", historyId, newHistoryies.Id}
			_, errx := history_io.CreateHistoryHistory(historyHistoriesObejct)
			if errx != nil {
				fmt.Println(err, " error create history-histories")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/history", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully Updated Details for the following History  : "+history.Title)
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}
		fmt.Println(" error updating History Details. One Field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return

	}
}

func UpdateHistoryDetailsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		historyId := r.PostFormValue("historyId")
		historyTitle := r.PostFormValue("historyTitle")
		historyDate := r.PostFormValue("historyDate")
		description := r.PostFormValue("description")

		//Check if the history tobe updated exists
		history, err := history_io.ReadHistory(historyId)
		if err != nil {
			fmt.Println(err, " error reading history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history", 301)
			return
		}
		if historyTitle != "" && historyDate != "" && description != "" {
			historyObject := history2.History{historyId, historyTitle, description, historyDate}

			_, err := history_io.UpdateHistory(historyObject)
			if err != nil {
				fmt.Println(err, " could not update History")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully Updated Details for the following History  : "+history.Title)
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}
		fmt.Println(" error updating History Details. One Field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return
	}

}

func DeleteHistoryImage(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		historyId := chi.URLParam(r, "historyId")
		imageId := chi.URLParam(r, "imageId")

		history, err := history_io.ReadHistory(historyId)
		if err != nil {
			fmt.Println(err, " error reading history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}
		//Checking History Image
		historyImage, err := history_io.ReadHistoryImageWithHistoryId(historyId)
		if err != nil {
			fmt.Println(err, " error reading history Image")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}
		//Checking the image
		_, err = image_io.ReadImage(imageId)
		if err != nil {
			fmt.Println(err, " error reading Image")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}

		//If all passes than.
		_, errs := history_io.DeleteHistoryImage(historyImage.Id)
		if errs != nil {
			fmt.Println(err, " error deleting history Image")
		} else {
			_, err := image_io.DeleteImage(historyImage.ImageId)
			if err != nil {
				fmt.Println(err, " error deleting Image")
			}
		}
		fmt.Println(" deleting Successful")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully Deleted image for the following History  : "+history.Title)
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return

	}
}

func UpdateHistoryImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		file, _, err := r.FormFile("file")
		historyId := r.PostFormValue("historyId")
		historyImageId := r.PostFormValue("historyImageId")
		imageId := r.PostFormValue("imageId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		//Checking the History
		history, err := history_io.ReadHistory(historyId)
		if err != nil {
			fmt.Println(err, " could not read History")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history", 301)
			return
		}
		// Reading History Image
		historyImageObejct, err := history_io.ReadHistoryImage(historyImageId)
		if err != nil {
			fmt.Println(err, " could not read HistoryImage")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}

		filesArray := []io.Reader{file}
		filesByteArray := misc.CheckFiles(filesArray)
		historyImage := history2.HistoryImage{historyImageId, imageId, historyId, historyImageObejct.Description}

		helper := history2.HistoryImageHelper{historyImage, filesByteArray}
		_, errr := history_io.UpdateHistoryImage(helper)
		if errr != nil {
			fmt.Println(errr, " error updating HistoryImage this error may occur")
			//if app.Session.GetString(r.Context(), "user-create-error") != "" {
			//	app.Session.Remove(r.Context(), "user-create-error")
			//}
			//app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			//http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			//return
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully Updated image for the following History  : "+history.Title)
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return
	}
}

func CreateHistoryImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		file, _, err := r.FormFile("file")
		historyId := r.PostFormValue("historyId")
		imageType := r.PostFormValue("imageType")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		history, err := history_io.ReadHistory(historyId)
		if err != nil {
			fmt.Println(err, " could not read History")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}

		filesArray := []io.Reader{file}
		filesByteArray := misc.CheckFiles(filesArray)
		historyImage := history2.HistoryImage{"", "", historyId, imageType}

		helper := history2.HistoryImageHelper{historyImage, filesByteArray}
		_, errr := history_io.CreateHistoryImage(helper)
		if errr != nil {
			fmt.Println(errr, " error creating HistoryImage")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully created image for the following History  : "+history.Title)
		http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
		return
	}
}

func CreateHistpory(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		title := r.PostFormValue("title")
		description := r.PostFormValue("description")
		date := r.PostFormValue("date")
		mytextarea := r.PostFormValue("mytextarea")
		fmt.Println("Title: ", title,
			"Date: ", date,
			"description: ", description,
			"mytextArea: ", mytextarea)

		if title != "" && mytextarea != "" {
			history := history2.History{"", title, description, date}
			createdHistory, err := history_io.CreateHistory(history)
			if err != nil {
				fmt.Println("error creating history: ", err)
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/history", 301)
				return
			}
			historyies := history2.Histories{"", misc.ConvertToByteArray(mytextarea)}
			newHistoryies, err := history_io.CreateHistorie(historyies)
			if err != nil {
				fmt.Println(err, " error create histories")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/history", 301)
				return
			}
			historyHistoriesObejct := history2.HistoryHistories{"", createdHistory.Id, newHistoryies.Id}
			_, errx := history_io.CreateHistoryHistory(historyHistoriesObejct)
			if errx != nil {
				fmt.Println(err, " error create history-histories")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/history", 301)
				return
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully created an new HistoryId : "+createdHistory.Title)
			http.Redirect(w, r, "/admin_user/history/create_step2/"+createdHistory.Id, 301)
			return
		}
		fmt.Println("fail to create, one field is empty")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/history", 301)
		return
	}
}
func EditHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		historyId := chi.URLParam(r, "historyId")

		pageFlow, err := history_io.ReadHistoryPageFLowsWithHistoryId(historyId)

		commentNumber, pendingcomments, activeComments := historyCommentCalculation(historyId)
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PageData struct {
			HistoryData     HistorySimpleData
			SidebarData     misc.SidebarData
			Comments        []comment.CommentHelper2
			Gallery         []misc.HistoryGalleryImages
			CommentNumber   int64
			PendingComments int64
			ActiveComments  int64
			Contritbution   []contribution.Contribution
			PageFlows       []history2.HistoryPageFlow
			AdminName       string
			AdminImage      string
		}
		data := PageData{GetHistorySimpleData(historyId),
			misc.GetSideBarData("history", ""),
			misc.GetHistoryComments(historyId), misc.GetHistoryGallery(historyId),
			commentNumber,
			pendingcomments,
			activeComments,
			GetContribution(historyId),
			pageFlow, adminName, adminImage}
		files := []string{
			app.Path + "admin/history/edit_history.html",
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

func NewHistoryHandler(app *config.Env) http.HandlerFunc {
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
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			AdminName     string
			AdminImage    string
		}
		data := PagePage{backend_error, unknown_error, adminName, adminImage}
		files := []string{
			app.Path + "admin/history/new_history.html",
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

func HistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		var unknown_error string
		var backend_error string
		var historieList []history2.History
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			unknown_error = app.Session.GetString(r.Context(), "creation-unknown-error")
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			backend_error = app.Session.GetString(r.Context(), "user-create-error")
			app.Session.Remove(r.Context(), "user-create-error")
		}
		histories, err := history_io.ReadHistorys()
		if err != nil {
			fmt.Println(err, "Error reading Histories")
		}
		for _, history := range histories {
			historyObejct := history2.History{history.Id, history.Title, history.Description, misc.FormatDateMonth(history.Date)}
			historieList = append(historieList, historyObejct)
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			Histories     []history2.History
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}

		data := PagePage{backend_error, unknown_error, historieList,
			misc.GetSideBarData("history", ""), adminName, adminImage}
		files := []string{
			app.Path + "admin/history/history.html",
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

func CreateImageHelper(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		historyId := chi.URLParam(r, "historyId")

		history, err := history_io.ReadHistory(historyId)
		if err != nil {
			fmt.Println(err, " error reading history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/history/edit/"+historyId, 301)
			return
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PageData struct {
			History    history2.History
			AdminName  string
			AdminImage string
		}

		data := PageData{history, adminName, adminImage}
		files := []string{
			app.Path + "admin/history/image_history.html",
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

//With historyId, you get the commentNumber, pending, active.
func historyCommentCalculation(historyId string) (commentNumber int64, pending int64, active int64) {
	var commentNumbers int64 = 0
	var pendings int64 = 0
	var actives int64 = 0
	historyComments, err := comment_io.ReadAllCommentHistory(historyId)
	if err != nil {
		fmt.Println(err, " error reading People comment")
		return commentNumbers, pendings, actives
	} else {
		for _, historyComment := range historyComments {
			comments, err := comment_io.ReadComment(historyComment.CommentId)
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
