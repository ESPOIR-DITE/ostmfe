package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventpeopleURL = api.BASE_URL + "event_people/"

func CreateEventPeople(E event.EventPeople) (event.EventPeople, error) {

	entity := event.EventPeople{}
	resp, _ := api.Rest().SetBody(E).Post(eventpeopleURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventPeople(E event.EventPeople) (event.EventPeople, error) {

	entity := event.EventPeople{}
	resp, _ := api.Rest().SetBody(E).Post(eventpeopleURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPeople(id string) (event.EventPeople, error) {

	entity := event.EventPeople{}
	resp, _ := api.Rest().Get(eventpeopleURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventPeopleWithBoth(eventId, peopleId string) (event.EventPeople, error) {

	entity := event.EventPeople{}
	resp, _ := api.Rest().Get(eventpeopleURL + "readWithBoth?eventId=" + eventId + "&peopleId=" + peopleId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventPeopleOf(eventId string) ([]event.EventPeople, error) {
	entity := []event.EventPeople{}
	resp, _ := api.Rest().Get(eventpeopleURL + "readAllOf?id=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPeopleWithPeopleId(peopleId string) ([]event.EventPeople, error) {
	entity := []event.EventPeople{}
	resp, _ := api.Rest().Get(eventpeopleURL + "readAllWithPeopleId?id=" + peopleId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteEventPeople(id string) (event.EventPeople, error) {
	entity := event.EventPeople{}
	resp, _ := api.Rest().Get(eventpeopleURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPeoples() ([]event.EventPeople, error) {
	entity := []event.EventPeople{}
	resp, _ := api.Rest().Get(eventpeopleURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
