package login

import (
	"ostmfe/api"
)

const roleURL = api.BASE_URL + "login/"

func AdminLogin(email, password string) bool {
	var entity bool
	resp, _ := api.Rest().Get(roleURL + "login?email=" + email + "&password=" + password)
	if resp.IsError() {
		return false
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return false
	}
	return entity
}
