package admin

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pages"
)

const eventURL = api.BASE_URL + "admin-pages/event/"

func GetEventEditData(id string) (pages.EventViewData, error) {
	entity := pages.EventViewData{}
	resp, _ := api.Rest().Get(eventURL + "edit?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
