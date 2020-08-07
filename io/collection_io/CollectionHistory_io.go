package collection_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/collection"
)

const collectionhistoryURL = api.BASE_URL + "collection_history/"

func CreateCollectionHistory(collectionHistory collection.CollectionHistory) (collection.CollectionHistory, error) {
	entity := collection.CollectionHistory{}
	resp, _ := api.Rest().SetBody(collectionHistory).Post(collectionhistoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCollectionHistory(image collection.CollectionHistory) (collection.CollectionHistory, error) {
	entity := collection.CollectionHistory{}
	resp, _ := api.Rest().SetBody(image).Post(collectionhistoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadCollectionHistory(id string) (collection.CollectionHistory, error) {
	entity := collection.CollectionHistory{}
	resp, _ := api.Rest().Get(collectionhistoryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCollectionHistoryWithCollectionId(id string) (collection.CollectionHistory, error) {
	entity := collection.CollectionHistory{}
	resp, _ := api.Rest().Get(collectionhistoryURL + "readWith?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteCollectionHistory(id string) (collection.CollectionHistory, error) {
	entity := collection.CollectionHistory{}
	resp, _ := api.Rest().Get(collectionhistoryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadCollectionHistorys() (collection.CollectionHistory, error) {
	entity := collection.CollectionHistory{}
	resp, _ := api.Rest().Get(collectionhistoryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
