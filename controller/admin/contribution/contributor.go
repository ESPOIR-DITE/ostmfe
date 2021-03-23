package contribution

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	"ostmfe/domain/contribution"
	"ostmfe/domain/place"
	"ostmfe/io/contribution_io"
	"ostmfe/io/place_io"
	"time"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Get("/contribution", homeContributionHanler(app))
	r.Get("/ipAddress", printIpAddressHandler(app))
	r.Get("/delete-contribution/{contributionId}/{contributionFile}", DeleteContributionHandler(app))
	r.Get("/delete-file-type/{fileTypeId}", DeleteFileTypeHanler(app))
	r.Get("/delete-place-category/{placeCategory}", DeletePlaceCategoryHandler(app))
	r.Get("/read-audio/{fileId}", ReadAudioFileHandler(app))
	r.Post("/createFileType", CreateFileType(app))
	r.Post("/new", CreateContribution(app))
	r.Post("/place-category", CreatePlaceCategoryHandler(app))

	return r
}

func DeletePlaceCategoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		fileTypeId := chi.URLParam(r, "placeCategory")

		if fileTypeId != "" {
			_, err := place_io.DeletePlaceCategory(fileTypeId)
			if err != nil {
				fmt.Println(err, " error deleting place category.")
			}
		}
		http.Redirect(w, r, "/admin_user/contribution", 301)
		return
	}
}

func CreatePlaceCategoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		r.ParseForm()

		placeCategory := r.FormValue("placeCategory")

		if placeCategory != "" {
			placeCategoryObject := place.PlaceCategory{"", placeCategory}
			_, err := place_io.CreatePlaceCategory(placeCategoryObject)
			if err != nil {
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, please try again later.")
				fmt.Println(err, " error creating category.")
			} else {
				app.InfoLog.Println(" successfully created placeCategory")
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new PlaceType : ")
			}
		} else {
			fmt.Println(" error creating category. Field missing.")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, The field is empty.")
		}
		http.Redirect(w, r, "/admin_user/contribution", 301)
		return
	}
}

func printIpAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		host, _ := os.Hostname()
		addrs, _ := net.LookupIP(host)
		for _, addr := range addrs {
			if ipv4 := addr.To4(); ipv4 != nil {
				fmt.Println("IPv4: ", ipv4)
			}
		}
		http.Redirect(w, r, "/admin_user/contribution/contribution", 301)
	}
}

func DeleteContributionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		contributionId := chi.URLParam(r, "contributionId")
		contributionFile := chi.URLParam(r, "contributionFile")

		contribution, err := contribution_io.DeleteContribution(contributionId)
		if err != nil {
			fmt.Println(err, " error deleting contribution")
		} else {
			_, err := contribution_io.DeleteContributionFile(contributionFile)
			if err != nil {
				fmt.Println(err, " error deleting contribution")
			} else {
				_, err := contribution_io.UpdateContribution(contribution)
				if err != nil {
					fmt.Println(err, " error updating contribution")
				}
			}
		}
		http.Redirect(w, r, "/admin_user/contribution", 301)
	}
}

func ReadAudioFileHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		fileId := chi.URLParam(r, "fileId")
		fmt.Println(fileId)

		response, err := contribution_io.ReadAudioContributionFile(fileId)
		if err != nil {
			fmt.Println(err, " error reading audio")
		}
		fmt.Println(response, " << audio")
	}
}

type ContributionFileAgregration struct {
	ContributionFile contribution.ContributionFileHelper
	Contribution     contribution.ContributionHelper
	FileType         string
	Audio            []byte
}

