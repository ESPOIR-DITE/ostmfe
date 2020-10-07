package event

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	museum "ostmfe/domain"
	"ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	"ostmfe/domain/project"
	io2 "ostmfe/io"
	"ostmfe/io/event_io"
	"ostmfe/io/project_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/single/{eventId}", EventHanler(app))
	r.Get("/ofayear/{yearId}", EventOfAYearHanler(app))
	return r
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
		type PageData struct {
			EventData EventData
			Place     place2.Place
			Peoples   []people.People
			GroupData []GroupData
			Project   project.Project
		}
		data := PageData{eventdata, GetEnventPlaceData(eventId), GetEventPeopleData(eventId), GetGroupsData(eventId), getEventProject(eventId)}
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
