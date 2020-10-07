package place_io

import (
	"errors"
	"ostmfe/api"
	place2 "ostmfe/domain/place"
)

const placeplaceURl = api.BASE_URL + "place_place/"

func CreatePlacePlace(history place2.PlacePlace) (place2.PlacePlace, error) {
	entity := place2.PlacePlace{}

	resp, _ := api.Rest().SetBody(history).Post(placeplaceURl + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdatePlacePlace(history place2.PlacePlace) (place2.PlacePlace, error) {
	entity := place2.PlacePlace{}

	resp, _ := api.Rest().SetBody(history).Post(placeplaceURl + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlacePlace(id string) (place2.PlacePlace, error) {
	entity := place2.PlacePlace{}

	resp, _ := api.Rest().Get(placeplaceURl + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPlacePlaceOf(placeId string) (place2.PlacePlace, error) {
	entity := place2.PlacePlace{}

	resp, _ := api.Rest().Get(placeplaceURl + "readOf?id=" + placeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePlacePlace(id string) (place2.PlacePlace, error) {
	entity := place2.PlacePlace{}

	resp, _ := api.Rest().Get(placeplaceURl + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlacePlaces() ([]place2.PlacePlace, error) {
	entity := []place2.PlacePlace{}

	resp, _ := api.Rest().Get(placeplaceURl + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
