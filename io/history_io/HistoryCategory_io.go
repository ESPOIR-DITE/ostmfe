package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const historycategoryhURL = api.BASE_URL + "history-category/"

func CreateHistoryCategory(hist history.HistoryCategory) (history.HistoryCategory, error) {
	entity := history.HistoryCategory{}
	resp, _ := api.Rest().SetBody(hist).Post(historycategoryhURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateHistoryCategory(hist history.HistoryCategory) (history.HistoryCategory, error) {
	entity := history.HistoryCategory{}
	resp, _ := api.Rest().SetBody(hist).Post(historycategoryhURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistoryCategory(id string) (history.HistoryCategory, error) {
	entity := history.HistoryCategory{}
	resp, _ := api.Rest().Get(historycategoryhURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllHistoryCategoryByHistoryId(historyId string) ([]history.HistoryCategory, error) {
	entity := []history.HistoryCategory{}
	resp, _ := api.Rest().Get(historycategoryhURL + "readAllByHistoryId?historyId=" + historyId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllHistoryCategoryByCategoryId(categoryId string) ([]history.HistoryCategory, error) {
	entity := []history.HistoryCategory{}
	resp, _ := api.Rest().Get(historycategoryhURL + "findAllByCategoryId?categoryId=" + categoryId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteHistoryCategory(id string) (history.HistoryCategory, error) {
	entity := history.HistoryCategory{}
	resp, _ := api.Rest().Get(historycategoryhURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadHistoryCategories() ([]history.HistoryCategory, error) {
	entity := []history.HistoryCategory{}
	resp, _ := api.Rest().Get(historycategoryhURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
