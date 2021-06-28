package image_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/image"
)

const galeryURL = api.BASE_URL + "galery/"

func CreateGalery(img image.Gallery) (image.Gallery, error) {

	entity := image.Gallery{}
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
func UpdateGallery(img image.Gallery) (image.Gallery, error) {

	entity := image.Gallery{}
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
func ReadGalleryH(id string) (image.GalleryHelper, error) {
	entity := image.GalleryHelper{}
	resp, _ := api.Rest().Get(galeryURL + "readH?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}

//Deprecated
func ReadGallery(id string) (image.Gallery, error) {
	entity := image.Gallery{}
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
func DeleteGalery(id string) (image.Gallery, error) {

	entity := image.Gallery{}
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
func ReadGaleries() (image.Gallery, error) {

	entity := image.Gallery{}
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
