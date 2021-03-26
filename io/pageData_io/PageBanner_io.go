package pageData_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pageData"
)

const pagebannerURL = api.BASE_URL + "page-banner/"

func CreatePageBanner(P pageData.PageBanner) (pageData.PageBanner, error) {
	entity := pageData.PageBanner{}
	resp, _ := api.Rest().SetBody(P).Post(pagebannerURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdatePageBanner(P pageData.PageBanner) (pageData.PageBanner, error) {
	entity := pageData.PageBanner{}
	resp, _ := api.Rest().SetBody(P).Post(pagebannerURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPageBanner(id string) (pageData.PageBanner, error) {
	entity := pageData.PageBanner{}
	resp, _ := api.Rest().Get(pagebannerURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPageBannerWIthPageName(id string) (pageData.PageBanner, error) {
	entity := pageData.PageBanner{}
	resp, _ := api.Rest().Get(pagebannerURL + "readByPageName?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePageBanner(id string) (pageData.PageBanner, error) {

	entity := pageData.PageBanner{}
	resp, _ := api.Rest().Get(pagebannerURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPageBanners() ([]pageData.PageBanner, error) {

	entity := []pageData.PageBanner{}
	resp, _ := api.Rest().Get(pagebannerURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
