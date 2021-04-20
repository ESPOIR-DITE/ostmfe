package place_io

import (
	"errors"
	"ostmfe/api"
	place2 "ostmfe/domain/place"
)

const placeTypeURL = api.BASE_URL + "place_type/"

func CreatePlaceType(placeType place2.PlaceType) (place2.PlaceType, error) {
	entity := place2.PlaceType{}

	resp, _ := api.Rest().SetBody(placeType).Post(placeTypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdatePlaceType(placeType place2.PlaceType) (place2.PlaceType, error) {
	entity := place2.PlaceType{}

	resp, _ := api.Rest().SetBody(placeType).Post(placeTypeURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceType(id string) (place2.PlaceType, error) {
	entity := place2.PlaceType{}

	resp, _ := api.Rest().Get(placeTypeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPlaceTypeOf(placeId string) (place2.PlaceType, error) {
	entity := place2.PlaceType{}

	resp, _ := api.Rest().Get(placeTypeURL + "readOf?id=" + placeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePlaceType(id string) (place2.PlaceType, error) {
	entity := place2.PlaceType{}

	resp, _ := api.Rest().Get(placeTypeURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceTypes() ([]place2.PlaceType, error) {
	entity := []place2.PlaceType{}

	resp, _ := api.Rest().Get(placeTypeURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
