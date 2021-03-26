package place_io

import (
	"errors"
	"ostmfe/api"
	place2 "ostmfe/domain/place"
)

const placecotegoryURl = api.BASE_URL + "place_category/"

func CreatePlaceCategory(history place2.PlaceCategory) (place2.PlaceCategory, error) {
	entity := place2.PlaceCategory{}

	resp, _ := api.Rest().SetBody(history).Post(placecotegoryURl + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdatePlaceCategory(history place2.PlaceCategory) (place2.PlaceCategory, error) {
	entity := place2.PlaceCategory{}

	resp, _ := api.Rest().SetBody(history).Post(placecotegoryURl + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceCategory(id string) (place2.PlaceCategory, error) {
	entity := place2.PlaceCategory{}

	resp, _ := api.Rest().Get(placecotegoryURl + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPlaceCategoryOf(placeId string) (place2.PlaceCategory, error) {
	entity := place2.PlaceCategory{}

	resp, _ := api.Rest().Get(placecotegoryURl + "readOf?id=" + placeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePlaceCategory(id string) (place2.PlaceCategory, error) {
	entity := place2.PlaceCategory{}

	resp, _ := api.Rest().Get(placecotegoryURl + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceCategories() ([]place2.PlaceCategory, error) {
	entity := []place2.PlaceCategory{}

	resp, _ := api.Rest().Get(placecotegoryURl + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
