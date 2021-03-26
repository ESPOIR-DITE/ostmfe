package project

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
	"ostmfe/domain/comment"
	contribution2 "ostmfe/domain/contribution"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	project2 "ostmfe/domain/project"
	"ostmfe/io/comment_io"
	"ostmfe/io/contribution_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/pageData_io"
	"ostmfe/io/project_io"
	"path/filepath"
	"strings"
	"time"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/read_single/{projectId}", ReadSingleProjectHanler(app))
	r.Post("/comment", createProjectComment(app))
	r.Post("/contribution", CreateContributionComment(app))
	//r.Use(middleware.LoginSession{SessionManager: app.Session}.RequireAuthenticatedUser)
	//r.Get("/home", indexHanler(app))
	//r.Get("/homeError", indexErrorHanler(app))

	return r
}

func CreateContributionComment(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		var isExtension bool
		var fileTypeId string
		r.ParseForm()
		file, m, err := r.FormFile("file")
		projectId := r.PostFormValue("projectId")
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
				http.Redirect(w, r, "/project/read_single/"+projectId, 301)
			}
		}
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		cellphone := r.PostFormValue("cellphone")

		if name != "" && email != "" && message != "" && projectId != "" {
			contributionObject := contribution2.Contribution{"", email, name, time.Now(), cellphone, misc.ConvertToByteArray(message)}

			contribution, err := contribution_io.CreateContribution(contributionObject)
			if err != nil {
				fmt.Println("error creating a new contribution")
			} else {
				contributorEventObject := contribution2.ContributionProject{"", contribution.Id, projectId, name}
				_, err := contribution_io.CreateContributionProject(contributorEventObject)
				if err != nil {
					_, _ = contribution_io.DeleteContribution(contribution.Id)
					fmt.Println("error creating a new contribution")
				} else {
					contributionFileObject := contribution2.ContributionFile{"", contribution.Id, content, fileTypeId, ""}
					_, err := contribution_io.CreateContributionFile(contributionFileObject)
					if err != nil {
						fmt.Println("error creating contributionFile")
					}
				}
			}
		}
		http.Redirect(w, r, "/event/read_single/"+projectId, 301)
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
func createProjectComment(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		message := r.PostFormValue("message")
		projectId := r.PostFormValue("projectId")

		if name != "" && email != "" && message != "" && projectId != "" {
			commentObject := comment.Comment{"", email, name, misc.FormatDateTime(time.Now()), misc.ConvertToByteArray(message), "", false}
			newComment, err := comment_io.CreateComment(commentObject)
			if err != nil {
				fmt.Println("error creating comment")
			} else {
				projectCommentObject := comment.CommentProject{"", projectId, newComment.Id}
				_, err := comment_io.CreateCommentProject(projectCommentObject)
				if err != nil {
					fmt.Println("error creating comment")
				}
			}
			http.Redirect(w, r, "/project/read_single/"+projectId, 301)
			return
		}
		fmt.Println("error one field missing")
		http.Redirect(w, r, "/project/read_single/"+projectId, 301)
		return
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
		projectComment := comment2.GetProjectComment(projectId)
		//fmt.Println(" projectComment: ",projectComment)
		type PageData struct {
			ProjectDataHistory ProjectDataHistory
			Projects           []project2.Project
			Comments           []comment.CommentStack
			CommentNumber      int64
			//GalleryString      []string
			GalleryImages []misc.ProjectGalleryImages
		}
		data := PageData{projectDataHistory, projects, projectComment, commentNumber, misc.GetProjectGallery(projectId)}

		files := []string{
			app.Path + "project/project_single.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
			app.Path + "base_templates/comments.html",
			//app.Path + "base_templates/reply-template.html",
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
		var bannerImage string
		pageBanner, err := pageData_io.ReadPageBannerWIthPageName("project-page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = misc.GetBannerImage(pageBanner.BannerId)
		}

		type PageData struct {
			Projects      []misc.ProjectContentsHome
			ProjectBanner string
		}
		data := PageData{misc.GetProjectContentsHomes(), bannerImage}
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

func getProjectGallery(projectId string) []string {
	var picture []string
	projectGallerys, err := project_io.ReadAllByProjectIdGallery(projectId)
	if err != nil {
		fmt.Println(err, " error peopleGalleries.")
	} else {
		for _, projectGallery := range projectGallerys {
			gallery, err := image_io.ReadGallery(projectGallery.GalleryId)
			if err != nil {
				fmt.Println(err, " error gallery")
			} else {
				picture = append(picture, misc.ConvertingToString(gallery.Image))
			}
		}
	}
	return picture
}
