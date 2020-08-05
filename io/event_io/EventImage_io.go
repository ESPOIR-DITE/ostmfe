package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventimageURL = api.BASE_URL + "event_image/"

func CreateEventImg(image event.EventImageHelper) (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().SetBody(image).Post(eventimageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventImg(image event.EventImageHelper) (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().SetBody(image).Post(eventimageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventImg(id string) (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().Get(eventimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventImgOf(eventId string) ([]event.EventImage, error) {
	entity := []event.EventImage{}
	resp, _ := api.Rest().Get(eventimageURL + "readAllOf?id=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteEventImg(id string) (event.EventImage, error) {

	entity := event.EventImage{}
	resp, _ := api.Rest().Get(eventimageURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventmgs() ([]event.EventImage, error) {
	entity := []event.EventImage{}
	resp, _ := api.Rest().Get(eventimageURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
