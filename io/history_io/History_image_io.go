package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const historyimageURL = api.BASE_URL + "historyImages/"

func CreateHistoryImage(hist history.HistoryImageHelper) (history.HistoryImage, error) {
	entity := history.HistoryImage{}
	resp, _ := api.Rest().SetBody(hist).Post(historyimageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateHistoryImage(hist history.HistoryImageHelper) (history.HistoryImage, error) {
	entity := history.HistoryImage{}
	resp, _ := api.Rest().SetBody(hist).Post(historyimageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistoryImage(id string) (history.HistoryImage, error) {
	entity := history.HistoryImage{}
	resp, _ := api.Rest().Get(historyimageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistoryImageWithHistoryId(id string) (history.HistoryImage, error) {
	entity := history.HistoryImage{}
	resp, _ := api.Rest().Get(historyimageURL + "readWithHistoryId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadHistoryImagesWithHistoryId(id string) ([]history.HistoryImage, error) {
	entity := []history.HistoryImage{}
	resp, _ := api.Rest().Get(historyimageURL + "readAllWithHistoryId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteHistoryImage(id string) (history.HistoryImage, error) {
	entity := history.HistoryImage{}
	resp, _ := api.Rest().Get(historyimageURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistoryImages() (history.HistoryImage, error) {
	entity := history.HistoryImage{}
	resp, _ := api.Rest().Get(historyimageURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
