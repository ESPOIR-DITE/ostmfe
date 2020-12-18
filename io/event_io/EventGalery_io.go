package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventgaleryURL = api.BASE_URL + "event-galery/"

func CreateEventGalery(myEvent event.EventGalery) (event.EventGalery, error) {
	entity := event.EventGalery{}
	resp, _ := api.Rest().SetBody(myEvent).Post(eventgaleryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateEventGalery(myEvent event.EventGalery) (event.EventGalery, error) {
	entity := event.EventGalery{}
	resp, _ := api.Rest().SetBody(myEvent).Post(eventgaleryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventGalery(id string) (event.EventGalery, error) {
	entity := event.EventGalery{}
	resp, _ := api.Rest().Get(eventgaleryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventGaleryWitheventId(id string) (event.EventGalery, error) {
	entity := event.EventGalery{}
	resp, _ := api.Rest().Get(eventgaleryURL + "readByEventId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllEventGalleryWithEventId(eventId string) ([]event.EventGalery, error) {
	entity := []event.EventGalery{}
	resp, _ := api.Rest().Get(eventgaleryURL + "readAllByEventId?eventId=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteEventGalery(id string) (event.EventGalery, error) {
	entity := event.EventGalery{}
	resp, _ := api.Rest().Get(eventgaleryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadEventGalerys() ([]event.EventGalery, error) {
	entity := []event.EventGalery{}
	resp, _ := api.Rest().Get(eventgaleryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
