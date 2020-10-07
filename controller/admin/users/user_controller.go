package users

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	image2 "ostmfe/domain/image"
	user2 "ostmfe/domain/user"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/user_io"
	"time"
)

func UserController(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", UserHandler(app))
	r.Get("/new", NewUserHandler(app))
	r.Get("/edit/{userId}", EditUserHandler(app))
	r.Get("/delete/{userId}", DeleteUserHandler(app))
	r.Post("/create", CreateUserHandler(app))
	r.Post("/update_user", UpdateUserHandler(app))
	r.Post("/update_history", UpdateUserHistoryHandler(app))
	return r
}

func UpdateUserHistoryHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		myArea := r.PostFormValue("myArea")
		historyId := r.PostFormValue("historyId")
		email := r.PostFormValue("email")

		if historyId != "" && myArea != "" && email != "" {
			history := history2.Histories{historyId, misc.ConvertToByteArray(myArea)}
			_, err := history_io.UpdateHistorie(history)
			if err != nil {
				fmt.Println(err, " error update history")
			}
		}

		fmt.Println("Creation of a new user successful")
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully update")
		http.Redirect(w, r, "/admin_user/users/edit/"+email, 301)
		return
	}
}

func DeleteUserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")

		_, err := user_io.ReadUser(userId)
		if err != nil {
			fmt.Println(err, " error reading user")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/users", 301)
			return
		}
		//Deleting User Image
		_, errx := user_io.DeleteUser(userId)
		if errx != nil {
			fmt.Println(errx, " error deleting user")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/users", 301)
			return
		}

		//If the deletion of user went well
		userRole, err := user_io.ReadUserRoleWithEmail(userId)
		if err != nil {
			fmt.Println(errx, " error reading user role")
		} else {
			_, err := user_io.DeleteUserRole(userRole.Id)
			if err != nil {
				fmt.Println(errx, " error deleting user role")
			}
		}

		//Deleting userAccount
		_, errorDeleteUserAccount := user_io.DeleteUserAccount(userId)
		if errorDeleteUserAccount != nil {
			fmt.Println(errorDeleteUserAccount, " error deleting userAccount")
		}

		//Deleting userImage
		UserImage, err := user_io.ReadUserImageWithEmail(userId)
		if err != nil {
			fmt.Println(errx, " error reading user Image")
		} else {
			_, err := user_io.DeleteUserImage(UserImage.Id)
			if err != nil {
				fmt.Println(errx, " error deleting userImage")
			}
		}
		fmt.Println("Creation of a new user successful")
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully deleted user")
		http.Redirect(w, r, "/admin_user/users", 301)
		return
	}
}
func EditUserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var role user2.Roles
		var unknown_error string
		var backend_error string
		var historie history2.HistoriesHelper
		var imagehelper image2.ImagesHelper
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			unknown_error = app.Session.GetString(r.Context(), "creation-unknown-error")
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			backend_error = app.Session.GetString(r.Context(), "user-create-error")
			app.Session.Remove(r.Context(), "user-create-error")
		}

		userId := chi.URLParam(r, "userId")

		//Reading the user
		user, err := user_io.ReadUser(userId)
		if err != nil {
			fmt.Println(err, "error reading user")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/users", 301)
			return
		}
		userAccount, err := user_io.ReadUserAccountwithEmail(userId)
		if err != nil {
			fmt.Println(err, "error reading new userAccount")
		}
		//Reading User Role with email
		userRole, err := user_io.ReadUserRoleWithEmail(userId)
		if err != nil {
			fmt.Println(err, "error reading new userRole")
		} else {
			role, err = user_io.ReadRole(userRole.RoleId)
			if err != nil {
				fmt.Println(err, "error reading new Role line: 1991")

			}
		}
		//UserImage
		userImage, err := user_io.ReadUserImageWithEmail(userId)
		if err != nil {
			fmt.Println(err, "error reading userImage")
		} else {
			imageObject, err := image_io.ReadImage(userImage.ImageId)
			if err != nil {
				fmt.Println(err, "error reading image for: ", userImage.ImageId)
			} else {
				imagehelper = image2.ImagesHelper{imageObject.Id, misc.ConvertingToString(imageObject.Image), userImage.Id}
			}
			//Reading history
			userHistory, err := history_io.ReadHistorie(userImage.HistoryId)
			if err != nil {
				fmt.Println(err, "error reading histories")
			} else {
				historie = history2.HistoriesHelper{userHistory.Id, misc.ConvertingToString(userHistory.History)}
			}
		}

		roles, err := user_io.ReadRoles()
		if err != nil {
			fmt.Println(err, " error reading roles")
		}

		type PageData struct {
			Backend_error string
			Unknown_error string
			User          user2.Users
			Role          user2.Roles
			Roles         []user2.Roles
			UserAccount   user2.UserAccount
			SidebarData   misc.SidebarData
			Image         image2.ImagesHelper
			History       history2.HistoriesHelper
		}
		data := PageData{backend_error, unknown_error, user, role, roles, userAccount, misc.GetSideBarData("user", "user"), imagehelper, historie}
		files := []string{
			app.Path + "admin/user/edit_user.html",
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

func UserHandler(app *config.Env) http.HandlerFunc {
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
		users := misc.GetUserAndRole()
		role, err := user_io.ReadRoles()
		if err != nil {
			fmt.Println("error: ", err)
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			Users         []misc.UsersAndRoles
			Roles         []user2.Roles
			SidebarData   misc.SidebarData
		}
		data := PagePage{backend_error, unknown_error, users, role, misc.GetSideBarData("user", "user")}
		files := []string{
			app.Path + "admin/user/users.html",
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

func NewUserHandler(app *config.Env) http.HandlerFunc {
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
			app.Path + "admin/user/new_user.html",
			//app.Path + "admin/template/navbar.html",
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

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var success_notice string
		var Error string
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			success_notice = app.Session.GetString(r.Context(), "creation-successful")
			app.Session.Remove(r.Context(), "creation-successful")
		}
		users, err := user_io.ReadUsers()
		if err != nil {
			Error = "Internal server error"
			fmt.Println(Error, " error reading Users")
		}
		type PageData struct {
			Success_notice string
			Error          string
			Users          []user2.Users
		}
		data := PageData{success_notice, Error, users}
		files := []string{
			app.Path + "admin/user/users.html",
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

/***
we are getting the form from html
we grab all the fields corresponding to the name assigned to them
we create an object with the records collected from the html
we then send the object to the backend, if an error occurs we will redirect back to new user html file to try again.
*/
func CreateUserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, _, err := r.FormFile("file")
		mytextarea := r.PostFormValue("mytextarea")
		name := r.PostFormValue("name")
		surname := r.PostFormValue("surname")
		email := r.PostFormValue("email")
		userRoleId := r.PostFormValue("userRoleId")
		password := r.PostFormValue("password")

		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should happen>>>")
		}
		reader := bufio.NewReader(file)
		content, _ := ioutil.ReadAll(reader)

		if mytextarea != "" {
			historyies := history2.Histories{"", misc.ConvertToByteArray(mytextarea)}
			newHistoryies, err := history_io.CreateHistorie(historyies)
			if err != nil {
				fmt.Println(err, " error create histories")
			} else {
				imageObejct := image2.Images{"", content, ""}
				image, err := image_io.CreateImage(imageObejct)
				if err != nil {
					fmt.Println(err, " error create image")
				} else {
					userImageObject := user2.UserImage{"", email, newHistoryies.Id, image.Id, ""}
					_, err := user_io.CreateUserImage(userImageObject)
					if err != nil {
						fmt.Println(err, " error create userImage")
					}
				}
			}
		}

		fmt.Println(name, "<<name  surname>>", surname, "  email>>", email)

		if name != "" && surname != "" && email != "" && userRoleId != "" && password != "" {
			//Creating user
			user := user2.Users{email, name, surname}
			newUser, err := user_io.CreateUser(user)
			if err != nil {
				fmt.Println(err, "error creating new user line: 57")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users", 301)
				return
			}
			//Creating User Role

			userRoleObject := user2.RoleOfUser{"", email, userRoleId}
			userRole, err := user_io.CreateUserRole(userRoleObject)
			if err != nil {
				//Here we deleting the user if userRole creation has failed.
				_, err = user_io.DeleteUser(email)
				if err != nil {
					fmt.Println(err, " error deleting user")
				}
				fmt.Println(err, "error creating new user Role")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users", 301)
				return
			}

			//Creating User Account

			userAccountObject := user2.UserAccount{email, time.Now(), password}
			_, err = user_io.CreateUserAccount(userAccountObject)
			if err != nil {
				//Here we deleting the user if userRole creation has failed.
				_, err = user_io.DeleteUser(email)
				if err != nil {
					fmt.Println(err, " error deleting user")
				}
				_, err = user_io.DeleteUserRole(userRole.Id)
				if err != nil {
					fmt.Println(err, " error deleting user role")
				}
				fmt.Println(err, "error creating new user Account")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users", 301)
				return
			}

			fmt.Println("Creation of a new user successful")
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create a new user : "+newUser.Name)
			http.Redirect(w, r, "/admin_user/users", 301)
			return

		}
		fmt.Println("One of the field is missing")
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			app.Session.Remove(r.Context(), "creation-unknown-error")
			return
		}
		app.Session.Put(r.Context(), "creation-unknown-error", "You have encountered an unknown error, please try again")
		http.Redirect(w, r, "/admin_user/users", 301)
		return
	}
}

func UpdateUserHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		var userRoleObejct user2.RoleOfUser
		var newUser user2.Users
		var content []byte
		var isfilether = false
		file, _, err := r.FormFile("file")
		roleId := r.PostFormValue("roleId")
		imageId := r.PostFormValue("imageId")
		email := r.PostFormValue("email")
		surname := r.PostFormValue("surname")
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")

		if err == nil {
			isfilether = true
			reader := bufio.NewReader(file)
			content, _ = ioutil.ReadAll(reader)
		}
		fmt.Println("email: " + email)
		user, err := user_io.ReadUser(email)
		if err != nil {
			fmt.Println(err, " could not read user Line: 113")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/users/edit/"+email, 301)
			return
		}
		if isfilether == true {
			imageObejct := image2.Images{imageId, content, ""}
			_, err := image_io.UpdateImage(imageObejct)
			if err != nil {
				fmt.Println(err, " error updating image")
			}
		}

		//We need to check if the user object has changed.
		if user.Name != name || user.Surname != surname || user.Email != email {
			fmt.Println(" Updating User")
			newUserObejct := user2.Users{email, name, surname}
			newUser, err = user_io.UpdateUser(newUserObejct)
			if err != nil {
				fmt.Println(err, " could not update User")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/users/edit/"+email, 301)
				return
			}
		}

		oldUserRole, err := user_io.ReadUserRoleWithEmail(email)
		if err != nil {
			fmt.Println(err, " user may not have a role yet or an error proceeding into creating one now")
			if roleId != "" {
				userRoleObejct = user2.RoleOfUser{"", email, roleId}
				_, err := user_io.CreateUserRole(userRoleObejct)
				if err != nil {
					fmt.Println(err, " Error creating user role ")
				} else {
					fmt.Println(" Creation of user role is completed")
				}
			}

		} else if oldUserRole.RoleId != roleId {
			fmt.Println(oldUserRole, " Updating userRole")
			userRoleObejct = user2.RoleOfUser{oldUserRole.RoleId, email, roleId}
			_, err := user_io.UpdateUserRole(userRoleObejct)
			if err != nil {
				fmt.Println(err, " Error updating user role ")
			}
		}

		//Reading the user Account
		userAccount, err := user_io.ReadUserAccountwithEmail(email)
		fmt.Println(userAccount, " <<user account ")
		if err != nil && password != "" {
			fmt.Println("Creating user account")
			userAccountObject := user2.UserAccount{email, time.Now(), password}
			_, err := user_io.CreateUserAccount(userAccountObject)
			if err != nil {
				fmt.Println(err, " Error creating userAccount ")
			}
		} else if password != "" && userAccount.Password != password {
			fmt.Println(" Updating User Account")
			userAccountObject := user2.UserAccount{email, userAccount.Date, password}
			_, err := user_io.UpdateUserAccount(userAccountObject)
			if err != nil {
				fmt.Println(err, " Error Updating userAccount ")
			}
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully updated the following User : "+newUser.Name)
		http.Redirect(w, r, "/admin_user/users/edit/"+user.Email, 301)
		return
	}
}
