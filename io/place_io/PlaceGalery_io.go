package place_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/place"
)

const placegalery = api.BASE_URL + "place-galery/"

func CreatePlaceGalery(plcs place.PlaceGallery) (place.PlaceGallery, error) {

	entity := place.PlaceGallery{}
	resp, _ := api.Rest().SetBody(plcs).Post(placegalery + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdatePlaceGalery(plcs place.PlaceGallery) (place.PlaceGallery, error) {

	entity := place.PlaceGallery{}
	resp, _ := api.Rest().SetBody(plcs).Post(placegalery + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPlaceGalery(id string) (place.PlaceGallery, error) {
	entity := place.PlaceGallery{}
	resp, _ := api.Rest().Get(placegalery + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByPlaceGallery(id string) ([]place.PlaceGallery, error) {
	entity := []place.PlaceGallery{}
	resp, _ := api.Rest().Get(placegalery + "readAllByPlaceId?placeId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPlaceGaleryWithPlaceId(id string) (place.PlaceGallery, error) {
	entity := place.PlaceGallery{}
	resp, _ := api.Rest().Get(placegalery + "readByPlaceId?placeId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CountPlaceGaleryWithPlaceId(id string) (place.PlaceGallery, error) {
	entity := place.PlaceGallery{}
	resp, _ := api.Rest().Get(placegalery + "countAllByPlaceId?placeId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePlaceGalery(id string) (place.PlaceGallery, error) {
	entity := place.PlaceGallery{}
	resp, _ := api.Rest().Get(placegalery + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPlaceGalerys() ([]place.PlaceGallery, error) {
	entity := []place.PlaceGallery{}
	resp, _ := api.Rest().Get(placegalery + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
