package project

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
	"ostmfe/domain/comment"
	"ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	image2 "ostmfe/domain/image"
	"ostmfe/domain/place"
	project2 "ostmfe/domain/project"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/place_io"
	"ostmfe/io/project_io"
	"ostmfe/utile"
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
	r.Post("/update_details", ProjectUpdateDetails(app))

	r.Post("/create_history", ProjectCreateHistoryHandler(app))
	r.Post("/addPlace", AddPlaceHandler(app))
	//Gallery
	r.Post("/create-gallery", CreateGalleryHandler(app))
	r.Get("/delete-gallery/{pictureId}/{projectId}/{projectGalleryPictureId}", DeleteGalleryHandler(app))

	r.Get("/delete_project/{projectId}", DeleteProjectHandler(app))
	r.Get("/activate_comment/{commentId}/{projectId}", ActivateCommentHandler(app))

	r.Post("/create-page-flow", CreatePageFlowHandler(app))
	r.Get("/delete-pageFlow/{projectPageFlowId}/{projectId}", DeletePageFlowHandler(app))
	return r
}

func DeletePageFlowHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		projectPageFlowId := chi.URLParam(r, "projectPageFlowId")
		projectId := chi.URLParam(r, "projectId")

		_, err := project_io.DeleteProjectPageFLow(projectPageFlowId)
		if err != nil {
			fmt.Println("error deleting History Page FLow")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: Project Gallery. ")
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return
	}
}

func CreatePageFlowHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		projectId := r.PostFormValue("projectId")
		pageFlowTitle := r.PostFormValue("pageFlowTitle")
		scr := r.PostFormValue("scr")

		if scr != "" && projectId != "" && pageFlowTitle != "" {
			_, err := project_io.CreateProjectPageFLow(project2.ProjectPageFlow{"", pageFlowTitle, projectId, scr})
			if err != nil {
				fmt.Println(err, " error creating page flow!")
			}
		} else {
			app.ErrorLog.Print("Error creating projectPageFlow")
		}
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return
	}
}

//Todo finish this method.
func AddPlaceHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		projectId := r.PostFormValue("projectId")
		PlaceId := r.PostFormValue("PlaceId")
		if PlaceId != "" && projectId != "" {

		}
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return
	}
}
func ActivateCommentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		commentId := chi.URLParam(r, "commentId")
		projectId := chi.URLParam(r, "projectId")
		result := misc.ActivateComment(commentId)
		fmt.Print("Activation Result: ", result)
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return
	}
}
func DeleteGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		pictureId := chi.URLParam(r, "pictureId")
		projectId := chi.URLParam(r, "projectId")
		projectGalleryPictureId := chi.URLParam(r, "projectGalleryPictureId")

		//Deleting project
		gallery, err := image_io.DeleteGalery(pictureId)
		if err != nil {
			fmt.Println("error deleting gallery")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		} else {
			_, err := project_io.DeleteProjectGallery(projectGalleryPictureId)
			if err != nil {

				fmt.Println("ROLLING BACK!!!")
				_, err := image_io.UpdateGallery(gallery)
				if err != nil {
					fmt.Println("error updating gallery")
				}

				fmt.Println("error deleting project gallery")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
				return
			}
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: Project Gallery. ")
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return
	}
}

func CreateGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		project := r.PostFormValue("project")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if project != "" && description != "" {
			galery := image2.Galery{"", content, description}
			galleryObject, err := image_io.CreateGalery(galery)
			if err != nil {
				fmt.Println(err, " error creating gallery")
			} else {
				placeGallery := project2.ProjectGallery{"", project, galleryObject.Id}
				_, err := project_io.CreateProjectGallery(placeGallery)
				if err != nil {
					fmt.Println(err, " error creating projectGallery")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/project/edit/"+project, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/project/edit/"+project, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/project/edit/"+project, 301)
		return
	}
}

func DeleteProjectHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		projectId := chi.URLParam(r, "projectId")

		//Reading the project
		project, err := project_io.ReadProject(projectId)
		if err != nil {
			fmt.Println("error reading project")
			fmt.Println(err, " error creating Project")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/project", 301)
			return
		}
		//Deleting project
		_, err = project_io.DeleteProject(project.Id)
		if err != nil {
			fmt.Println("error deleting project")
		}

		//Reading Project Image
		projectImage, err := project_io.ReadWithProjectIdProjectImage(projectId)
		if err != nil {
			fmt.Println("error deleting projectImage. this project may not have an image")
		} else {
			_, err = project_io.DeleteProjectImage(projectImage.Id)
			if err != nil {
				fmt.Println("error deleting project Image")
			} else {
				_, err = image_io.DeleteImage(projectImage.Id)
				if err != nil {
					fmt.Println("error deleting image")
				}
			}
		}

		//Reading History
		projectHistory, err := project_io.ReadProjectHistoryOf(projectId)
		if err != nil {
			fmt.Println("error reading projectHistory. this project may not have History")
		} else {
			_, err = history_io.DeleteHistorie(projectHistory.HistoryId)
			if err != nil {
				fmt.Println("error deleting history. this project may not have an history")
			}
			_, err := project_io.DeleteProjectHistory(projectHistory.Id)
			if err != nil {
				fmt.Println("error deleting projectHistory")
			}
		}

		//Reading ProjectMembers
		projectMembers, err := project_io.ReadAllOfProjectMembers(projectId)
		if err != nil {
			fmt.Println("error reading projectMembers. this project may not have an Members")
		} else {
			for _, projectMember := range projectMembers {
				_, err = project_io.DeleteProjectMember(projectMember.Id)
				if err != nil {
					fmt.Println("error deleting project Member for the following projectMemberId: ", projectMember.Id)
				}
			}
		}
		//Reading ProjectPartners
		ProjectPartners, err := project_io.ReadAllOfProjectPartner(projectId)
		if err != nil {
			fmt.Println("error reading Project Partners. this project may not have an Partners")
		} else {
			for _, projectPartner := range ProjectPartners {
				_, err = project_io.DeleteProjectMember(projectPartner.Id)
				if err != nil {
					fmt.Println("error deleting project Partner for the following projectPartnerId: ", projectPartner.Id)
				}
			}
		}

		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: Project Details. ")
		http.Redirect(w, r, "/admin_user/project", 301)
		return
	}
}