func getContributionWithFile() []ContributionFileAgregration {
	var contributionFileAgregration []ContributionFileAgregration
	var myAudio []byte
	contributionFiles, err := contribution_io.ReadContributionFiles()
	if err != nil {
		fmt.Println(err, " error reading all the contributionFiles")
	} else {
		for _, contributionFile := range contributionFiles {
			contributions, err := contribution_io.ReadContribution(contributionFile.ContributionId)
			if err != nil {
				fmt.Println(err, " error reading all the contribution")
			} else {
				if contributionFile.FileType != "jpg" {
					myAudio, err = contribution_io.ReadAudioContributionFile(contributionFile.Id)
					if err != nil {
						fmt.Println(err, " not an Audio")
					}
				}

				filetype, err := contribution_io.ReadContributionFileType(contributionFile.FileType)
				if err != nil {
					fmt.Println(err, " error reading all the contribution")
				} else {
					contributionFileAgregration = append(contributionFileAgregration, ContributionFileAgregration{contribution.ContributionFileHelper{contributionFile.Id, contributionFile.ContributionId, misc.ConvertingToString(contributionFile.File), contributionFile.FileType, contributionFile.Description},
						contribution.ContributionHelper{contributions.Id, contributions.Email, contributions.Name, misc.FormatDateTime(contributions.Date), contributions.PhoneNumber, misc.ConvertingToString(contributions.Description)},
						filetype.FileType, myAudio})
				}
			}
		}
	}
	return contributionFileAgregration
}
func homeContributionHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		type PageData struct {
			Contributions []ContributionFileAgregration
			SidebarData   misc.SidebarData
			Backend_error string
			Unknown_error string
		}

		data := PageData{getContributionWithFile(), misc.GetSideBarData("contribution", ""), backend_error, unknown_error}
		files := []string{
			app.Path + "admin/contribution/contribution.html",
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

func DeleteFileTypeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		fileTypeId := chi.URLParam(r, "fileTypeId")

		if fileTypeId != "" {
			_, err := contribution_io.DeleteContributionFileType(fileTypeId)
			if err != nil {
				fmt.Println(err, " error deleting file type")
			}
		}
		http.Redirect(w, r, "/admin_user/contribution", 301)
		return
	}
}

//This is metho is called on the front page of the client side.
func CreateContribution(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		var isExtension bool
		var fileType string
		file, m, err := r.FormFile("file")
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		fileTypeId := r.PostFormValue("filetypeId")
		cellphone := r.PostFormValue("cellphone")
		message := r.PostFormValue("message")

		if err != nil {
			fmt.Println(err, "<<<error reading file>>>>This error may happen if there is no picture selected>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
			//Check the file extension
			isExtension, fileType = misc.GetFileExtension(m)

			fmt.Println("result from assement: ", isExtension)

			if isExtension == false {
				fmt.Println("error creating contribution", content)
				fmt.Println("wrong file: ", m.Filename)
				http.Redirect(w, r, "/", 301)
			}
		}
		if email != "" {
			contributionObject := contribution.Contribution{"", email, name, time.Now(), cellphone, misc.ConvertToByteArray(message)}
			contributionResult, err := contribution_io.CreateContribution(contributionObject)
			if err != nil {
				fmt.Println(err, " error creating contribution")
			} else if fileTypeId != "" {
				fileCOntribution := contribution.ContributionFile{"", contributionResult.Id, content, fileTypeId, fileType}
				_, err := contribution_io.CreateContributionFile(fileCOntribution)
				if err != nil {
					fmt.Println(err, " error creating contributionFile")
				}
			} else {
				_, err := contribution_io.DeleteContribution(contributionResult.Id)
				if err != nil {
					fmt.Println(err, " error deleting contribution")
				}
			}
		}
		http.Redirect(w, r, "/", 301)
	}
}

func CreateFileType(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}

		r.ParseForm()
		fileType := r.PostFormValue("fileType")

		if fileType != "" {
			contributionFileType := contribution.ContributionFileType{"", fileType}
			_, err := contribution_io.CreateContributionFileType(contributionFileType)
			if err != nil {
				fmt.Println("error creating file type")
				http.Redirect(w, r, "/admin_user/contribution/", 301)
			}
		}
		http.Redirect(w, r, "/admin_user/contribution/", 301)
	}
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var success_notice string
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			success_notice = app.Session.GetString(r.Context(), "creation-successful")
			app.Session.Remove(r.Context(), "creation-successful")
		}
		contributions, err := contribution_io.ReadContributions()
		if err != nil {
			fmt.Println("error reading contributions")
		}
		contributionFileTypes, err := contribution_io.ReadContributionFileTypes()
		if err != nil {
			fmt.Println("error reading contributions types")
		}
		placeCategories, err := place_io.ReadPlaceCategories()
		if err != nil {
			fmt.Println("error reading placeCategories")
		}

		type PageData struct {
			Success_notice    string
			SidebarData       misc.SidebarData
			Contribution      []contribution.Contribution
			ContributionTypes []contribution.ContributionFileType
			PlaceCategories   []place.PlaceCategory
		}
		data := PageData{success_notice, misc.GetSideBarData("setting", ""), contributions, contributionFileTypes, placeCategories}
		files := []string{
			app.Path + "admin/settings/admin-settings.html",
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
