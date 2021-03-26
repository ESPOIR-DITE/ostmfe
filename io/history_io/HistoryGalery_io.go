package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const historygalleryURL = api.BASE_URL + "history-galery/"

func CreateHistoryGallery(hist history.HistoryGalery) (history.HistoryGalery, error) {
	entity := history.HistoryGalery{}
	resp, _ := api.Rest().SetBody(hist).Post(historygalleryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateHistoryGallery(hist history.HistoryGalery) (history.HistoryGalery, error) {
	entity := history.HistoryGalery{}
	resp, _ := api.Rest().SetBody(hist).Post(historygalleryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistoryGallery(id string) (history.HistoryGalery, error) {
	entity := history.HistoryGalery{}
	resp, _ := api.Rest().Get(historygalleryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistoryGalleryWithHistoryId(historyId string) (history.HistoryGalery, error) {
	entity := history.HistoryGalery{}
	resp, _ := api.Rest().Get(historygalleryURL + "readByHistoryId?historyId=" + historyId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllHistoryGallery(historyId string) ([]history.HistoryGalery, error) {
	entity := []history.HistoryGalery{}
	resp, _ := api.Rest().Get(historygalleryURL + "readAllByHistoryId?historyId=" + historyId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CountHistoryGalleryWithHistoryId(historyId string) (history.HistoryGalery, error) {
	entity := history.HistoryGalery{}
	resp, _ := api.Rest().Get(historygalleryURL + "countByHistoryId?historyId=" + historyId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteHistoryGallery(id string) (history.HistoryGalery, error) {
	entity := history.HistoryGalery{}
	resp, _ := api.Rest().Get(historygalleryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistoryGalleries() ([]history.HistoryGalery, error) {
	entity := []history.HistoryGalery{}
	resp, _ := api.Rest().Get(historygalleryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
