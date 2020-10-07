package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventgroupURL = api.BASE_URL + "event_group/"

func CreateEventGroup(E event.EventGroup) (event.EventGroup, error) {

	entity := event.EventGroup{}
	resp, _ := api.Rest().SetBody(E).Post(eventgroupURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventGroup(E event.EventGroup) (event.EventGroup, error) {

	entity := event.EventGroup{}
	resp, _ := api.Rest().SetBody(E).Post(eventgroupURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventGroup(id string) (event.EventGroup, error) {

	entity := event.EventGroup{}
	resp, _ := api.Rest().Get(eventgroupURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventGroupWithBoth(eventId, groupId string) (event.EventGroup, error) {

	entity := event.EventGroup{}
	resp, _ := api.Rest().Get(eventgroupURL + "readWithBoth?eventId=" + eventId + "&groupId=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventGroupOf(eventId string) (event.EventGroup, error) {
	entity := event.EventGroup{}
	resp, _ := api.Rest().Get(eventgroupURL + "readOf?id=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteEventGroup(id string) (event.EventGroup, error) {

	entity := event.EventGroup{}
	resp, _ := api.Rest().Get(eventgroupURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventGroupAllOfs(eventId string) ([]event.EventGroup, error) {
	entity := []event.EventGroup{}
	resp, _ := api.Rest().Get(eventgroupURL + "readWithEventId?id=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventGroups() ([]event.EventGroup, error) {
	entity := []event.EventGroup{}
	resp, _ := api.Rest().Get(eventgroupURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventGroupWithGroupId(groupId string) ([]event.EventGroup, error) {
	entity := []event.EventGroup{}
	resp, _ := api.Rest().Get(eventgroupURL + "readWithGroupId?id=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
