package admin

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pages"
)

const adminPage = api.BASE_URL + "admin-pages/"

func GetHomeAdminPage(email string) (pages.AdminLandingPageData, error) {
	entity := pages.AdminLandingPageData{}
	resp, _ := api.Rest().Get(adminPage + "home?email=" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUsersPageData(email string) (pages.UserPageData, error) {
	entity := pages.UserPageData{}
	resp, _ := api.Rest().Get(adminPage + "users?email=" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserPageData(email, adminEmail string) (pages.UserPageData, error) {
	entity := pages.UserPageData{}
	resp, _ := api.Rest().Get(adminPage + "user?email=" + email + "email2=" + adminEmail)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

//Event
