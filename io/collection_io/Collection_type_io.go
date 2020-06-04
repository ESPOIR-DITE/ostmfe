package collection_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/collection"
)

const collection_typeURL = api.BASE_URL + "collection_type/"

func CreateCollection_Type(types collection.Collection_type) (collection.Collection_type, error) {
	entity := collection.Collection_type{}
	resp, _ := api.Rest().SetBody(types).Post(collection_typeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCollection_Type(types collection.Collection_type) (collection.Collection_type, error) {
	entity := collection.Collection_type{}
	resp, _ := api.Rest().SetBody(types).Post(collection_typeURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCollection_Type(id string) (collection.Collection_type, error) {
	entity := collection.Collection_type{}
	resp, _ := api.Rest().Get(collection_typeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWithCollectionId(id string) (collection.Collection_type, error) {
	entity := collection.Collection_type{}
	resp, _ := api.Rest().Get(collection_typeURL + "readWihtCollectionId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCollection_Type(id string) (collection.Collection_type, error) {
	entity := collection.Collection_type{}
	resp, _ := api.Rest().Get(collection_typeURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCollection_Types() ([]collection.Collection_type, error) {
	entity := []collection.Collection_type{}
	resp, _ := api.Rest().Get(collection_typeURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
