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
	"ostmfe/io/event_io"
	"ostmfe/io/image_io"
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

type EventData struct {
	Event        event2.Event
	ProfileImage image3.Images
	Images       []image3.Images
	//Location string
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects := misc.GetProjectContentsHomes()
		var images []image3.Images
		var profileImage image3.Images
		var eventDataList []EventData

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
					fmt.Println(" profileImage.Id: ", profileImage.Id)
					eventData := EventData{event, profileImage, images}
					eventDataList = append(eventDataList, eventData)
					eventData = EventData{}
				}
				fmt.Println("This error may occur if there is no events created error:  profileImage is empty")

			}

		}
		type PageData struct {
			Projects      []misc.ProjectContentsHome
			EventDataList []EventData
		}
		date := PageData{projects, eventDataList}
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
