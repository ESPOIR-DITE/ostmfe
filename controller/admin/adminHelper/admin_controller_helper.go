package adminHelper

import (
	"fmt"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/constates"
	"ostmfe/io/image_io"
	"ostmfe/io/pages"
)

/***
This method will check in the session if there is the email of the admin
if there is the request will execute fine.
if there is not the page will be directed to a login page
*/
func CheckAdminInSession(app *config.Env, r *http.Request) bool {
	email := app.Session.GetString(r.Context(), "email")
	//fmt.Println("Email int the session: ", email)
	if email != "" {
		return true
	}
	fmt.Println(r.URL.Path, 3001)
	return false
}

//returns adminName, adminImage, boolean.
func CheckAdminDataInSession(app *config.Env, r *http.Request) (string, string, bool) {
	adminName := app.Session.GetString(r.Context(), "adminName")
	adminImage := app.Session.GetString(r.Context(), "adminImage")
	if adminName != "" {
		return adminName, adminImage, true
	}
	return adminName, adminImage, false
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

func PutAdminDataInSession(app *config.Env, r *http.Request, email string) {
	helper, err := pages.GetAdminData(email)
	if err != nil {
		fmt.Println(err, "error reading AdminData")
		return
	}
	app.Session.Put(r.Context(), "adminName", helper.Users.Name)
	app.Session.Put(r.Context(), "adminImage", helper.Images.Id)
}

func GetDescriptiveImageId() (string, error) {
	var id string
	imageType, err := image_io.ReadImageTypeWithName(constates.DESCRIPTIVE)
	if err != nil {
		fmt.Println(err, "error reading ImageType")
		return id, err
	}
	return imageType.Id, nil
}
func GetProfileImageId() string {
	var id string
	imageType, err := image_io.ReadImageTypeWithName(constates.PROFILE)
	if err != nil {
		fmt.Println(err, "error reading ImageType")
		return id
	}
	return imageType.Id
}
