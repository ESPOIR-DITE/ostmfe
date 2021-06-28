package group

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/constates"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	"ostmfe/domain/pages"
	partner2 "ostmfe/domain/partner"
	project2 "ostmfe/domain/project"
	"ostmfe/io/group_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/pages/admin"
	"ostmfe/io/partner_io"
	"ostmfe/io/project_io"
	"ostmfe/utile"
)

func GroupHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", GroupsHandler(app))
	r.Post("/create", CreateGroupsHandler(app))
	//r.Post("/create_history", CreateHistory_ImageHandler(app))
	r.Get("/picture/{groupId}", GroupPictureHandler(app))
	r.Get("/edit/{groupId}", GroupEditHandler(app))

	r.Post("/create_pictures", CreateImageHandler(app))
	r.Post("/create_history", CreateHistoryHandler(app))
	r.Post("/update_pictures", UpdateImageHandler(app))
	r.Post("/update_history", UpdateHistoryHandler(app))
	r.Post("/update_details", UpdateDetailsHandler(app))

	r.Get("/activate_comment/{commentId}/{groupId}", ActivateCommentHandler(app))
	r.Get("/delete-group/{groupId}", DeleteGroup(app))
	r.Post("/add_group_pictures/{groupId}", AddGroupGallery(app))

	//Gallery
	r.Post("/create-gallery", CreateGroupGalleryHandler(app))
	r.Get("/delete-gallery/{pictureId}/{groupId}/{groupGalleryId}", DeleteGalleryHandler(app))

	r.Post("/create-descriptive-Image", AddDescriptiveImage(app))

	return r
}

