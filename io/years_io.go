package io

import (
	"errors"
	"ostmfe/api"
	museum "ostmfe/domain"
)

const yearURL = api.BASE_URL + "years/"

func CreateYear(M museum.Years) (museum.Years, error) {

	entity := museum.Years{}
	resp, _ := api.Rest().SetBody(M).Post(yearURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UpdateYear(M museum.Years) (museum.Years, error) {

	entity := museum.Years{}
	resp, _ := api.Rest().SetBody(M).Post(yearURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadYear(id string) (museum.Years, error) {

	entity := museum.Years{}
	resp, _ := api.Rest().Get(yearURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteYear(id string) (museum.Years, error) {

	entity := museum.Years{}
	resp, _ := api.Rest().Get(yearURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadYears() ([]museum.Years, error) {
	entity := []museum.Years{}
	resp, _ := api.Rest().Get(yearURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
