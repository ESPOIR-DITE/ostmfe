package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const categoryURl = api.BASE_URL + "category"

func CreateCategory(history people.Category) (people.Category, error) {
	entity := people.Category{}
	resp, _ := api.Rest().SetBody(history).Post(categoryURl + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCategory(history people.Category) (people.Category, error) {
	entity := people.Category{}
	resp, _ := api.Rest().SetBody(history).Post(categoryURl + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCategory(id string) (people.Category, error) {
	entity := people.Category{}
	resp, _ := api.Rest().Get(categoryURl + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCategory(id string) (people.Category, error) {
	entity := people.Category{}
	resp, _ := api.Rest().Get(categoryURl + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCategories() ([]people.Category, error) {
	entity := []people.Category{}
	resp, _ := api.Rest().Get(categoryURl + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
