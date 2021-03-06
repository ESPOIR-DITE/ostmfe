package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const pple = api.BASE_URL + "people/"

func CreatePeople(P people.People) (people.People, error) {
	entity := people.People{}
	resp, _ := api.Rest().SetBody(P).Post(pple + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePeople(P people.People) (people.People, error) {
	entity := people.People{}
	resp, _ := api.Rest().SetBody(P).Post(pple + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeople(id string) (people.People, error) {
	entity := people.People{}
	resp, _ := api.Rest().Get(pple + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeletePeople(id string) (people.People, error) {
	entity := people.People{}
	resp, _ := api.Rest().Get(pple + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeoples() ([]people.People, error) {
	entity := []people.People{}
	resp, _ := api.Rest().Get(pple + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetAggregatedPeople(peopleId string) (people.PeopleAggregate, error) {
	entity := people.PeopleAggregate{}
	resp, _ := api.Rest().Get(pple + "readAggregated?peopleId=" + peopleId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
