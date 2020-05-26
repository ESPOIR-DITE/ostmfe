package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const projectimageURL = api.BASE_URL + "project_Image/"

func CreateProjectImage(helper project.ProjectImageHelper) (project.ProjectImage, error) {
	entity := project.ProjectImage{}
	resp, _ := api.Rest().SetBody(helper).Post(projectimageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateProjectImage(helper project.ProjectImage) (project.ProjectImage, error) {
	entity := project.ProjectImage{}
	resp, _ := api.Rest().SetBody(helper).Post(projectimageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectImage(id string) (project.ProjectImage, error) {
	entity := project.ProjectImage{}
	resp, _ := api.Rest().Get(projectimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadProjectImages(id string) ([]project.ProjectImage, error) {
	entity := []project.ProjectImage{}
	resp, _ := api.Rest().Get(projectimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteProjectImage(id string) (project.ProjectImage, error) {
	entity := project.ProjectImage{}
	resp, _ := api.Rest().Get(projectimageURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
