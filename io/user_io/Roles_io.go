package user_io

import (
	"errors"
	"ostmfe/api"
	user2 "ostmfe/domain/user"
)

const roleURL = api.BASE_URL + "role/"

func CreateRole(roles user2.Roles) (user2.Roles, error) {
	var entity user2.Roles
	resp, _ := api.Rest().SetBody(roles).Post(roleURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateRole(roles user2.Roles) (user2.Roles, error) {
	var entity user2.Roles
	resp, _ := api.Rest().SetBody(roles).Post(roleURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadRole(id string) (user2.Roles, error) {
	var entity user2.Roles
	resp, _ := api.Rest().Get(roleURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteRole(id string) (user2.Roles, error) {
	var entity user2.Roles
	resp, _ := api.Rest().Get(roleURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadRoles() ([]user2.Roles, error) {
	var entity []user2.Roles
	resp, _ := api.Rest().Get(roleURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
