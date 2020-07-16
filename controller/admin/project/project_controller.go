package project

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	image2 "ostmfe/domain/image"
	project2 "ostmfe/domain/project"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/project_io"
	"time"
)

func ProjectHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", ProjectsHandler(app))
	r.Get("/new", NewProjectsHandler(app))
	r.Get("/new_history/{projectId}", NewProjectHistoryHandler(app))
	r.Get("/edit/{projectId}", EditeProjectsHandler(app))
	r.Post("/create", CreateProjectHandler(app))
	r.Post("/create_project_history", CreateProjectHistoryHandler(app))
	r.Post("/update_pictures", ProjectUpdatePicturesHandler(app))
	r.Post("/update_picture", ProjectUpdatePictureHandler(app))
	r.Post("/update_history", ProjectUpdateHistoryHandler(app))
	return r
}

//TODO for now we are accepting mytextarea maybe empty.
func ProjectUpdateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		historyContent := r.PostFormValue("mytextarea")
		projectId := r.PostFormValue("projectId")
		historyId := r.PostFormValue("historyId")
		//checking if the projectHistory exists
		history, err := history_io.ReadHistory(historyId)
		fmt.Println(historyId)
		if err != nil {
			//TODO need to create a new entity called Histories. will have to change the logic of history of each entity that is related to this class.
			fmt.Println(err, " could not read history")
			fmt.Println(err, " proceeding into creation of a project history.....")
			history := history2.History{""}

			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		}
		updatedProjectHistory := history2.History{history.Id, history.Title, history.Description, misc.ConvertToByteArray(historyContent), history.Date}
		//Now Updating
		_, errr := history_io.UpdateHistory(updatedProjectHistory)
		if errr != nil {
			fmt.Println(errr, " could not update history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		}
		fmt.Println(" successfully updated")
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return
	}
}

//Todo implement the HTML page to get these data
func EditeProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")
		selectedProjest := misc.GetProjectEditable(projectId)
		projectDetails, err := project_io.ReadProject(projectId)

		if err != nil {
			fmt.Println(err, " Error reading project Details")
		}
		type PageData struct {
			Project        misc.ProjectEditable
			ProjectDetails project2.Project
		}
		data := PageData{selectedProjest, projectDetails}
		files := []string{
			app.Path + "admin/project/edite_project.html",
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

func NewProjectsHandler(app *config.Env) http.HandlerFunc {
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
		type PagePage struct {
			Backend_error string
			Unknown_error string
		}
		data := PagePage{backend_error, unknown_error}
		files := []string{
			app.Path + "admin/project/new_project.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
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

func ProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading projects")
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
		type PagePage struct {
			Backend_error string
			Unknown_error string
			Projects      []project2.Project
		}
		data := PagePage{backend_error, unknown_error, projects}

		files := []string{
			app.Path + "admin/project/projects.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
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
func CreateProjectHandler(app *config.Env) http.HandlerFunc {
	/***
	Here we create a new project

	*/
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		project_name := r.PostFormValue("project_name")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>>>>")
		}

		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		fmt.Println(project_name, "<<<Project Name|| description>>>", description)

		if project_name != "" && description != "" {
			project := project2.Project{"", project_name, description}
			new_project, err := project_io.CreateProject(project)
			if err != nil {
				fmt.Println(err, " could not create project Line: 190")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new", 301)
				return
			}

			projectImage := project2.ProjectImage{"", new_project.Id, "", ""}
			helper := project2.ProjectImageHelper{filesByteArray, projectImage}
			_, errr := project_io.CreateProjectImage(helper)
			if errr != nil {
				fmt.Println(errr, " error creating projectImage")
				_, err := project_io.DeleteProject(new_project.Id)
				if err != nil {
					fmt.Println(err, " error deleting project")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+project_name)
			http.Redirect(w, r, "/admin_user/project/new_history/"+new_project.Id, 301)
			return
			//event_name := r.PostFormValue("event_name")
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/project/new", 301)
		return

	}

}
func CreateProjectHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		projectId := r.PostFormValue("projectId")
		history := r.PostFormValue("mytextarea")
		description := r.PostFormValue("description")
		title := r.PostFormValue("title")

		if projectId != "" && history != "" {
			project, err := project_io.ReadProject(projectId)
			if err != nil {
				fmt.Println(err, " error reading Project")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new_history/"+projectId, 301)
				return
			}
			historyByteArray := []byte(history)
			historyObject := history2.History{"", title, description, historyByteArray, time.Now()}
			history, err := history_io.CreateHistory(historyObject)
			if err != nil {
				fmt.Println(err, " error creating History")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new_history/"+projectId, 301)
				return
			}
			projectHistoryObject := project2.ProjectHistory{"", projectId, history.Id}
			_, errr := project_io.CreateProjectHistory(projectHistoryObject)
			if errr != nil {
				_, err := history_io.DeleteHistory(history.Id)
				if err != nil {
					fmt.Println(err, " error Delete History")
				}
				fmt.Println(err, " error creating History")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/new_history/"+projectId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new People Type : "+project.Title)
			http.Redirect(w, r, "/admin_user", 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/project/new_history/"+projectId, 301)
		return
	}
}

func NewProjectHistoryHandler(app *config.Env) http.HandlerFunc {
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
		projectId := chi.URLParam(r, "projectId")
		project, err := project_io.ReadProject(projectId)
		if err != nil {
			fmt.Println(" error reading project")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/new", 301)
			return
		}

		type PageData struct {
			Project       project2.Project
			Backend_error string
			Unknown_error string
		}
		data := PageData{project, backend_error, unknown_error}
		files := []string{
			app.Path + "admin/project/new_project_history.html",
			app.Path + "admin/template/navbar.html",
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
func ProjectUpdatePictureHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//fileslist := r.Form["file"]

		file, _, err := r.FormFile("file")
		imageId := r.PostFormValue("imageId")
		decription := r.PostFormValue("decription")
		projectId := r.PostFormValue("projectId")
		fmt.Println(projectId)
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		} else if file != nil {
			reader := bufio.NewReader(file)
			content, _ := ioutil.ReadAll(reader)
			//First reading the project to make sure that the project that we want to update exist.
			new_project, err := project_io.ReadProject(projectId)
			if err != nil {
				fmt.Println(err, " could not read project Line: 113")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project", 301)
				return
			}
			image := image2.Images{imageId, content, decription}
			_, errr := image_io.UpdateImage(image)
			if errr != nil {
				fmt.Println(err, "Error Updating image")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following project : "+new_project.Title)
			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return
	}
}

func ProjectUpdatePicturesHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//fileslist := r.Form["file"]

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		projectId := r.PostFormValue("projectId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		new_project, err := project_io.ReadProject(projectId)
		if err != nil {
			fmt.Println(err, " could not read project Line: 113")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/new", 301)
			return
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)
		projectImage := project2.ProjectImage{"", new_project.Id, "", ""}
		helper := project2.ProjectImageHelper{filesByteArray, projectImage}
		_, errr := project_io.CreateProjectImage(helper)
		if errr != nil {
			fmt.Println(errr, " error creating projectImage")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/edit/"+new_project.Id, 301)
			return
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following project : "+new_project.Title)
		http.Redirect(w, r, "/admin_user/project/new_history/"+new_project.Id, 301)
		return
	}
}
