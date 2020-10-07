package controller

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/about_us"
	"ostmfe/controller/admin"
	"ostmfe/controller/admin/administration"
	"ostmfe/controller/collection"
	"ostmfe/controller/event"
	"ostmfe/controller/history"
	"ostmfe/controller/home"
	"ostmfe/controller/misc"
	"ostmfe/controller/people"
	"ostmfe/controller/place"
	"ostmfe/controller/project"
	"ostmfe/controller/user"
	"ostmfe/controller/visit"
	event2 "ostmfe/domain/event"
	image3 "ostmfe/domain/image"
	project2 "ostmfe/domain/project"
	"ostmfe/io/event_io"
	"ostmfe/io/image_io"
	"ostmfe/io/pageData_io"
	"ostmfe/io/project_io"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", homeHanler(env))
	mux.Mount("/home", home.Home(env))
	mux.Mount("/visit", visit.Home(env))
	mux.Mount("/history", history.Home(env))
	mux.Mount("/collection", collection.Home(env))
	mux.Mount("/place", place.Home(env))
	mux.Mount("/people", people.Home(env))
	mux.Mount("/admin_user", admin.Home(env))
	mux.Mount("/administration", administration.AdministrationController(env))
	mux.Mount("/user", user.Home(env))
	mux.Mount("/event", event.Home(env))
	mux.Mount("/about_us", about_us.Home(env))
	mux.Mount("/project", project.Home(env))

	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
}

type SimpleEventData struct {
	Event        event2.Event
	ProfileImage image3.Images
	Images       []image3.Images
	//Location string
}
type PageData struct {
	Projects      []misc.ProjectContentsHome
	EventDataList []misc.SimpleEventData
	//EventDataListLeft   []misc.SimpleEventData
	AllProjects     []project2.Project
	PagePageSection HomePageData
	CheckOdds       func(index int) bool
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects := misc.GetProjectContentsHomes()

		//var eventDataListLeft []EventData
		//var eventDataListRight []EventData

		allProjects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading all the project")
		}
		eventdataLeft := misc.GetSimpleEventData(6)

		date := PageData{projects,
			eventdataLeft,
			allProjects,
			GetPageData("HomePage"),
			func(index int) bool {
				if index%2 == 0 {
					return true
				} else {
					return false
				}
			},
		}

		files := []string{
			app.Path + "index.html",
			app.Path + "base_templates/navigator.html",
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

func CheckEventAndOdd(index int) bool {
	if index%2 == 0 {
		return true
	} else {
		return false
	}
}

type HomePageData struct {
	Notification    string
	ProjectIntro    string
	EventIntro      string
	ExhibitionIntro string
}

func GetPageData(pageName string) HomePageData {
	var notification string
	var projectintro string
	var eventintro string
	var exibitionintro string

	page, err := pageData_io.ReadPageDataWIthName(pageName)
	if err != nil {
		fmt.Println(err, " error reading page")
	} else {
		pageDateSectionObject, err := pageData_io.ReadPageSectionAllOf(page.Id)
		if err != nil {
			fmt.Println(err, " error reading page")
		}
		for _, pageDateSection := range pageDateSectionObject {
			pageSection, err := pageData_io.ReadSection(pageDateSection.SectionId)
			if err != nil {
				fmt.Println(err, " error reading page")
			} else {
				if pageSection.SectionName == "Notification" {
					fmt.Println(" Notification", pageSection)
					notification = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "ProjectIntro" {
					fmt.Println(" ProjectIntro", pageSection)
					projectintro = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "EventIntro" {
					fmt.Println(" EventIntro", pageSection)
					eventintro = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "ExhibitionIntro" {
					fmt.Println(" exhibition", pageSection)
					exibitionintro = misc.ConvertingToString(pageDateSection.Content)
				}
			}
		}
	}
	return HomePageData{notification, projectintro, eventintro, exibitionintro}
}

func GetEvents() []SimpleEventData {
	var images []image3.Images
	var profileImage image3.Images
	var eventDataList []SimpleEventData
	//Here we are reading all the events
	events, err := event_io.ReadEvents()
	if err != nil {
		fmt.Println(err, " error reading events")
	} else {
		for _, event := range events {
			eventImages, err := event_io.ReadEventImgOf(event.Id)
			if err != nil {
				fmt.Println(err, " error reading events Images")
			} else {
				fmt.Println(" Looping eventImages")
				for _, eventImage := range eventImages {

					fmt.Println(" eventImage.Description: ", eventImage.Description)
					if eventImage.Description == "1" || eventImage.Description == "profile" {
						fmt.Println(" We have a profile Image")
						profileImage, err = image_io.ReadImage(eventImage.ImageId)
						if err != nil {
							fmt.Println(err, " error reading profile event image")
						}
					}
					fmt.Println(" eventImage.ImageId: ", eventImage.ImageId)
					image, err := image_io.ReadImage(eventImage.ImageId)
					if err != nil {
						fmt.Println(err, " error reading image")
					}
					images = append(images, image)
				}
				//eventLocation,err:= ReadEvent
			}
			//we need to make sure that profileImage is not empty
			if profileImage.Id != "" {
				//fmt.Println(" profileImage.Id: ", profileImage.Id)
				eventData := SimpleEventData{event, profileImage, images}
				eventDataList = append(eventDataList, eventData)
				eventData = SimpleEventData{}

				//adding data to the correct list
				//if CheckEventAndOdd(index)
			}
			fmt.Println("This error may occur if there is no events created error:  profileImage is empty")

		}

	}
	return eventDataList
}
