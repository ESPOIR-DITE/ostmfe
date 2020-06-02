package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const projecthistoryURL = api.BASE_URL + "project_history/"

func CreateProjectHistory(history project.ProjectHistory) (project.ProjectHistory, error) {
	entity := project.ProjectHistory{}
	resp, _ := api.Rest().SetBody(history).Post(projecthistoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateProjectHistory(history project.ProjectHistory) (project.ProjectHistory, error) {
	entity := project.ProjectHistory{}
	resp, _ := api.Rest().SetBody(history).Post(projecthistoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectHistory(id string) (project.ProjectHistory, error) {
	entity := project.ProjectHistory{}
	resp, _ := api.Rest().Get(projecthistoryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectHistoryOf(id string) (project.ProjectHistory, error) {
	entity := project.ProjectHistory{}
	resp, _ := api.Rest().Get(projecthistoryURL + "readofproject?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteProjectHistory(id string) (project.ProjectHistory, error) {
	entity := project.ProjectHistory{}
	resp, _ := api.Rest().Get(projecthistoryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectHistories() ([]project.ProjectHistory, error) {
	entity := []project.ProjectHistory{}
	resp, _ := api.Rest().Get(projecthistoryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
