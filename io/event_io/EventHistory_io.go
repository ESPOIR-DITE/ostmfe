package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const eventhistoryURL = api.BASE_URL + "event_history/"

func CreateEventHistory(myEvent event.EventHistory) (event.EventHistory, error) {

	entity := event.EventHistory{}
	resp, _ := api.Rest().SetBody(myEvent).Post(eventhistoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventHistory(myEvent event.EventHistory) (event.EventHistory, error) {

	entity := event.EventHistory{}
	resp, _ := api.Rest().SetBody(myEvent).Post(eventhistoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventHistory(id string) (event.EventHistory, error) {

	entity := event.EventHistory{}
	resp, _ := api.Rest().Get(eventhistoryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventHistoryWithEventId(id string) (event.EventHistory, error) {
	entity := event.EventHistory{}
	resp, _ := api.Rest().Get(eventhistoryURL + "readWithEventId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteEventHistory(id string) (event.EventHistory, error) {

	entity := event.EventHistory{}
	resp, _ := api.Rest().Get(eventhistoryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventHistories() ([]event.EventHistory, error) {
	entity := []event.EventHistory{}
	resp, _ := api.Rest().Get(evenT + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
