package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const projectpageflowURL = api.BASE_URL + "project_page_low/"

func CreateProjectPageFLow(hist project.ProjectPageFlow) (project.ProjectPageFlow, error) {
	entity := project.ProjectPageFlow{}
	resp, _ := api.Rest().SetBody(hist).Post(projectpageflowURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateProjectPageFLow(hist project.ProjectPageFlow) (project.ProjectPageFlow, error) {
	entity := project.ProjectPageFlow{}
	resp, _ := api.Rest().SetBody(hist).Post(projectpageflowURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProjectPageFLow(id string) (project.ProjectPageFlow, error) {
	entity := project.ProjectPageFlow{}
	resp, _ := api.Rest().Get(projectpageflowURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadProjectPageFLowWithHistoryId(id string) (project.ProjectPageFlow, error) {
	entity := project.ProjectPageFlow{}
	resp, _ := api.Rest().Get(projectpageflowURL + "readWithHistoryId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectPageFLowsWithProjectId(id string) ([]project.ProjectPageFlow, error) {
	entity := []project.ProjectPageFlow{}
	resp, _ := api.Rest().Get(projectpageflowURL + "readAllWithHistoryId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteProjectPageFLow(id string) (project.ProjectPageFlow, error) {
	entity := project.ProjectPageFlow{}
	resp, _ := api.Rest().Get(projectpageflowURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProjectPageFLows() (project.ProjectPageFlow, error) {
	entity := project.ProjectPageFlow{}
	resp, _ := api.Rest().Get(projectpageflowURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
