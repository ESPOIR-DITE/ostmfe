package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const projectpartnerURL = api.BASE_URL + "place_partner/"

func CreateProjectPartner(prjP project.ProjectPartner) (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().SetBody(prjP).Post(projectpartnerURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func updateProjectPartner(prjP project.ProjectPartner) (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().SetBody(prjP).Post(projectpartnerURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProjectPartner(id string) (project.ProjectPartner, error) {
	entity := project.ProjectPartner{}
	resp, _ := api.Rest().Get(projectpartnerURL + "read?id" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllOfProjectPartner(id string) ([]project.ProjectPartner, error) {
	entity := []project.ProjectPartner{}
	resp, _ := api.Rest().Get(projectpartnerURL + "readAllOf?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteProjectPartner(id string) (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().Get(projectpartnerURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadProjectPartners() (project.ProjectPartner, error) {

	entity := project.ProjectPartner{}
	resp, _ := api.Rest().Get(projectpartnerURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
