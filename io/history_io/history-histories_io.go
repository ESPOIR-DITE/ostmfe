package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const historyhistoriesURL = api.BASE_URL + "historyhistories/"

func CreateHistoryHistory(hist history.HistoryHistories) (history.HistoryHistories, error) {
	entity := history.HistoryHistories{}
	resp, _ := api.Rest().SetBody(hist).Post(historyhistoriesURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateHistoryHistories(hist history.HistoryHistories) (history.HistoryHistories, error) {
	entity := history.HistoryHistories{}
	resp, _ := api.Rest().SetBody(hist).Post(historyhistoriesURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadHistoryHistory(id string) (history.HistoryHistories, error) {
	entity := history.HistoryHistories{}
	resp, _ := api.Rest().Get(historyhistoriesURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistoryHistoriesWithHistoryId(id string) (history.HistoryHistories, error) {
	entity := history.HistoryHistories{}
	resp, _ := api.Rest().Get(historyhistoriesURL + "readWithHistoryId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteHistoryHistories(id string) (history.HistoryHistories, error) {
	entity := history.HistoryHistories{}
	resp, _ := api.Rest().Get(historyhistoriesURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistoryHistories() ([]history.HistoryHistories, error) {
	entity := []history.HistoryHistories{}
	resp, _ := api.Rest().Get(historyhistoriesURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
