package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/contribution"
)

const eventpageflowURL = api.BASE_URL + "event_page-flow/"

func CreateEventPageFlow(contributionObject contribution.EventPageFlow) (contribution.EventPageFlow, error) {
	entity := contribution.EventPageFlow{}
	resp, _ := api.Rest().SetBody(contributionObject).Post(eventpageflowURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventPageFlow(commentObject contribution.EventPageFlow) (contribution.EventPageFlow, error) {
	entity := contribution.EventPageFlow{}
	resp, _ := api.Rest().SetBody(commentObject).Post(eventpageflowURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPageFlow(id string) (contribution.EventPageFlow, error) {
	entity := contribution.EventPageFlow{}
	resp, _ := api.Rest().Get(eventpageflowURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteEventPageFlow(id string) (contribution.EventPageFlow, error) {
	entity := contribution.EventPageFlow{}
	resp, _ := api.Rest().Get(eventpageflowURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventPageFlows() ([]contribution.EventPageFlow, error) {
	entity := []contribution.EventPageFlow{}
	resp, _ := api.Rest().Get(eventpageflowURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllEventPageFlowByEventId(eventId string) ([]contribution.EventPageFlow, error) {
	entity := []contribution.EventPageFlow{}
	resp, _ := api.Rest().Get(eventpageflowURL + "readWithEventId?id=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
