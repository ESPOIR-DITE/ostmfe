package pageData_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pageData"
)

const pagedataURL = api.BASE_URL + "page_data/"

func CreatePageData(P pageData.PageData) (pageData.PageData, error) {

	entity := pageData.PageData{}
	resp, _ := api.Rest().SetBody(P).Post(pagedataURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePageData(P pageData.PageData) (pageData.PageData, error) {

	entity := pageData.PageData{}
	resp, _ := api.Rest().SetBody(P).Post(pagedataURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPageData(id string) (pageData.PageData, error) {
	entity := pageData.PageData{}
	resp, _ := api.Rest().Get(pagedataURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPageDataWIthName(id string) (pageData.PageData, error) {
	entity := pageData.PageData{}
	resp, _ := api.Rest().Get(pagedataURL + "readWithName?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePageData(id string) (pageData.PageData, error) {

	entity := pageData.PageData{}
	resp, _ := api.Rest().Get(pagedataURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPageDatas() ([]pageData.PageData, error) {

	entity := []pageData.PageData{}
	resp, _ := api.Rest().Get(pagedataURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
