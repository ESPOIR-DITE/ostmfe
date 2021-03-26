package place_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/place"
)

const placepageflowURL = api.BASE_URL + "place_page-flow/"

func CreatePlacePageFlow(contributionObject place.PlacePageFlow) (place.PlacePageFlow, error) {
	entity := place.PlacePageFlow{}
	resp, _ := api.Rest().SetBody(contributionObject).Post(placepageflowURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePlacePageFlow(commentObject place.PlacePageFlow) (place.PlacePageFlow, error) {
	entity := place.PlacePageFlow{}
	resp, _ := api.Rest().SetBody(commentObject).Post(placepageflowURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlacePageFlow(id string) (place.PlacePageFlow, error) {
	entity := place.PlacePageFlow{}
	resp, _ := api.Rest().Get(placepageflowURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeletePlacePageFlow(id string) (place.PlacePageFlow, error) {
	entity := place.PlacePageFlow{}
	resp, _ := api.Rest().Get(placepageflowURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlacePageFlows() ([]place.PlacePageFlow, error) {
	entity := []place.PlacePageFlow{}
	resp, _ := api.Rest().Get(placepageflowURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllPlacePageFlowByPlaceId(eventId string) ([]place.PlacePageFlow, error) {
	entity := []place.PlacePageFlow{}
	resp, _ := api.Rest().Get(placepageflowURL + "readWithPlaceId?id=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
