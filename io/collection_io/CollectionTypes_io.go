package collection_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/collection"
)

const collectiontypeURL = api.BASE_URL + "collection_types/"

func CreateCollectionTyupe(types collection.CollectionTypes) (collection.CollectionTypes, error) {
	entity := collection.CollectionTypes{}
	resp, _ := api.Rest().SetBody(types).Post(collectiontypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCollectionTyupe(types collection.CollectionTypes) (collection.CollectionTypes, error) {
	entity := collection.CollectionTypes{}
	resp, _ := api.Rest().SetBody(types).Post(collectiontypeURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCollectionTyupe(id string) (collection.CollectionTypes, error) {
	entity := collection.CollectionTypes{}
	resp, _ := api.Rest().Get(collectiontypeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCollectionTyupe(id string) (collection.CollectionTypes, error) {
	entity := collection.CollectionTypes{}
	resp, _ := api.Rest().Get(collectiontypeURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCollectionTyupes() ([]collection.CollectionTypes, error) {
	entity := []collection.CollectionTypes{}
	resp, _ := api.Rest().Get(collectiontypeURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
