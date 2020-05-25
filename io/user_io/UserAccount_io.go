package user_io

import (
	"errors"
	"ostmfe/api"
	user2 "ostmfe/domain/user"
)

const useraccountURL = api.BASE_URL + "user_account"

func CreateUserAccount(account user2.UserAccount) (user2.UserAccount, error) {
	var entity user2.UserAccount
	resp, _ := api.Rest().SetBody(account).Post(useraccountURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserAccount(account user2.UserAccount) (user2.UserAccount, error) {
	var entity user2.UserAccount
	resp, _ := api.Rest().SetBody(account).Post(useraccountURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserAccount(id string) (user2.UserAccount, error) {
	var entity user2.UserAccount
	resp, _ := api.Rest().Get(useraccountURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUserAccount(id string) (user2.UserAccount, error) {
	var entity user2.UserAccount
	resp, _ := api.Rest().Get(useraccountURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserAccounts() ([]user2.UserAccount, error) {
	var entity []user2.UserAccount
	resp, _ := api.Rest().Get(useraccountURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
