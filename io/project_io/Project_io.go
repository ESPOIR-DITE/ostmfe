package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const projectURL = api.BASE_URL + "project/"

func CreateProject(P project.Project) (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().SetBody(P).Post(projectURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func updateProject(P project.Project) (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().SetBody(P).Post(projectURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProject(id string) (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().Get(projectURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteProject(id string) (project.Project, error) {

	entity := project.Project{}
	resp, _ := api.Rest().Get(projectURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProjects() ([]project.Project, error) {

	entity := []project.Project{}
	resp, _ := api.Rest().Get(projectURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
