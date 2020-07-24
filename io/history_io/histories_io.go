package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const historiesURL = api.BASE_URL + "histories/"

func CreateHistorie(hist history.Histories) (history.Histories, error) {
	entity := history.Histories{}
	resp, _ := api.Rest().SetBody(hist).Post(historiesURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateHistorie(hist history.Histories) (history.Histories, error) {
	entity := history.Histories{}
	resp, _ := api.Rest().SetBody(hist).Post(historiesURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistorie(id string) (history.Histories, error) {
	entity := history.Histories{}
	resp, _ := api.Rest().Get(historiesURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteHistorie(id string) (history.Histories, error) {
	entity := history.Histories{}
	resp, _ := api.Rest().Get(historiesURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistories() ([]history.Histories, error) {
	entity := []history.Histories{}
	resp, _ := api.Rest().Get(historiesURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
