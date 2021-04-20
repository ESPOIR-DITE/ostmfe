package adminHelper

import (
	"fmt"
	"net/http"
	"ostmfe/config"
	"ostmfe/io/image_io"
)

/***
This method will check in the session if there is the email of the admin
if there is the request will execute fine.
if there is not the page will be directed to a login page
*/
func CheckAdminInSession(app *config.Env, r *http.Request) bool {
	email := app.Session.GetString(r.Context(), "email")
	fmt.Println("Email int the session: ", email)
	if email != "" {
		return true
	}
	fmt.Println(r.URL.Path, 3001)
	return false
}

func CheckIfImageTypeIsProfile(imageTypeId string) bool {
	imageType, err := image_io.ReadImageType(imageTypeId)
	if err != nil {
		return false
	}
	if imageType.Name == "Profile" {
		return true
	}
	return false
}
