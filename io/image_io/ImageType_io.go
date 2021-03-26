package image_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/image"
)

const imageTypeURL = api.BASE_URL + "ImageType/"

func CreateImageType(img image.ImageType) (image.ImageType, error) {
	entity := image.ImageType{}
	resp, _ := api.Rest().SetBody(img).Post(imageTypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateImageType(img image.ImageType) (image.ImageType, error) {
	entity := image.ImageType{}
	resp, _ := api.Rest().SetBody(img).Post(imageTypeURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadImageType(id string) (image.ImageType, error) {

	entity := image.ImageType{}
	resp, _ := api.Rest().Get(imageTypeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteImageType(id string) (image.ImageType, error) {

	entity := image.ImageType{}
	resp, _ := api.Rest().Get(imageTypeURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadImageTypes() ([]image.ImageType, error) {
	entity := []image.ImageType{}
	resp, _ := api.Rest().Get(imageTypeURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadImageTypeWithName(name string) (image.ImageType, error) {
	entity := image.ImageType{}
	resp, _ := api.Rest().Get(imageTypeURL + "findByName?name=" + name)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