func ProjectUpdateDetails(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		projectTitle := r.PostFormValue("projectTitle")
		projectId := r.PostFormValue("projectId")
		description := r.PostFormValue("description")

		_, err := project_io.ReadProject(projectId)
		if err != nil {
			fmt.Println("error reading project")
			fmt.Println(err, " error creating Project")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		}
		if projectId != "" && projectTitle != "" && description != "" {
			projectObject := project2.Project{projectId, projectTitle, description}
			_, errs := project_io.UpdateProject(projectObject)
			if errs != nil {
				fmt.Println("error updating project")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
				http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
				return
			}

			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: Project Details. ")
			http.Redirect(w, r, "/admin_user/project", 301)
			return
		}
		fmt.Println(" error creating project One field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred due to selected Place, Please try again late")
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return

	}
}

func ProjectCreateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		projectId := r.PostFormValue("projectId")
		//checking if there is contents in the variables
		if historyContent != "" && projectId != "" {
			history := history2.Histories{"", misc.ConvertToByteArray(historyContent)}

			newHistory, err := history_io.CreateHistorie(history)
			if err != nil {
				fmt.Println(err, " something went wrong! could not create history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
				return
			}
			fmt.Println("HistoryId created successfully ..")
			fmt.Println(" proceeding into creation of a project_history.....")
			projectHistory := project2.ProjectHistory{"", projectId, newHistory.Id}
			_, errr := project_io.CreateProjectHistory(projectHistory)
			if errr != nil {
				fmt.Println(err, " could not create ProjectHistory")
				fmt.Println("RollBack ...")
				fmt.Println("deleting history ...")
				_, err := history_io.DeleteHistorie(newHistory.Id)
				if err != nil {
					fmt.Println("Error deleting history ...")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
				return
			}
			fmt.Println(" successfully created")
			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		}
		fmt.Println("one or more fields missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return

	}
}

//TODO for now we are accepting mytextarea maybe empty. but this will will need to be investigates
func ProjectUpdateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		projectId := r.PostFormValue("projectId")
		historyId := r.PostFormValue("historyId")
		//checking if the projectHistory exists
		_, err := history_io.ReadHistorie(historyId)
		fmt.Println(historyContent)
		if err != nil {
			fmt.Println(err, " something went wrong! could not read history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		}
		histories := history2.Histories{historyId, misc.ConvertToByteArray(historyContent)}

		_, errr := history_io.UpdateHistorie(histories)
		if errr != nil {
			fmt.Println(err, " something went wrong! could not update history")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
			return
		}
		project, errx := project_io.ReadProject(projectId)
		if errx != nil {
			fmt.Println("error reading project")
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		fmt.Println(" successfully updated")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: "+project.Title+"  project. ")
		http.Redirect(w, r, "/admin_user/project/edit/"+projectId, 301)
		return
	}
}
func EditeProjectsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		projectId := chi.URLParam(r, "projectId")

		selectedProjest := misc.GetProjectEditable(projectId)
		projectDetails, err := project_io.ReadProject(projectId)
		if err != nil {
			fmt.Println(err, " Error reading project Details")
		}
		Places, err := place_io.ReadPlaces()
		if err != nil {
			fmt.Println(err, " Error reading places")
		}
		events, err := event_io.ReadEvents()
		if err != nil {
			fmt.Println(err, " Error reading events")
		}
		projectPageFlow, err := project_io.ReadProjectPageFLowsWithProjectId(projectId)
		if err != nil {
			fmt.Println(err, " Error reading pageFlows")
		}
		commentNumber, pendingcomments, activeComments := projectCommentCalculation(projectId)
		type PageData struct {
			Events          []event.Event
			Places          []place.Place
			Project         misc.ProjectEditable
			ProjectDetails  project2.Project
			SidebarData     misc.SidebarData
			Comments        []comment.CommentHelper2
			Gallery         []misc.ProjectGalleryImages
			CommentNumber   int64
			PendingComments int64
			ActiveComments  int64
			PageFlows       []project2.ProjectPageFlow
		}
		data := PageData{events,
			Places,
			selectedProjest,
			projectDetails,
			misc.GetSideBarData("project", ""),
			GetProjectCommentsWithProjectId(projectId),
			misc.GetProjectGallery(projectId),
			commentNumber, pendingcomments, activeComments,
			projectPageFlow,
		}
		files := []string{
			app.Path + "admin/project/edite_project.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/cards.html",
			app.Path + "admin/template/topbar.html",
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
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
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
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
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
			SidebarData   misc.SidebarData
		}
		data := PagePage{backend_error, unknown_error, projects, misc.GetSideBarData("project", "")}

		files := []string{
			app.Path + "admin/project/projects.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
			app.Path + "base_templates/footer.html",
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
func CreateProjectHandler(app *config.Env) http.HandlerFunc {
	/***
	Here we create a new project

	*/
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		project_name := r.PostFormValue("project_name")
		description := r.PostFormValue("description")
		mytextarea := r.PostFormValue("mytextarea")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>>>>")
		}

		if err != nil {
			fmt.Println(err, "<<<error reading file>>>>This error may happen if there is no picture selected>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}

		fmt.Println(project_name, "<<<Project Name|| description>>>", description)

		if project_name != "" && description != "" && mytextarea != "" {
			project := project2.Project{"", project_name, description}
			new_project, err := project_io.CreateProject(project)
			if err != nil {
				fmt.Println(err, " could not create project Line: 190")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project", 301)
				return
			}

			//project Image
			//Image
			imageObject := image2.Images{"", content, utile.PROFILE}
			imageObjectNew, err := image_io.CreateImage(imageObject)
			if err != nil {
				fmt.Println(err, " error creating a new image")
			}
			projectImage := project2.ProjectImage{"", new_project.Id, imageObjectNew.Id, utile.PROFILE}
			_, errr := project_io.CreateProjectImage(projectImage)
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
				http.Redirect(w, r, "/admin_user/project", 301)
				return
			}

			historyObject := history2.Histories{"", misc.ConvertToByteArray(mytextarea)}
			history, err := history_io.CreateHistorie(historyObject)
			if err != nil {
				fmt.Println(err, " error reading History")
			} else {
				projectHistoryObject := project2.ProjectHistory{"", new_project.Id, history.Id}
				_, err := project_io.CreateProjectHistory(projectHistoryObject)
				if err != nil {
					fmt.Println(err, " error reading project History")
				}
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new project : "+project_name)
			http.Redirect(w, r, "/admin_user/project", 301)
			return
			//event_name := r.PostFormValue("event_name")
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/project", 301)
		return

	}

}
func CreateProjectHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		projectId := r.PostFormValue("projectId")
		history := r.PostFormValue("mytextarea")
		//description := r.PostFormValue("description")
		//title := r.PostFormValue("title")

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

			historyObject := history2.Histories{"", historyByteArray}
			history, err := history_io.CreateHistorie(historyObject)
			if err != nil {
				fmt.Println(err, " error creating HistoryId")
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
					fmt.Println(err, " error Delete HistoryId")
				}
				fmt.Println(err, " error creating HistoryId")
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
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		var unknown_error string
		var backend_error string
		projects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading projects")
		}
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
			Projects      []project2.Project
			Backend_error string
			Unknown_error string
		}
		data := PageData{project, projects, backend_error, unknown_error}
		files := []string{
			//app.Path + "admin/project/new_project_history.html",
			app.Path + "admin/project/projects_history.html",
			app.Path + "admin/template/topbar.html",
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
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
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
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()

		var content []byte
		file, _, err := r.FormFile("file")
		projectId := r.PostFormValue("projectId")
		imageId := r.PostFormValue("imageId")
		//projectImageId := r.PostFormValue("projectImageId")
		//imageType := r.PostFormValue("imageType")
		if err != nil {
			fmt.Println(err, "<<<error reading file>>>>This error may happen if there is no picture selected>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		project, err := project_io.ReadProject(projectId)
		if err != nil {
			fmt.Println(err, " could not read project Line: 113")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/project/", 301)
			return
		}
		if imageId != "" && projectId != "" {
			newImageObject := image2.Images{imageId, content, ""}
			_, err := image_io.UpdateImage(newImageObject)
			if err != nil {
				fmt.Println(err, " error updating image")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/project/edit/"+project.Id, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following project : "+project.Title)
			http.Redirect(w, r, "/admin_user/project/edit/"+project.Id, 301)
			return
		}
	}
}
