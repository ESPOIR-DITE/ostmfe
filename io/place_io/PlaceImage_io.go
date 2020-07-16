package place_io

import (
	"errors"
	"ostmfe/api"
	place2 "ostmfe/domain/place"
)

const placeimageURL = api.BASE_URL + "place_image/"

func CreatePlaceImage(helper place2.PlaceImageHelper) (place2.PlaceImage, error) {
	entity := place2.PlaceImage{}

	resp, _ := api.Rest().SetBody(helper).Post(placeimageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePlaceImage(image place2.PlaceImage) (place2.PlaceImage, error) {
	entity := place2.PlaceImage{}

	resp, _ := api.Rest().SetBody(image).Post(placeimageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceImage(id string) (place2.PlaceImage, error) {
	entity := place2.PlaceImage{}

	resp, _ := api.Rest().Get(placeimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceImageAllOf(placeId string) ([]place2.PlaceImage, error) {
	entity := []place2.PlaceImage{}

	resp, _ := api.Rest().Get(placeimageURL + "readAllOf?id=" + placeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeletePlaceImage(id string) (place2.PlaceImage, error) {
	entity := place2.PlaceImage{}

	resp, _ := api.Rest().Get(placeimageURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPlaceImages() ([]place2.PlaceImage, error) {
	entity := []place2.PlaceImage{}

	resp, _ := api.Rest().Get(placeimageURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
