package event_io

import (
	"errors"
	"ostmfe/api"
	event2 "ostmfe/domain/event"
)

const eventPrtnr = api.BASE_URL + "event_partner/"

func CreateEventPartener(prtnr event2.EventPartener) (event2.EventPartener, error) {

	entity := event2.EventPartener{}
	resp, _ := api.Rest().SetBody(prtnr).Post(eventPrtnr + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventPartener(prtnr event2.EventPartener) (event2.EventPartener, error) {

	entity := event2.EventPartener{}
	resp, _ := api.Rest().SetBody(prtnr).Post(eventPrtnr + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventPartener(id string) (event2.EventPartener, error) {

	entity := event2.EventPartener{}
	resp, _ := api.Rest().Get(eventPrtnr + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteEventPartener(id string) (event2.EventPartener, error) {

	entity := event2.EventPartener{}
	resp, _ := api.Rest().Get(eventPrtnr + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventParteners() (event2.EventPartener, error) {

	entity := event2.EventPartener{}
	resp, _ := api.Rest().Get(eventPrtnr + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
