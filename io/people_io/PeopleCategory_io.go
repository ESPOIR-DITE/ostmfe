package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const peoplecategoryURL = api.BASE_URL + "people_category/"

func CreatePeopleCategory(plcs people.PeopleCategory) (people.PeopleCategory, error) {
	entity := people.PeopleCategory{}
	resp, _ := api.Rest().SetBody(plcs).Post(peoplecategoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePeopleCategory(plcs people.PeopleCategory) (people.PeopleCategory, error) {
	entity := people.PeopleCategory{}
	resp, _ := api.Rest().SetBody(plcs).Post(peoplecategoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleCategory(id string) (people.PeopleCategory, error) {
	entity := people.PeopleCategory{}
	resp, _ := api.Rest().Get(peoplecategoryURL + "read?=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleCategoryWithPplId(id string) ([]people.PeopleCategory, error) {
	entity := []people.PeopleCategory{}
	resp, _ := api.Rest().Get(peoplecategoryURL + "readWithPplId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleCategoryWithCategoryId(id string) ([]people.PeopleCategory, error) {
	entity := []people.PeopleCategory{}
	resp, _ := api.Rest().Get(peoplecategoryURL + "readWithCategoryId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeletePeopleCategory(id string) (people.PeopleCategory, error) {
	entity := people.PeopleCategory{}
	resp, _ := api.Rest().Get(peoplecategoryURL + "delete?=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleCategorys() ([]people.PeopleCategory, error) {
	entity := []people.PeopleCategory{}
	resp, _ := api.Rest().Get(peoplecategoryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
