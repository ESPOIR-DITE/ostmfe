package group_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/group"
)

const groupURL = api.BASE_URL + "group/"

func CreateGroup(myEvent group.Group) (group.Group, error) {
	entity := group.Group{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGroup(myEvent group.Group) (group.Group, error) {
	entity := group.Group{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGroup(id string) (group.Group, error) {
	entity := group.Group{}
	resp, _ := api.Rest().Get(groupURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteGroup(id string) (group.Group, error) {
	entity := group.Group{}
	resp, _ := api.Rest().Get(groupURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroups() ([]group.Group, error) {
	entity := []group.Group{}
	resp, _ := api.Rest().Get(groupURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
