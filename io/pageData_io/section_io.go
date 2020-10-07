package pageData_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pageData"
)

const sectionURL = api.BASE_URL + "section/"

func CreateSection(P pageData.SectionBlock) (pageData.SectionBlock, error) {

	entity := pageData.SectionBlock{}
	resp, _ := api.Rest().SetBody(P).Post(sectionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateSection(P pageData.SectionBlock) (pageData.SectionBlock, error) {

	entity := pageData.SectionBlock{}
	resp, _ := api.Rest().SetBody(P).Post(sectionURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadSection(id string) (pageData.SectionBlock, error) {

	entity := pageData.SectionBlock{}
	resp, _ := api.Rest().Get(sectionURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteSection(id string) (pageData.SectionBlock, error) {

	entity := pageData.SectionBlock{}
	resp, _ := api.Rest().Get(sectionURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadSections() ([]pageData.SectionBlock, error) {

	entity := []pageData.SectionBlock{}
	resp, _ := api.Rest().Get(sectionURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
