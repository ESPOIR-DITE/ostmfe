package group_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/group"
)

const groupprojectURL = api.BASE_URL + "group_project/"

func CreateGroupProject(myEvent group.GroupProject) (group.GroupProject, error) {
	entity := group.GroupProject{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupprojectURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGroupProject(myEvent group.GroupProject) (group.GroupProject, error) {
	entity := group.GroupProject{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupprojectURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGroupProject(id string) (group.GroupProject, error) {
	entity := group.GroupProject{}
	resp, _ := api.Rest().Get(groupprojectURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupProjectWithGroupId(groupId string) (group.GroupHistory, error) {
	entity := group.GroupHistory{}
	resp, _ := api.Rest().Get(groupprojectURL + "readWith?id=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteGroupProject(id string) (group.GroupProject, error) {
	entity := group.GroupProject{}
	resp, _ := api.Rest().Get(groupprojectURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupProjects() ([]group.GroupProject, error) {
	entity := []group.GroupProject{}
	resp, _ := api.Rest().Get(grouphistoryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
