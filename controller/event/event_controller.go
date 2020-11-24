package event

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"ostmfe/config"
	comment2 "ostmfe/controller/comment"
	"ostmfe/controller/misc"
	museum "ostmfe/domain"
	"ostmfe/domain/comment"
	contribution2 "ostmfe/domain/contribution"
	"ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	"ostmfe/domain/project"
	io2 "ostmfe/io"
	"ostmfe/io/comment_io"
	"ostmfe/io/contribution_io"
	"ostmfe/io/event_io"
	"ostmfe/io/project_io"
	"path/filepath"
	"time"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/single/{eventId}", EventHanler(app))
	r.Get("/ofayear/{yearId}", EventOfAYearHanler(app))
	r.Post("/create", CreateComment(app))
	r.Post("/contribution", CreateContributionComment(app))
	return r
}

func CreateContributionComment(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var content []byte
		var isExtension bool
		var fileTypeId string
		r.ParseForm()
		file, m, err := r.FormFile("file")
		eventId := r.PostFormValue("eventId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
			//Check the file extension
			isExtension, fileTypeId = getFileExtension(m)

			if !isExtension {
				fmt.Println("error creating contribution")
				fmt.Println("wrong file: ", m.Filename)
				http.Redirect(w, r, "/event/single/"+eventId, 301)
			}
		}
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		cellphone := r.PostFormValue("cellphone")

		if name != "" && email != "" && message != "" && eventId != "" {
			contributionObject := contribution2.Contribution{"", email, name, misc.FormatDateTime(time.Now()), cellphone, misc.ConvertToByteArray(message)}

			contribution, err := contribution_io.CreateContribution(contributionObject)
			if err != nil {
				fmt.Println("error creating a new contribution")
			} else {
				contributorEventObject := contribution2.ContributionEvent{"", contribution.Id, eventId, name}
				_, err := contribution_io.CreateContributionEvent(contributorEventObject)
				if err != nil {
					contribution_io.DeleteContribution(contribution.Id)
					fmt.Println("error creating a new contribution")
				} else {
					contributionFileObject := contribution2.ContributionFile{"", content, fileTypeId}
					_, err := contribution_io.CreateContributionFile(contributionFileObject)
					if err != nil {
						fmt.Println("error creating contributionFile")
					}
				}
			}
		}
		http.Redirect(w, r, "/event/single/"+eventId, 301)
	}
}

func getFileExtension(fileData *multipart.FileHeader) (bool, string) {
	var extension = filepath.Ext(fileData.Filename)
	contributionFileTypes, err := contribution_io.ReadContributionFileTypes()
	if err != nil {
		fmt.Println("error reading contributionFileType")
		return true, ""
	} else {
		for _, contributionFileType := range contributionFileTypes {
			if extension == contributionFileType.FileType {
				return true, contributionFileType.Id
			}
		}
	}
	return false, ""
}
func CreateComment(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		eventId := r.PostFormValue("eventId")

		if name != "" && email != "" && message != "" {
			commentObject := comment.Comment{"", email, name, misc.FormatDateTime(time.Now()), misc.ConvertToByteArray(message), ""}
			newComment, err := comment_io.CreateComment(commentObject)
			if err != nil {
				fmt.Println("error creating comment")
			} else {
				_, err := comment_io.CreateCommentEvent(comment.CommentEvent{"", eventId, newComment.Id})
				if err != nil {
					fmt.Println("error creating comment")
				}
				http.Redirect(w, r, "/event/single/"+eventId, 301)
			}
		}
		http.Redirect(w, r, "/event/single/"+eventId, 301)
	}
}

func EventOfAYearHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		yearId := chi.URLParam(r, "yearId")
		year, err := io2.ReadYear(yearId)
		if err != nil && year.Id == "" {
			fmt.Println(err, " error reading year")
			http.Redirect(w, r, "/event", 301)
			return
		}
		_, errs := event_io.ReadEventYearsWithYearId(yearId)
		if errs != nil {
			fmt.Println(errs, " error reading event years")
			http.Redirect(w, r, "/event", 301)
			return
		}

		events := misc.GetSimpleEventDataOfYear(yearId)
		if len(events) == 0 {
			http.Redirect(w, r, "/event", 301)
			return
		}
		type PageData struct {
			Events []misc.SimpleEventData
			Years  []YearData
			Year   museum.Years
		}
		data := PageData{events, getYearDate(), year}
		files := []string{
			app.Path + "event/events_year.html",
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

func EventHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventId := chi.URLParam(r, "eventId")
		eventdata := GetEventData(eventId)
		eventNumber, err := comment_io.CountCommentEvent(eventId)
		if err != nil {
			fmt.Println("error reading COmmentEvent")
		}

		type PageData struct {
			EventData     EventData
			Place         place2.Place
			Peoples       []people.People
			GroupData     []GroupData
			Project       project.Project
			CommentNumber int64
			Comments      []comment.CommentStack
		}
		data := PageData{eventdata, GetEnventPlaceData(eventId), GetEventPeopleData(eventId), GetGroupsData(eventId), getEventProject(eventId), eventNumber, comment2.GetEventComments(eventId)}
		files := []string{
			app.Path + "event/event_single.html",
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
		events := misc.GetSimpleEventData(5)
		projects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading projects")
		}
		type PageData struct {
			Events   []misc.SimpleEventData
			Years    []YearData
			Projects []project.Project
		}
		data := PageData{events, getYearDate(), projects}
		files := []string{
			app.Path + "event/events.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
			app.Path + "base_templates/projects_template.html",
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
			fmt.Println("Year with number of event: ", yearResult, " number: ", amount)
			if err != nil {
				fmt.Println(err, "error reading year with yearId.")
			} else {
				year = append(year, YearData{yearResult, amount})
			}
		}
	}
	return year
}

func getEventProject(eventId string) project.Project {
	var projectObject project.Project

	eventProject, err := event_io.ReadEventProjectWithEventId(eventId)
	if err != nil {
		fmt.Println(err, " could not find read the eventProject")
		return projectObject
	}
	project, err := project_io.ReadProject(eventProject.ProjectId)
	if err != nil {
		fmt.Println(err, " could not find read the Project")
		return projectObject
	}
	return project
}
