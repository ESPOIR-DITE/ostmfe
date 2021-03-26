package slider

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
	"ostmfe/domain/slider"
	"ostmfe/io/slider_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/delete/{sliderId}", DeleteHanler(app))
	r.Post("/create", CreateSlider(app))
	return r
}

func DeleteHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !adminHelper.CheckAdminInSession(app, r) {
			fmt.Println("Admin session isseu")
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		sliderId := chi.URLParam(r, "sliderId")
		_, err := slider_io.ReadSlider(sliderId)
		if err != nil {
			fmt.Println(err, "could not create slider")
			http.Redirect(w, r, "/admin_user/slider", 301)
			return
		} else {
			_, err := slider_io.DeleteSlider(sliderId)
			if err != nil {
				fmt.Println(err, "could not delete slider")
				http.Redirect(w, r, "/admin_user/slider", 301)
				return
			}
		}
		fmt.Println(err, "Successfully deleted slider")
		http.Redirect(w, r, "/admin_user/slider", 301)
		return
	}
}

func CreateSlider(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		r.ParseForm()

		var content []byte
		file, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should not happen>>>")
			http.Redirect(w, r, "/admin_user/slider", 301)
			return
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		sliderName := r.PostFormValue("sliderName")
		description := r.PostFormValue("description")
		message := r.PostFormValue("message")

		if message != "" && description != "" {
			sliderObject := slider.Slider{"", sliderName, description, misc.ConvertToByteArray(message), content}
			_, err := slider_io.CreateSlider(sliderObject)
			if err != nil {
				fmt.Println("error create slider")
			}
		}
		http.Redirect(w, r, "/admin_user/slider", 301)
	}
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Session result: ", adminHelper.CheckAdminInSession(app, r))
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
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
		var sliderList []slider.SliderHelper
		sliders, err := slider_io.ReadSliders()
		if err != nil {
			fmt.Println("error reading Sliders")
		} else {
			for _, mySlider := range sliders {
				sliderOblect := slider.SliderHelper{mySlider.Id, mySlider.SliderName, mySlider.Description, misc.ConvertingToString(mySlider.SliderMessage), misc.ConvertingToString(mySlider.SliderImage)}
				sliderList = append(sliderList, sliderOblect)
				//sliderOblect = slider.SliderHelper{}
			}

		}
		type PageData struct {
			Sliders       []slider.SliderHelper
			Backend_error string
			Unknown_error string
			SidebarData   misc.SidebarData
		}
		data := PageData{sliderList, backend_error, unknown_error, misc.GetSideBarData("slider", "")}
		files := []string{
			app.Path + "admin/slider/slider.html",
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
