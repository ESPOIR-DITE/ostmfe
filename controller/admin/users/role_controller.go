package users

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	user2 "ostmfe/domain/user"
	"ostmfe/io/user_io"
)

func RoleController(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", RoleHandler(app))
	r.Post("/create", CreateRoleHandler(app))
	r.Get("/edit/{roleId}", EditRoleHandler(app))
	r.Post("/update", UpdateRoleHandler(app))
	return r
}
func EditRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleId := chi.URLParam(r, "roleId")
		//var role user2.Roles
		role, err := user_io.ReadRole(roleId)
		if err != nil {
			fmt.Println(err, "error reading role")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/users", 301)
			return
		}

		roles, err := user_io.ReadRoles()
		if err != nil {
			fmt.Println(err, " error reading roles")
		}
		type PageData struct {
			Role        user2.Roles
			Roles       []user2.Roles
			SidebarData misc.SidebarData
		}
		data := PageData{role, roles, misc.GetSideBarData("user", "role")}
		files := []string{
			app.Path + "admin/user/role_edit.html",
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
func CreateRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		role := r.PostFormValue("role")
		description := r.PostFormValue("description")
		if role != "" && description != "" {
			roleObject := user2.Roles{"", role, description}
			roleResult, err := user_io.CreateRole(roleObject)
			if err != nil {
				fmt.Println(err, "error creating new Role")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/role", 301)
				return
			}
			fmt.Println(err, "Creation of a new user successful")
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new Role : "+roleResult.Role)
			http.Redirect(w, r, "/admin_user/role", 301)
			return

		}
	}
}
func RoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var unknown_error string
		var backend_error string
		var Error string
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			unknown_error = app.Session.GetString(r.Context(), "creation-unknown-error")
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			backend_error = app.Session.GetString(r.Context(), "user-create-error")
			app.Session.Remove(r.Context(), "user-create-error")
		}
		roles, err := user_io.ReadRoles()
		if err != nil {
			Error = "Internal error,Please try again later"
			fmt.Println("error reading roles")
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			Error         string
			RoleList      []user2.Roles
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}
		data := PagePage{backend_error, unknown_error, Error, roles,
			misc.GetSideBarData("user", "role"),
			adminName, adminImage,
		}
		files := []string{
			app.Path + "admin/user/roles.html",
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
func UpdateRoleHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		roleId := r.PostFormValue("roleId")
		Role := r.PostFormValue("role")
		Description := r.PostFormValue("description")

		//We are proceeding into updating the role.
		if roleId != "" && Role != "" && Description != "" {
			newRoleObejct := user2.Roles{roleId, Role, Description}
			_, err := user_io.UpdateRole(newRoleObejct)
			if err != nil {
				fmt.Println(err, " could not updating User Line: 124")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/role/edit/"+roleId, 301)
					return
				}
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following Role : "+Role)
			http.Redirect(w, r, "/admin_user/role", 301)
			return
		}
		fmt.Println("one of the field is empty", roleId, "<<<<RoleId", Role, "<<<<role", Description, "<<<<<<description")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		}
		http.Redirect(w, r, "/admin_user/role/edit/"+roleId, 301)
		return
	}
}