func AddDescriptiveImage(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		groupId := r.PostFormValue("groupId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if groupId != "" && description != "" {
			ImageObject := image.Images{"", content, description}
			imageObj, err := image_io.CreateImage(ImageObject)
			if err != nil {
				fmt.Println(err, " error creating image")
			} else {
				ImageType, err := image_io.ReadImageTypeWithName(constates.DESCRIPTIVE)
				if err != nil {
					fmt.Println(err, " error Reading image Type")
				}
				groupImage := group.GroupImage{"", imageObj.Id, groupId, ImageType.Id, description}
				_, errx := group_io.CreateGroupImage(groupImage)
				if errx != nil {
					fmt.Println(err, " error creating GroupImage")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return
	}
}

func AddGroupGallery(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		groupId := r.PostFormValue("groupId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if groupId != "" && description != "" {
			galery := image.Gallery{"", content, description}
			galleryObject, err := image_io.CreateGalery(galery)
			if err != nil {
				fmt.Println(err, " error creating gallery")
			} else {
				groupGallery := group.GroupGalery{"", groupId, galleryObject.Id}
				_, err := group_io.CreateGroupGalery(groupGallery)
				if err != nil {
					fmt.Println(err, " error creating GroupGallery")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/group/", 301)
		return
	}
}

func DeleteGroup(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		groupId := chi.URLParam(r, "groupId")
		if groupId != "" {
			_, err := group_io.DeleteGroup(groupId)
			if err != nil {
				app.ErrorLog.Println(err, " error deleting group")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
		} else {
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
			return
		}
		http.Redirect(w, r, "/admin_user/group", 301)
		return
	}
}

func ActivateCommentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentId := chi.URLParam(r, "commentId")
		groupId := chi.URLParam(r, "groupId")
		result := misc.ActivateComment(commentId)
		fmt.Print("Activation Result: ", result)
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return
	}
}

func DeleteGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		pictureId := chi.URLParam(r, "pictureId")
		groupId := chi.URLParam(r, "groupId")
		groupGalleryId := chi.URLParam(r, "groupGalleryId")

		//Deleting project
		gallery, err := image_io.DeleteGalery(pictureId)
		if err != nil {
			fmt.Println("error deleting gallery")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
			return
		} else {
			_, err := group_io.DeleteGroupGalery(groupGalleryId)
			if err != nil {
				fmt.Println("error deleting group gallery")
				fmt.Println("ROLLING BACK!!!")
				_, err := image_io.UpdateGallery(gallery)
				if err != nil {
					fmt.Println("error updating gallery")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
		}
		fmt.Println(" successful deletion.")
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted: group Gallery. ")
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return
	}
}

func CreateGroupGalleryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		var content []byte
		r.ParseForm()
		file, _, err := r.FormFile("file")
		groupId := r.PostFormValue("groupId")
		description := r.PostFormValue("description")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading contribution file>>>>This error should happen>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if groupId != "" && description != "" {
			galery := image.Gallery{"", content, description}
			galleryObject, err := image_io.CreateGalery(galery)
			if err != nil {
				fmt.Println(err, " error creating gallery")
			} else {
				groupGalery := group.GroupGalery{"", groupId, galleryObject.Id}
				_, err := group_io.CreateGroupGalery(groupGalery)
				if err != nil {
					fmt.Println(err, " error creating GroupGallery")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
					return
				}
				if app.Session.GetString(r.Context(), "creation-successful") != "" {
					app.Session.Remove(r.Context(), "creation-successful")
				}
				app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted an event Group")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return
	}
}

func UpdateDetailsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		groupId := r.PostFormValue("groupId")
		groupName := r.PostFormValue("groupName")
		Description := r.PostFormValue("Description")
		historyId := r.PostFormValue("historyId")

		if Description != "" && groupId != "" && groupName != "" {
			groupObejct := group.Groupes{groupId, groupName, Description, historyId}
			group, err := group_io.UpdateGroup(groupObejct)
			if err != nil {
				fmt.Println(err, " something went wrong! could not create group")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated Details of the following group: "+group.Name)
			http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
			return
		}
		fmt.Println("one or more fields missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return

	}
}

func UpdateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		groupId := r.PostFormValue("groupId")
		historyId := r.PostFormValue("historyId")
		//checking if there is contents in the variables
		if historyContent != "" && groupId != "" && historyId != "" {
			history := history2.Histories{historyId, misc.ConvertToByteArray(historyContent)}

			_, err := history_io.UpdateHistorie(history)
			if err != nil {
				fmt.Println(err, " something went wrong! could not create history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated Group History")
			http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
			return
		}
		fmt.Println("one or more fields missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return
	}
}

func CreateHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		historyContent := r.PostFormValue("myArea")
		groupId := r.PostFormValue("groupId")
		//checking if there is contents in the variables
		if historyContent != "" && groupId != "" {
			history := history2.Histories{"", misc.ConvertToByteArray(historyContent)}

			newHistory, err := history_io.CreateHistorie(history)
			if err != nil {
				fmt.Println(err, " something went wrong! could not create history")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
			fmt.Println("HistoryId created successfully ..")
			fmt.Println(" proceeding into creation of a group_history.....")
			groupHistory := group.GroupHistory{"", groupId, newHistory.Id}
			_, errr := group_io.CreateGroupHistory(groupHistory)
			if errr != nil {
				fmt.Println(err, " could not create GroupHistory")
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
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			fmt.Println(" successfully created")
			http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
			return
		}
		fmt.Println("one or more fields missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return
	}
}

func UpdateImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		var content []byte
		file, _, err := r.FormFile("file")
		groupId := r.PostFormValue("groupId")
		imageId := r.PostFormValue("imageId")
		if err != nil {
			fmt.Println(err, "<<<error reading file>>>>This error may happen if there is no picture selected>>>")
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		if groupId != "" {
			imageObject := image.Images{imageId, content, ""}
			_, err := image_io.CreateImage(imageObject)
			if err != nil {
				fmt.Println(err, " error creating GroupImage")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
				return
			}
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updated an image for a Group")
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return
	}
}

func CreateImageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		r.ParseForm()
		var content []byte
		file, _, err := r.FormFile("file")
		groupId := r.PostFormValue("groupId")
		imageType := adminHelper.GetProfileImageId()
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
			http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
			return
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		myGroup, err := group_io.ReadGroup(groupId)
		if err != nil {
			fmt.Println(err, " could not read group")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/group", 301)
			return
		}

		imageObject, err := image_io.CreateImage(image.Images{"", content, groupId})
		if err != nil {
			fmt.Println(err, " error creating Image")
			http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		}
		groupImage := group.GroupImage{"", imageObject.Id, groupId, imageType, ""}
		_, errr := group_io.CreateGroupImage(groupImage)
		if errr != nil {
			fmt.Println(errr, " error creating GroupImage")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
			return
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully created image(s) for the following Group  : "+myGroup.Name)
		http.Redirect(w, r, "/admin_user/group/edit/"+groupId, 301)
		return
	}
}

func GroupEditHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		groupId := chi.URLParam(r, "groupId")

		commentNumber, pendingcomments, activeComments := groupCommentCalculation(groupId)
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}

		groupAdinData, err := admin.GetGroupAdminEditPageData(groupId)
		members, number, err := GetMembers(groupId)
		if err != nil {
			fmt.Println(err, " error reading group Member")
		}

		type PageData struct {
			Members         []MemberData
			GroupNumber     int64
			PageData        pages.GroupAdminEditPresentation
			Groups          GroupData
			SidebarData     misc.SidebarData
			Gallery         []misc.GroupGalleryImages
			Comments        []comment.CommentHelper2
			CommentNumber   int64
			PendingComments int64
			ActiveComments  int64
			AdminName       string
			AdminImage      string
		}

		data := PageData{members, number, groupAdinData,
			GetGroupData(groupId),
			misc.GetSideBarData("group", ""),
			misc.GetGroupGallery(groupId),
			GetGroupCommentsWithEventId(groupId),
			commentNumber,
			pendingcomments,
			activeComments, adminName, adminImage}

		files := []string{
			app.Path + "admin/group/edit_group.html",
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
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PageData struct {
			Projects      []project2.Project
			Partners      []partner2.Partner
			Group         group.Groupes
			Backend_error string
			Unknown_error string
			AdminName     string
			AdminImage    string
		}
		data := PageData{projects, partners, groupObject, backend_error, unknown_error, adminName, adminImage}
		files := []string{
			app.Path + "admin/group/group_image.html",
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
		var newGroup group.Groupes
		var content []byte
		file, _, err := r.FormFile("file")

		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should not happen>>>")
			http.Redirect(w, r, "/admin_user/group", 301)
			return
		} else {
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		groupName := r.PostFormValue("name")
		project := r.PostFormValue("project")
		description := r.PostFormValue("description")
		partner := r.PostFormValue("partner")
		historyContent := r.PostFormValue("mytextarea")

		if groupName != "" && description != "" {

			var historyReturnObject history2.Histories
			//Creating group history.
			if historyContent != "" {
				history := history2.Histories{"", misc.ConvertToByteArray(historyContent)}

				historyReturnObject, err = history_io.CreateHistorie(history)
				if err != nil {
					fmt.Println(err, " error when creating history")
				} else {
					_, err := group_io.CreateGroupHistory(group.GroupHistory{"", newGroup.Id, historyReturnObject.Id})
					if err != nil {
						fmt.Println(err, " error when creating group history")
					}
				}
			}

			groupObject := group.Groupes{"", groupName, description, historyReturnObject.Id}
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

			//Group Image
			imageObject := image.Images{"", content, "profile"}
			imageReturnObject, err := image_io.CreateImage(imageObject)
			if err != nil {
				fmt.Println(err, " error when creating image")
			} else {
				imageType, err := image_io.ReadImageTypeWithName(utile.PROFILE)
				if err != nil {
					fmt.Println(err, " error reading imageType")
				}
				groupImage := group.GroupImage{"", imageReturnObject.Id, newGroup.Id, imageType.Name, groupName}
				_, err = group_io.CreateGroupImage(groupImage)
				if err != nil {
					fmt.Println(err, " error when creating group image")
				}
			}

			//Group Partner
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
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Groupes : "+groupName)
			http.Redirect(w, r, "/admin_user/group", 301)
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
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
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
			Groups        []group.Groupes
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}
		data := PageData{projects, partners,
			backend_error,
			unknown_error, groups,
			misc.GetSideBarData("group", ""), adminName, adminImage}
		files := []string{
			app.Path + "admin/group/groups.html",
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
