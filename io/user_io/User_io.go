package user_io

import (
	"errors"
	"ostmfe/api"
	user2 "ostmfe/domain/user"
)

const userURL = api.BASE_URL + "user/"

func CreateUser(userObject user2.Users) (user2.Users, error) {
	var entity user2.Users
	resp, _ := api.Rest().SetBody(userObject).Post(userURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUser(userObject user2.Users) (user2.Users, error) {
	var entity user2.Users
	resp, _ := api.Rest().SetBody(userObject).Post(userURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUser(id string) (user2.Users, error) {
	var entity user2.Users
	resp, _ := api.Rest().Get(userURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUser(id string) (user2.Users, error) {
	var entity user2.Users
	resp, _ := api.Rest().Get(userURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUsers() ([]user2.Users, error) {
	var entity []user2.Users
	resp, _ := api.Rest().Get(userURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
