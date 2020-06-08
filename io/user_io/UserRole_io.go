package user_io

import (
	"errors"
	"ostmfe/api"
	user2 "ostmfe/domain/user"
)

const userroleURL = api.BASE_URL + "user_role/"

func CreateUserRole(role user2.UserRole) (user2.UserRole, error) {
	var entity user2.UserRole
	resp, _ := api.Rest().SetBody(role).Post(userroleURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserRole(role user2.UserRole) (user2.UserRole, error) {
	var entity user2.UserRole
	resp, _ := api.Rest().SetBody(role).Post(userroleURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserRole(id string) (user2.UserRole, error) {
	var entity user2.UserRole
	resp, _ := api.Rest().Get(userroleURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserRoleWithEmail(id string) (user2.UserRole, error) {
	var entity user2.UserRole
	resp, _ := api.Rest().Get(userroleURL + "readWithemail?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUserRole(id string) (user2.UserRole, error) {
	var entity user2.UserRole
	resp, _ := api.Rest().Get(userroleURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserRoles(id string) (user2.UserRole, error) {
	var entity user2.UserRole
	resp, _ := api.Rest().Get(userroleURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
