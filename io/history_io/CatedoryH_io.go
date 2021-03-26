package history_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/history"
)

const categoryhURL = api.BASE_URL + "categoryH/"

func CreateCategoryH(hist history.CategoryH) (history.CategoryH, error) {
	entity := history.CategoryH{}
	resp, _ := api.Rest().SetBody(hist).Post(categoryhURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateCategoryH(hist history.CategoryH) (history.CategoryH, error) {
	entity := history.CategoryH{}
	resp, _ := api.Rest().SetBody(hist).Post(categoryhURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadCategoryH(id string) (history.CategoryH, error) {
	entity := history.CategoryH{}
	resp, _ := api.Rest().Get(categoryhURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteCategoryH(id string) (history.CategoryH, error) {
	entity := history.CategoryH{}
	resp, _ := api.Rest().Get(categoryhURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadCategoryHs() ([]history.CategoryH, error) {
	entity := []history.CategoryH{}
	resp, _ := api.Rest().Get(categoryhURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
