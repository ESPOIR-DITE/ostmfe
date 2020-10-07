package place_io

import (
	"errors"
	"ostmfe/api"
	place2 "ostmfe/domain/place"
)

const placehistoryURl = api.BASE_URL + "place_history/"

func CreatePlaceHistpory(history place2.PlaceHistories) (place2.PlaceHistories, error) {
	entity := place2.PlaceHistories{}

	resp, _ := api.Rest().SetBody(history).Post(placehistoryURl + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdatePlaceHistpory(history place2.PlaceHistories) (place2.PlaceHistories, error) {
	entity := place2.PlaceHistories{}

	resp, _ := api.Rest().SetBody(history).Post(placehistoryURl + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceHistpory(id string) (place2.PlaceHistories, error) {
	entity := place2.PlaceHistories{}

	resp, _ := api.Rest().Get(placehistoryURl + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPlaceHistporyOf(placeId string) (place2.PlaceHistories, error) {
	entity := place2.PlaceHistories{}

	resp, _ := api.Rest().Get(placehistoryURl + "readOf?id=" + placeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePlaceHistpory(id string) (place2.PlaceHistories, error) {
	entity := place2.PlaceHistories{}

	resp, _ := api.Rest().Get(placehistoryURl + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceHistporys() ([]place2.PlaceHistories, error) {
	entity := []place2.PlaceHistories{}

	resp, _ := api.Rest().Get(placehistoryURl + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
