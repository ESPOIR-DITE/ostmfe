package image_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/image"
)

const imageURL = api.BASE_URL + "image/"

func CreateImage(img image.Images) (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().SetBody(img).Post(imageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func updateImage(img image.Images) (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().SetBody(img).Post(imageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadImage(id string) (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().Get(imageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteImage(id string) (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().Get(imageURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadImages() (image.Images, error) {

	entity := image.Images{}
	resp, _ := api.Rest().Get(imageURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
