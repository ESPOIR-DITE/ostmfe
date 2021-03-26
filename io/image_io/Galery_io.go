package image_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/image"
)

const galeryURL = api.BASE_URL + "galery/"

func CreateGalery(img image.Galery) (image.Galery, error) {

	entity := image.Galery{}
	resp, _ := api.Rest().SetBody(img).Post(galeryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateGallery(img image.Galery) (image.Galery, error) {

	entity := image.Galery{}
	resp, _ := api.Rest().SetBody(img).Post(galeryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGallery(id string) (image.Galery, error) {

	entity := image.Galery{}
	resp, _ := api.Rest().Get(galeryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteGalery(id string) (image.Galery, error) {

	entity := image.Galery{}
	resp, _ := api.Rest().Get(galeryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGaleries() (image.Galery, error) {

	entity := image.Galery{}
	resp, _ := api.Rest().Get(galeryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
