package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const historyURL = api.BASE_URL + "history/"

func CreateHistory(hist history.History) (history.History, error) {
	entity := history.History{}
	resp, _ := api.Rest().SetBody(hist).Post(historyURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateHistory(hist history.History) (history.History, error) {
	entity := history.History{}
	resp, _ := api.Rest().SetBody(hist).Post(historyURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistory(id string) (history.History, error) {
	entity := history.History{}
	resp, _ := api.Rest().Get(historyURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteHistory(id string) (history.History, error) {
	entity := history.History{}
	resp, _ := api.Rest().Get(historyURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistorys() ([]history.History, error) {
	entity := []history.History{}
	resp, _ := api.Rest().Get(historyURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
