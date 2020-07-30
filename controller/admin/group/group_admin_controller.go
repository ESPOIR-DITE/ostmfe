package group

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	partner2 "ostmfe/domain/partner"
	project2 "ostmfe/domain/project"
	"ostmfe/io/group_io"
	"ostmfe/io/history_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/project_io"
)

func EventHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", GroupsHandler(app))
	r.Post("/create", CreateGroupsHandler(app))
	r.Post("/create_history", CreateHistory_ImageHandler(app))
	r.Get("/picture/{groupId}", GroupPictureHandler(app))
	return r
}

func CreateHistory_ImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		var histories history2.Histories
		var groupHistory group.GroupHistory

		file, _, err := r.FormFile("file")
		file2, _, err := r.FormFile("file2")
		file3, _, err := r.FormFile("file3")
		file4, _, err := r.FormFile("file4")
		file5, _, err := r.FormFile("file5")
		file6, _, err := r.FormFile("file6")
		mytextarea := r.PostFormValue("mytextarea")
		groupId := r.PostFormValue("group")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		filesArray := []io.Reader{file, file2, file3, file4, file5, file6}
		filesByteArray := misc.CheckFiles(filesArray)

		//Creating EventHistory and History
		//fmt.Println("eventIed: ", groupId, " test>>>>", mytextarea)
		if groupId != "" && mytextarea != "" {

			//Creating Histories Object
			historyObject := history2.Histories{"", misc.ConvertToByteArray(mytextarea)}
			histories, err = history_io.CreateHistorie(historyObject)
			if err != nil {
				fmt.Println("could not create history and wont create group history")
				if app.Session.GetString(r.Context(), "user-read-error") != "" {
					app.Session.Remove(r.Context(), "user-read-error")
				}
				app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group", 301)
				return
			}

			//creating Event History
			groupHistoryObject := group.GroupHistory{"", histories.Id, groupId}
			groupHistory, err = group_io.CreateGroupHistory(groupHistoryObject)
			if err != nil {
				fmt.Println("could not create group history")
				_, err := history_io.DeleteHistorie(histories.Id)
				if err != nil {
					fmt.Println("error deleting history")
				}
				if app.Session.GetString(r.Context(), "user-read-error") != "" {
					app.Session.Remove(r.Context(), "user-read-error")
				}
				app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group", 301)
				return
			}

		} else {
			fmt.Println("One of the field is empty")
			if app.Session.GetString(r.Context(), "user-read-error") != "" {
				app.Session.Remove(r.Context(), "user-read-error")
			}
			app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/group", 301)
			return
		}

		//creating EventImage
		groupImageObejct := group.GroupImage{"", groupId, "", ""}
		groupImageHelper := group.GroupImageHelper{groupImageObejct, filesByteArray}
		_, errx := group_io.CreateGroupImage(groupImageHelper)
		/**
		Rolling back in case of error in the creation of the group image or image itself.
		*/
		if errx != nil {
			fmt.Println(err, " error could not create eventImage Proceeding into rol back.....")
			if histories.Id != "" {
				fmt.Println(err, " Deleting histories of this event....")
				_, err := history_io.DeleteHistorie(histories.Id)
				if err != nil {
					fmt.Println(err, " !!!!!error could not delete history")
				} else {
					fmt.Println(err, " Deleted")
				}
			}
			if groupHistory.Id != "" {
				fmt.Println(err, " Deleting Event histories of this event....")
				_, err := group_io.DeleteGroupHistory(groupHistory.Id)
				if err != nil {
					fmt.Println(err, " !!!!!error could not delete group history")
				} else {
					fmt.Println(err, " Deleted")
				}
			}
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/group/picture/"+groupId, 301)
			return
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Event : ")
		http.Redirect(w, r, "/admin_user", 301)
		return
	}
}

func GroupPictureHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupId := chi.URLParam(r, "groupId")
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
		//Checking if the eventiId passed is for an existing event
		groupObject, err := group_io.ReadGroup(groupId)
		if err != nil {
			fmt.Println(err, " error reading the group")
			if app.Session.GetString(r.Context(), "user-read-error") != "" {
				app.Session.Remove(r.Context(), "user-read-error")
			}
			app.Session.Put(r.Context(), "user-read-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/group", 301)
			return
		}
		//Reading all the Projects
		projects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading all the projects")
		}
		partners, err := partner_io.ReadPartners()
		if err != nil {
			fmt.Println(err, " error reading all the partners")
		}
		type PageData struct {
			Projects      []project2.Project
			Partners      []partner2.Partner
			Group         group.Group
			Backend_error string
			Unknown_error string
		}
		data := PageData{projects, partners, groupObject, backend_error, unknown_error}
		files := []string{
			app.Path + "admin/event/group_image.html",
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

func CreateGroupsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var newGroup group.Group
		groupName := r.PostFormValue("name")
		project := r.PostFormValue("project")
		description := r.PostFormValue("description")
		partner := r.PostFormValue("partner")

		if groupName != "" && description != "" {
			groupObject := group.Group{"", groupName, description}
			errs := errors.New("")
			newGroup, errs = group_io.CreateGroup(groupObject)
			if errs != nil {
				fmt.Println(errs, " error when creating a new group")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group", 301)
				return
			}
			if partner != "" {
				grouoPartnerOnejct := group.GroupPartener{"", partner, newGroup.Id, ""}
				_, err := group_io.CreateGroupPartner(grouoPartnerOnejct)
				if err != nil {
					fmt.Println(err, " error when creating group partner")
				}
			}

			//TODO will need to create EventProject description Field in HTML.
			if project != "" {
				eventProject := group.GroupProject{"", project, newGroup.Id, ""}
				_, err := group_io.CreateGroupProject(eventProject)
				if err != nil {
					fmt.Println(err, " error when creating group project")
				}
			}

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Group : "+groupName)
			http.Redirect(w, r, "/admin_user/group/picture/"+newGroup.Id, 301)
			return
		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/group", 301)
		return
	}
}

func GroupsHandler(app *config.Env) http.HandlerFunc {
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
		groups, err := group_io.ReadGroups()
		if err != nil {
			fmt.Println(err, " error reading groups")
		}
		//Reading all the Projects
		projects, err := project_io.ReadProjects()
		if err != nil {
			fmt.Println(err, " error reading all the projects")
		}
		partners, err := partner_io.ReadPartners()
		if err != nil {
			fmt.Println(err, " error reading all the partners")
		}
		type PageData struct {
			Projects      []project2.Project
			Partners      []partner2.Partner
			Backend_error string
			Unknown_error string
			Events        []group.Group
		}
		data := PageData{projects, partners, backend_error, unknown_error, groups}
		files := []string{
			app.Path + "admin/event/groups.html",
			app.Path + "admin/template/navbar.html",
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
