package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const historypageflowURL = api.BASE_URL + "history_page_low/"

func CreateHistoryPageFLow(hist history.HistoryPageFlow) (history.HistoryPageFlow, error) {
	entity := history.HistoryPageFlow{}
	resp, _ := api.Rest().SetBody(hist).Post(historypageflowURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateHistoryPageFLow(hist history.HistoryPageFlow) (history.HistoryPageFlow, error) {
	entity := history.HistoryPageFlow{}
	resp, _ := api.Rest().SetBody(hist).Post(historypageflowURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistoryPageFLow(id string) (history.HistoryPageFlow, error) {
	entity := history.HistoryPageFlow{}
	resp, _ := api.Rest().Get(historypageflowURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistoryPageFLowWithHistoryId(id string) (history.HistoryPageFlow, error) {
	entity := history.HistoryPageFlow{}
	resp, _ := api.Rest().Get(historypageflowURL + "readWithHistoryId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadHistoryPageFLowsWithHistoryId(id string) ([]history.HistoryPageFlow, error) {
	entity := []history.HistoryPageFlow{}
	resp, _ := api.Rest().Get(historypageflowURL + "readAllWithHistoryId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteHistoryPageFLow(id string) (history.HistoryPageFlow, error) {
	entity := history.HistoryPageFlow{}
	resp, _ := api.Rest().Get(historypageflowURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadHistoryPageFLows() (history.HistoryPageFlow, error) {
	entity := history.HistoryPageFlow{}
	resp, _ := api.Rest().Get(historypageflowURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
