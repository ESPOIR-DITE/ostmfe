package pageData_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pageData"
)

const pagesectionURL = api.BASE_URL + "page_section/"

func CreatePageSection(P pageData.PageSection) (pageData.PageSection, error) {

	entity := pageData.PageSection{}
	resp, _ := api.Rest().SetBody(P).Post(pagesectionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePageSection(P pageData.PageSection) (pageData.PageSection, error) {

	entity := pageData.PageSection{}
	resp, _ := api.Rest().SetBody(P).Post(pagesectionURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPageSection(id string) (pageData.PageSection, error) {

	entity := pageData.PageSection{}
	resp, _ := api.Rest().Get(pagesectionURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeletePageSection(id string) (pageData.PageSection, error) {

	entity := pageData.PageSection{}
	resp, _ := api.Rest().Get(pagesectionURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPageSections() ([]pageData.PageSection, error) {

	entity := []pageData.PageSection{}
	resp, _ := api.Rest().Get(pagesectionURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPageSectionAllOf(pageId string) ([]pageData.PageSection, error) {
	entity := []pageData.PageSection{}
	resp, _ := api.Rest().Get(pagesectionURL + "readAllOf?id=" + pageId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
