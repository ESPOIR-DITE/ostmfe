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
	"ostmfe/domain/pages"
	"ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	"ostmfe/domain/project"
	io2 "ostmfe/io"
	"ostmfe/io/comment_io"
	"ostmfe/io/contribution_io"
	"ostmfe/io/event_io"
	"ostmfe/io/pageData_io"
	"ostmfe/io/pages/client"
	"ostmfe/io/project_io"
	"path/filepath"
	"strings"
	"time"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	//r.Get("/", homeHanler(app))
	r.Get("/", homeHandler(app))
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
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
			//Check the file extension
			isExtension, fileTypeId = getFileExtension(m)

			fmt.Println("result from assement: ", isExtension)

			if isExtension == false {
				fmt.Println("error creating contribution")
				fmt.Println("wrong file: ", m.Filename)
				http.Redirect(w, r, "/event/single/"+eventId, 301)
			}
		}
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		fileId := r.PostFormValue("fileId")
		cellphone := r.PostFormValue("cellphone")

		if name != "" && email != "" && message != "" && eventId != "" && fileId != "" {
			contributionObject := contribution2.Contribution{"", email, name, time.Now(), cellphone, misc.ConvertToByteArray(message)}

			contribution, err := contribution_io.CreateContribution(contributionObject)
			if err != nil {
				fmt.Println("error creating a new contribution")
			} else {
				contributorEventObject := contribution2.ContributionEvent{"", contribution.Id, eventId, name}
				_, err := contribution_io.CreateContributionEvent(contributorEventObject)
				if err != nil {
					_, _ = contribution_io.DeleteContribution(contribution.Id)
					fmt.Println("error creating a new contribution")
				} else {
					contributionFileObject := contribution2.ContributionFile{"", contribution.Id, content, fileId, fileTypeId}
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
			fmt.Println("extension: " + extension + " file extension: " + contributionFileType.FileType)
			//t := strings.Trim(extension, ".")
			t := strings.Replace(extension, ".", "", -1)
			fmt.Println("extension2: " + t + " file extension: " + contributionFileType.FileType)
			if t == contributionFileType.FileType {
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
			commentObject := comment.Comment{"", email, name, misc.FormatDateTime(time.Now()), misc.ConvertToByteArray(message), "", false}
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
		var bannerImage string
		banner, err := misc.GetBanner("Event-Page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = banner.Id
		}

		events := misc.GetSimpleEventDataOfYear(yearId)
		if len(events) == 0 {
			http.Redirect(w, r, "/event", 301)
			return
		}
		type PageData struct {
			ProjectBanner string
			Events        []misc.SimpleEventData
			Years         []YearData
			Year          museum.Years
		}
		data := PageData{bannerImage, events, getYearDate(), year}
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
		if eventId == "" {
			http.Redirect(w, r, "/", 301)
		}
		var bannerImage string
		banner, err := misc.GetBanner("Event-Page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = banner.Id
		}
		eventdata := GetEventData(eventId)

		eventNumber, err := comment_io.CountCommentEvent(eventId)
		if err != nil {
			fmt.Println("error reading counting CommentEvent")
		}
		pageFlow, err := event_io.ReadAllEventPageFlowByEventId(eventId)
		if err != nil {
			fmt.Println("error reading event Page flow")
		}

		type PageData struct {
			ProjectBanner string
			EventData     EventData
			Place         place2.Place
			Peoples       []people.People
			GroupData     []GroupData
			Project       project.Project
			CommentNumber int64
			Comments      []comment.CommentStack
			GalleryImages []misc.EventGalleryImages
			PageFlow      []contribution2.EventPageFlow
		}
		data := PageData{bannerImage, eventdata,
			GetEnventPlaceData(eventId),
			GetEventPeopleData(eventId),
			GetGroupsData(eventId),
			getEventProject(eventId),
			eventNumber,
			comment2.GetEventComments(eventId),
			misc.GetEventGallery(eventId),
			pageFlow,
		}
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

		var bannerImage string
		pageBanner, err := pageData_io.ReadPageBannerWIthPageName("event-page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = misc.GetBannerImage(pageBanner.BannerId)
		}

		type PageData struct {
			Events      []misc.SimpleEventData
			Years       []YearData
			Projects    []project.Project
			EventBanner string
		}
		data := PageData{events, getYearDate(), projects, bannerImage}
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
func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var bannerImage string
		banner, err := misc.GetBanner("Event-Page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = banner.Id
		}
		pageData, err := client.EVentClientPage()
		if err != nil {
			// todo redirect to the home page!
		}
		type PageData struct {
			ProjectBanner string
			PageData      pages.EventPageData
			Projects      []project.Project
		}
		data := PageData{bannerImage, pageData, pageData.Project}
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
