package history

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/single_history/{historyId}", SingleHistoryHandler(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func SingleHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		historyId := chi.URLParam(r, "historyId")

		type PageData struct {
			History HistoryData
		}
		data := PageData{getHistoryData(historyId)}
		files := []string{
			app.Path + "history/history_single.html",
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
		history, err := history_io.ReadHistorys()
		if err != nil {
			fmt.Println(err, " error reading histories")
		}

		type PageData struct {
			History []history2.History
		}
		data := PageData{history}
		files := []string{
			app.Path + "history/history.html",
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

type HistoryData struct {
	History   history2.History
	Profile   image3.Images
	Images    []image3.Images
	Histories history2.HistoriesHelper
}

func getHistoryData(historyId string) HistoryData {
	var historyData HistoryData
	var profile image3.Images
	var images []image3.Images
	var histories history2.HistoriesHelper
	//Check if the history exist
	History, err := history_io.ReadHistory(historyId)
	if err != nil {
		fmt.Println(err, " error reading history")
		return historyData
	}

	//Images
	historyImages, err := history_io.ReadHistoryImagesWithHistoryId(historyId)
	if err != nil {
		fmt.Println(err, " error reading history Images")
	} else {
		for _, historyImage := range historyImages {
			if historyImage.Description == "1" || historyImage.Description == "profile" {
				profile, err = image_io.ReadImage(historyImage.ImageId)
				if err != nil {
					fmt.Println(err, " error reading profile Images")
				}
			}
			image, err := image_io.ReadImage(historyImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading profile Images")
			}
			images = append(images, image)
		}
	}

	//History
	historyHistorie, err := history_io.ReadHistoryHistoriesWithHistoryId(historyId)
	if err != nil {
		fmt.Println(err, " error reading HistoryHistory")
	} else {
		history, err := history_io.ReadHistorie(historyHistorie.HistoriesId)
		if err != nil {
			fmt.Println(err, " error reading Historie")
		}
		histories = history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}
	}

	historyDataObject := HistoryData{History, profile, images, histories}

	return historyDataObject
}
