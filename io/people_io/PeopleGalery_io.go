package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const peoplegaleryURL = api.BASE_URL + "people-galery/"

func CreatePeopleGalery(P people.PeopleGalery) (people.PeopleGalery, error) {
	entity := people.PeopleGalery{}
	resp, _ := api.Rest().SetBody(P).Post(peoplegaleryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePeopleGalery(P people.PeopleGalery) (people.PeopleGalery, error) {
	entity := people.PeopleGalery{}
	resp, _ := api.Rest().SetBody(P).Post(peoplegaleryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleGalery(id string) (people.PeopleGalery, error) {
	entity := people.PeopleGalery{}
	resp, _ := api.Rest().Get(peoplegaleryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPeopleGaleryWithPeopleId(id string) (people.PeopleGalery, error) {
	entity := people.PeopleGalery{}
	resp, _ := api.Rest().Get(peoplegaleryURL + "readByPeopleId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByPeopleIdGalery(id string) ([]people.PeopleGalery, error) {
	entity := []people.PeopleGalery{}
	resp, _ := api.Rest().Get(peoplegaleryURL + "readAllByPeopleId?peopleId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CountPeopleGalery(id string) (people.PeopleGalery, error) {
	entity := people.PeopleGalery{}
	resp, _ := api.Rest().Get(peoplegaleryURL + "countByPeopleId?peopleId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePeopleGalery(id string) (people.PeopleGalery, error) {
	entity := people.PeopleGalery{}
	resp, _ := api.Rest().Get(peoplegaleryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleGaleries() ([]people.PeopleGalery, error) {
	entity := []people.PeopleGalery{}
	resp, _ := api.Rest().Get(peoplegaleryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
