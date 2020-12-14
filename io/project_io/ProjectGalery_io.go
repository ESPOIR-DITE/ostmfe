package project_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/project"
)

const projectgalleryURL = api.BASE_URL + "project-galery/"

func CreateProjectGallery(P project.ProjectGallery) (project.ProjectGallery, error) {
	entity := project.ProjectGallery{}
	resp, _ := api.Rest().SetBody(P).Post(projectgalleryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateProjectGallery(P project.ProjectGallery) (project.ProjectGallery, error) {
	entity := project.ProjectGallery{}
	resp, _ := api.Rest().SetBody(P).Post(projectgalleryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadProjectGallery(id string) (project.ProjectGallery, error) {
	entity := project.ProjectGallery{}
	resp, _ := api.Rest().Get(projectgalleryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByProjectIdGallery(id string) ([]project.ProjectGallery, error) {
	entity := []project.ProjectGallery{}
	resp, _ := api.Rest().Get(projectgalleryURL + "readAllByProjectId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadProjectGalleryWithProjectId(id string) (project.ProjectGallery, error) {
	entity := project.ProjectGallery{}
	resp, _ := api.Rest().Get(projectgalleryURL + "readByProjectId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CountProjectGallery(id string) (int64, error) {
	var entity int64
	resp, _ := api.Rest().Get(projectgalleryURL + "countAllByProjectId?projectId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteProjectGallery(id string) (project.ProjectGallery, error) {
	entity := project.ProjectGallery{}
	resp, _ := api.Rest().Get(projectgalleryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadProjectGallerys() ([]project.ProjectGallery, error) {
	entity := []project.ProjectGallery{}
	resp, _ := api.Rest().Get(projectgalleryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllProjectGalleryWithProjectId(projectId string) ([]project.ProjectGallery, error) {
	entity := []project.ProjectGallery{}
	resp, _ := api.Rest().Get(projectgalleryURL + "readAllByProjectId?id=" + projectId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
