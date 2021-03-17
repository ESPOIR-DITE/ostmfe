package classroom

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	classroom2 "ostmfe/domain/classroom"
	"ostmfe/io/classroom"
)

func ClassHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", classHomeHandler(app))
	r.Get("/edit/{classroomId}", editClassHomeHandler(app))
	r.Post("/create", createClassroom(app))
	r.Post("/update_details", updateClassroom(app))
	r.Post("/update_image", updateImageClassroom(app))

	return r
}

func updateImageClassroom(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		file, _, err := r.FormFile("file")
		classroomId := r.PostFormValue("classroomId")

		if err != nil {
			fmt.Println(err, "<<<error reading file>>>>This error may happen if there is no picture selected>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}

		if classroomId != "" {
			classRoom, err := classroom.ReadClassroom(classroomId)
			if err != nil {
				fmt.Println(err, " error updating classroom")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/classroom", 301)
				return
			}
			classroomObject := classroom2.Classroom{classroomId, classRoom.Name, classRoom.Description, classRoom.Details, content}
			classroomResturn, err := classroom.UpdateClassroom(classroomObject)
			if err != nil {
				fmt.Println(err, " error updating classroom")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/classroom", 301)
				return
			}
			fmt.Println(err, " creation successful classroom")

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully update an new classroom : "+classroomResturn.Name)
			http.Redirect(w, r, "/admin_user/classroom/", 301)
			return
		}
		fmt.Println("One of the field is missing when creating classroom")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/classroom", 301)
		return

	}
}

func updateClassroom(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("name")
		classroomId := r.PostFormValue("classroomId")
		description := r.PostFormValue("description")
		details := r.PostFormValue("details")

		if name != "" && description != "" && details != "" && classroomId != "" {
			classRoom, err := classroom.ReadClassroom(classroomId)
			if err != nil {
				fmt.Println(err, " error updating classroom")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/classroom", 301)
				return
			}
			classroomObject := classroom2.Classroom{classroomId, name, description, misc.ConvertToByteArray(details), classRoom.Icon}
			classroomResturn, err := classroom.UpdateClassroom(classroomObject)
			if err != nil {
				fmt.Println(err, " error updating classroom")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/classroom", 301)
				return
			}
			fmt.Println(err, " creation successful classroom")

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully update an new classroom : "+classroomResturn.Name)
			http.Redirect(w, r, "/admin_user/classroom/", 301)
			return
		}
		fmt.Println("One of the field is missing when creating classroom")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/classroom", 301)
		return

	}
}

func editClassHomeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		classroomId := chi.URLParam(r, "classroomId")

		classroom, err := getClassroom(classroomId)
		if err != nil {
			fmt.Println(err, " error reading classroom")
			http.Redirect(w, r, "/admin_user/classroom", 301)
			return
		}
		type PageData struct {
			Classroom classroom2.ClassroomHelper
		}

		data := PageData{classroom}

		files := []string{
			app.Path + "admin/classroom/edit_classroom.html",
			app.Path + "admin/template/navbar.html",
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

func createClassroom(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var content []byte
		file, _, err := r.FormFile("file")
		name := r.PostFormValue("name")
		description := r.PostFormValue("description")
		details := r.PostFormValue("details")

		if err != nil {
			fmt.Println(err, "<<<error reading file>>>>This error may happen if there is no picture selected>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}

		if name != "" && description != "" && details != "" {
			classroomObject := classroom2.Classroom{"", name, description, misc.ConvertToByteArray(details), content}
			classroom, err := classroom.CreateClassroom(classroomObject)
			if err != nil {
				fmt.Println(err, " error creating classroom")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/classroom", 301)
				return
			}
			fmt.Println(err, " creation successful classroom")

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new classroom : "+classroom.Name)
			http.Redirect(w, r, "/admin_user/classroom/", 301)
			return
		}
		fmt.Println("One of the field is missing when creating classroom")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/classroom", 301)
		return

	}
}

func classHomeHandler(app *config.Env) http.HandlerFunc {
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

		classrooms, err := classroom.ReadClassrooms()
		if err != nil {
			fmt.Println(err, " error reading classrooms")
		}
		type PagePage struct {
			Classroom     []classroom2.Classroom
			Backend_error string
			Unknown_error string
			SidebarData   misc.SidebarData
		}
		data := PagePage{classrooms,
			backend_error,
			unknown_error,
			misc.GetSideBarData("classroom", "")}
		files := []string{
			app.Path + "admin/classroom/classroom.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
			app.Path + "admin/template/cards.html",
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

func getClassroom(classroomId string) (classroom2.ClassroomHelper, error) {
	var myClassroom classroom2.ClassroomHelper

	classroomObject, err := classroom.ReadClassroom(classroomId)
	if err != nil {
		fmt.Println(err, " error reading classroom")
		return myClassroom, err
	}
	myClassroom = classroom2.ClassroomHelper{classroomObject.Id, classroomObject.Name, classroomObject.Description, misc.ConvertingToString(classroomObject.Details), misc.ConvertingToString(classroomObject.Icon)}

	return myClassroom, nil
}
