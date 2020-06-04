package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const projectmemberURL = api.BASE_URL + "project_member/"

func CreateProjectMember(pm project.ProjectMember) (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().SetBody(pm).Post(projectmemberURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateProjectMember(pm project.ProjectMember) (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().SetBody(pm).Post(projectmemberURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectMember(id string) (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().Get(projectmemberURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteProjectMember(id string) (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().Get(projectmemberURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectMembers() (project.ProjectMember, error) {

	entity := project.ProjectMember{}
	resp, _ := api.Rest().Get(projectmemberURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllOfProjectMembers(id string) ([]project.ProjectMember, error) {
	entity := []project.ProjectMember{}
	resp, _ := api.Rest().Get(projectmemberURL + "readAllOf?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
