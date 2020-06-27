package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventprojectURL = api.BASE_URL + "event_project/"

func CreateEventProject(prj event.EventProject) (event.EventProject, error) {
	entity := event.EventProject{}
	resp, _ := api.Rest().SetBody(prj).Post(eventprojectURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventProject(prj event.EventProject) (event.EventProject, error) {
	entity := event.EventProject{}
	resp, _ := api.Rest().SetBody(prj).Post(eventprojectURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventProject(id string) (event.EventProject, error) {
	entity := event.EventProject{}
	resp, _ := api.Rest().Get(eventprojectURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventProjectOf(id string) ([]event.EventProject, error) {
	entity := []event.EventProject{}
	resp, _ := api.Rest().Get(eventprojectURL + "readOf?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteEventProject(id string) (event.EventProject, error) {
	entity := event.EventProject{}
	resp, _ := api.Rest().Get(eventprojectURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventProjects() ([]event.EventProject, error) {
	entity := []event.EventProject{}
	resp, _ := api.Rest().Get(eventprojectURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
