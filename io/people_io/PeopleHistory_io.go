package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/people"
)

const peoplehistoryURL = api.BASE_URL + "people_history"

func CreatePeopleHistory(history people.PeopleHistory) (people.PeopleHistory, error) {
	entity := people.PeopleHistory{}
	resp, _ := api.Rest().SetBody(history).Post(peoplehistoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdatePeopleHistory(history people.PeopleHistory) (people.PeopleHistory, error) {
	entity := people.PeopleHistory{}
	resp, _ := api.Rest().SetBody(history).Post(peoplehistoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleHistory(id string) (people.PeopleHistory, error) {
	entity := people.PeopleHistory{}
	resp, _ := api.Rest().Get(peoplehistoryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleHistoryWithPplId(id string) (people.PeopleHistory, error) {
	entity := people.PeopleHistory{}
	resp, _ := api.Rest().Get(peoplehistoryURL + "readWithPplId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeletePeopleHistory(id string) (people.PeopleHistory, error) {
	entity := people.PeopleHistory{}
	resp, _ := api.Rest().Get(peoplehistoryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleHistorys(id string) ([]people.PeopleHistory, error) {
	entity := []people.PeopleHistory{}
	resp, _ := api.Rest().Get(peoplehistoryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
