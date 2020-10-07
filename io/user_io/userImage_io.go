package user_io

import (
	"errors"
	"ostmfe/api"
	user2 "ostmfe/domain/user"
)

const userimageURL = api.BASE_URL + "user_image/"

func CreateUserImage(roles user2.UserImage) (user2.UserImage, error) {
	var entity user2.UserImage
	resp, _ := api.Rest().SetBody(roles).Post(userimageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateUserImage(roles user2.UserImage) (user2.UserImage, error) {
	var entity user2.UserImage
	resp, _ := api.Rest().SetBody(roles).Post(userimageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserImage(id string) (user2.UserImage, error) {
	var entity user2.UserImage
	resp, _ := api.Rest().Get(userimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadUserImageWithEmail(email string) (user2.UserImage, error) {
	var entity user2.UserImage
	resp, _ := api.Rest().Get(userimageURL + "readWithEmail?id=" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserImage(id string) (user2.UserImage, error) {
	var entity user2.UserImage
	resp, _ := api.Rest().Get(userimageURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserImages() ([]user2.UserImage, error) {
	var entity []user2.UserImage
	resp, _ := api.Rest().Get(userimageURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
