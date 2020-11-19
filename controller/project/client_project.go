package project

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	comment2 "ostmfe/controller/comment"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	project2 "ostmfe/domain/project"
	"ostmfe/io/comment_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/project_io"
	"time"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/read_single/{projectId}", ReadSingleProjectHanler(app))
	r.Post("/comment/{projectId}", createProjectComment(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func createProjectComment(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")
		r.ParseForm()
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		Subject := r.PostFormValue("Subject")
		message := r.PostFormValue("message")

		if name != "" && email != "" && message != "" && Subject != "" {
			commentObject := comment.Comment{"", email, name, misc.FormatDateTime(time.Now()), misc.ConvertToByteArray(message), ""}
			newComment, err := comment_io.CreateComment(commentObject)
			if err != nil {
				fmt.Println("error creating comment")
			} else {
				_, err := comment_io.CreateCommentProject(comment.CommentProject{"", projectId, newComment.Id})
				if err != nil {
					fmt.Println("error creating comment")
				}
			}
		}
		http.Redirect(w, r, "/project/read_single/"+projectId, 301)
	}
}

func ReadSingleProjectHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "projectId")
		var projectDataHistory ProjectDataHistory

		if projectId != "" {
			projectDataHistory = getProjectDataHistory(projectId)
		}
		//read All the projects
		projects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading projects")
		}
		commentNumber, err := comment_io.CountProjectComment(projectId)
		if err != nil {
			fmt.Println(err, " error reading projects comment Number")
		}
		type PageData struct {
			ProjectDataHistory ProjectDataHistory
			Projects           []project2.Project
			Comments           []comment.CommentStack
			CommentNumber      int64
		}
		data := PageData{projectDataHistory, projects, comment2.GetProjectComment(projectId), commentNumber}

		files := []string{
			app.Path + "project/project_single.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
			app.Path + "base_templates/comments.html",
			app.Path + "base_templates/reply-template.html",
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
		type PageData struct {
			Projects []misc.ProjectContentsHome
		}
		data := PageData{misc.GetProjectContentsHomes()}
		files := []string{
			app.Path + "project/projects.html",
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

type ProjectDataHistory struct {
	Project      project2.Project
	ProfileImage image.Images
	Images       []image.Images
	History      history2.HistoriesHelper
}

func getProjectDataHistory(projectId string) ProjectDataHistory {
	var projectDataHistory ProjectDataHistory
	var profileImage image.Images
	var images []image.Images

	project, err := project_io.ReadProject(projectId)
	if err != nil {
		fmt.Println(err, " error has occurred when reading ")
		return projectDataHistory
	}
	//Images
	projectImages, err := project_io.ReadWithProjectIdProjectImages(projectId)
	if err != nil {
		fmt.Println(err, " error has occurred when reading ")
		return projectDataHistory
	}
	for _, projectImage := range projectImages {
		if projectImage.ImageType == "1" || projectImage.ImageType == "profile" {
			profileImage, err = image_io.ReadImage(projectImage.ImageId)
			if err != nil {
				fmt.Println(err, " error has occurred when reading image")
			}
		}
		image, err := image_io.ReadImage(projectImage.ImageId)
		if err != nil {
			fmt.Println(err, " error has occurred when reading image")
		}
		images = append(images, image)
	}
	//History
	fmt.Println(project.Id)
	projectHistory, err := project_io.ReadProjectHistoryOf(project.Id)
	if err != nil {
		fmt.Println(err, " error has occurred when reading project History")
	}
	history, err := history_io.ReadHistorie(projectHistory.HistoryId)
	if err != nil {
		fmt.Println(err, " error has occurred when reading History")
	}
	historyHelp := history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}
	projectDataHistory = ProjectDataHistory{project, profileImage, images, historyHelp}

	return projectDataHistory
}

//type ProjectContentsHome struct {
//	ProjectId   string
//	Title       string
//	Picture     string
//	Description string
//}
//func getProjectContentsHomes() []ProjectContentsHome {
//	projectContentsHomeObject := []ProjectContentsHome{}
//	image := image.Images{}
//	projects, err := project_io.ReadProjects()
//	if err != nil {
//		fmt.Println(err, " Error reading all the projects")
//		return projectContentsHomeObject
//	}
//	for _, project := range projects {
//		//fmt.Println(project.Title)
//		projectImage, err := project_io.ReadWithProjectIdProjectImage(project.Id)
//		if err != nil {
//			fmt.Println(err, " Can not find the following project in project image table: ", project.Title)
//		} else {
//			image, err = image_io.ReadImage(projectImage.ImageId)
//			//fmt.Println(image.Image)
//			if err != nil {
//				fmt.Println(err, " Can not find the following project image Id in Image table: ", projectImage.ImageId)
//			}
//		}
//		projectObject := ProjectContentsHome{project.Id, project.Title, image.Id, project.Description}
//		projectContentsHomeObject = append(projectContentsHomeObject, projectObject)
//		projectObject = ProjectContentsHome{}
//	}
//	return projectContentsHomeObject
//}
