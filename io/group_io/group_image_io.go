package group_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/group"
)

const groupimageURL = api.BASE_URL + "group_image/"

func CreateGroupImage(myEvent group.GroupImage) (group.GroupImage, error) {
	entity := group.GroupImage{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupimageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGroupImage(myEvent group.GroupImageHelper) (group.GroupImage, error) {
	entity := group.GroupImage{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupimageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGroupImage(id string) (group.GroupImage, error) {
	entity := group.GroupImage{}
	resp, _ := api.Rest().Get(groupimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupImageWithGroupId(groupId string) (group.GroupImage, error) {
	entity := group.GroupImage{}
	resp, _ := api.Rest().Get(groupimageURL + "readWith?id=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupImagesWithGroupId(groupId string) ([]group.GroupImage, error) {
	entity := []group.GroupImage{}
	resp, _ := api.Rest().Get(groupimageURL + "readAllWith?id=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteGroupImage(id string) (group.GroupImage, error) {
	entity := group.GroupImage{}
	resp, _ := api.Rest().Get(groupimageURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupImages() ([]group.GroupImage, error) {
	entity := []group.GroupImage{}
	resp, _ := api.Rest().Get(groupimageURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
