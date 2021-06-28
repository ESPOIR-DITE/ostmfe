package pages

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/user"
)

const indexURL = api.BASE_URL + "admin-index/"

func GetAdminData(email string) (user.UserImageHelper, error) {
	entity := user.UserImageHelper{}
	resp, _ := api.Rest().Get(indexURL + "admin-data?email=" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
