package collection_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/collection"
)

const collectionImg = api.BASE_URL + "collection_image/"

func CreateCollectionImg(image collection.CollectionImageHelper) (collection.Collection_image, error) {

	entity := collection.Collection_image{}
	resp, _ := api.Rest().SetBody(image).Post(collectionImg + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCollectionImg(image collection.Collection_image) (collection.Collection_image, error) {
	entity := collection.Collection_image{}
	resp, _ := api.Rest().SetBody(image).Post(collectionImg + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadCollectionImg(id string) (collection.Collection_image, error) {
	entity := collection.Collection_image{}
	resp, _ := api.Rest().Get(collectionImg + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCollectionImgWithCollectionId(id string) (collection.Collection_image, error) {
	entity := collection.Collection_image{}
	resp, _ := api.Rest().Get(collectionImg + "readWith?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadCollectionImgsWithCollectionId(id string) ([]collection.Collection_image, error) {
	entity := []collection.Collection_image{}
	resp, _ := api.Rest().Get(collectionImg + "readsWith?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteCollectionImg(id string) (collection.Collection_image, error) {
	entity := collection.Collection_image{}
	resp, _ := api.Rest().Get(collectionImg + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadCollectionImgs() ([]collection.Collection_image, error) {
	entity := []collection.Collection_image{}
	resp, _ := api.Rest().Get(collectionImg + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
