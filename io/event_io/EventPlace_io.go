package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventplaceURL = api.BASE_URL + "event_place/"

func CreateEventPlace(E event.EventPlace) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().SetBody(E).Post(eventplaceURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventPlace(E event.EventPlace) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().SetBody(E).Post(eventplaceURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPlace(id string) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().Get(eventplaceURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPlaceOf(eventId string) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().Get(eventplaceURL + "readOf?id=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventFindByPlaceId(placeId string) ([]event.EventPlace, error) {

	entity := []event.EventPlace{}
	resp, _ := api.Rest().Get(eventplaceURL + "findAllBy?PlaceId=" + placeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteEventPlace(id string) (event.EventPlace, error) {

	entity := event.EventPlace{}
	resp, _ := api.Rest().Get(eventplaceURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPlaces() ([]event.EventPlace, error) {

	entity := []event.EventPlace{}
	resp, _ := api.Rest().Get(eventplaceURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
