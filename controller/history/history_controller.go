package history

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	contribution2 "ostmfe/domain/contribution"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/io/comment_io"
	"ostmfe/io/contribution_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"time"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/single_history/{historyId}", SingleHistoryHandler(app))
	r.Post("/create-comment", CreateHistoryComment(app))
	r.Post("/contribution", CreateContributionComment(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func CreateContributionComment(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		var isExtension bool
		var fileTypeId string
		r.ParseForm()
		file, m, err := r.FormFile("file")
		historyId := r.PostFormValue("historyId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
			//Check the file extension
			isExtension, fileTypeId = misc.GetFileExtension(m)

			fmt.Println("result from assement: ", isExtension)

			if isExtension == false {
				fmt.Println("error creating contribution")
				fmt.Println("wrong file: ", m.Filename)
				http.Redirect(w, r, "/event/single/"+historyId, 301)
			}
		}
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		cellphone := r.PostFormValue("cellphone")

		if name != "" && email != "" && message != "" && historyId != "" {
			contributionObject := contribution2.Contribution{"", email, name, misc.FormatDateTime(time.Now()), cellphone, misc.ConvertToByteArray(message)}

			contribution, err := contribution_io.CreateContribution(contributionObject)
			if err != nil {
				fmt.Println("error creating a new contribution")
			} else {
				contributorEventObject := contribution2.ContributionEvent{"", contribution.Id, historyId, name}
				_, err := contribution_io.CreateContributionEvent(contributorEventObject)
				if err != nil {
					_, _ = contribution_io.DeleteContribution(contribution.Id)
					fmt.Println("error creating a new contribution")
				} else {
					contributionFileObject := contribution2.ContributionFile{"", contribution.Id, content, fileTypeId}
					_, err := contribution_io.CreateContributionFile(contributionFileObject)
					if err != nil {
						fmt.Println("error creating contributionFile")
					}
				}
			}
		}
		http.Redirect(w, r, "/event/single_history/"+historyId, 301)
	}
}

func CreateHistoryComment(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		historyId := r.PostFormValue("historyId")

		if historyId != "" && email != "" && message != "" {
			commentObject := comment.Comment{"", email, name, misc.FormatDateTime(time.Now()), misc.ConvertToByteArray(message), ""}
			newComment, err := comment_io.CreateComment(commentObject)
			if err != nil {
				fmt.Println("error creating comment")
			} else {
				_, err := comment_io.CreateCommentHistory(comment.CommentHistory{"", historyId, newComment.Id})
				if err != nil {
					fmt.Println("error creating comment")
				}
				http.Redirect(w, r, "/history/single_history/"+historyId, 301)
			}
		}
		http.Redirect(w, r, "/history/single_history/"+historyId, 301)
	}
}

func SingleHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		historyId := chi.URLParam(r, "historyId")
		if historyId == "" {
			http.Redirect(w, r, "/", 301)
		}
		commentNumber, err := comment_io.CountCommentHistory(historyId)
		if err != nil {
			fmt.Println("error reading CommentEvent")
		}

		type PageData struct {
			History       HistoryData
			CommentNumber int64
			Comments      []comment.CommentStack
		}
		data := PageData{getHistoryData(historyId), commentNumber, getHistoryComments(historyId)}
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

	historyDataObject := HistoryData{history2.History{History.Id, History.Title, History.Description, misc.FormatDateMonth(History.Date)}, profile, images, histories}

	return historyDataObject
}

//History Comments
func getHistoryComments(historyId string) []comment.CommentStack {
	var parentCommentStack []comment.CommentStack
	var subCommentStack []comment.CommentHelper

	for _, commentObject := range getComments(historyId) {
		myComment, err := comment_io.ReadComment(commentObject.Id)
		if err != nil {
			fmt.Println("error reading Comment")
		}
		if myComment.ParentCommentId != "" && myComment.Comment != nil {
			subCommentStack = getSubComment(commentObject.Id)
		}
		parentCommentStack = append(parentCommentStack, comment.CommentStack{commentObject, subCommentStack})
	}

	fmt.Println("parentStack ", parentCommentStack)

	return parentCommentStack
}

//This method returns a list of either parent or subcomment
func getComments(historyId string) []comment.CommentHelper {
	var myCommentObject []comment.CommentHelper
	eventComments, err := comment_io.ReadAllByHistoryId(historyId)
	if err != nil {
		fmt.Println("error reading event Comment")
		return myCommentObject
	}
	for _, eventComment := range eventComments {
		myComment, err := comment_io.ReadComment(eventComment.CommentId)
		if err != nil {
			fmt.Println("error reading Comment")
		} else {
			commentHelper := comment.CommentHelper{myComment.Id, myComment.Email, myComment.Name, misc.FormatDateMonth(myComment.Date), misc.ConvertingToString(myComment.Comment), myComment.ParentCommentId}
			myCommentObject = append(myCommentObject, commentHelper)
		}
	}
	return myCommentObject
}

func getSubComment(parentComment string) []comment.CommentHelper {
	var myComments []comment.CommentHelper
	subComments, err := comment_io.ReadCommentWithParentId(parentComment)
	if err != nil {
		return myComments
	}
	for _, eventComment := range subComments {
		if eventComment.ParentCommentId == parentComment && eventComment.Comment != nil {
			commentHelper := comment.CommentHelper{eventComment.Id, eventComment.Email, eventComment.Name, misc.FormatDateMonth(eventComment.Date), misc.ConvertingToString(eventComment.Comment), eventComment.ParentCommentId}
			myComments = append(myComments, commentHelper)
		}
	}
	return myComments
}
