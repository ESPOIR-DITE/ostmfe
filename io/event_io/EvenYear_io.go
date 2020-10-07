package event_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/event"
)

const evenyearURL = api.BASE_URL + "event_year/"

func CreateEventYear(myEvent event.EventYear) (event.EventYear, error) {
	entity := event.EventYear{}
	resp, _ := api.Rest().SetBody(myEvent).Post(evenyearURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateEventYear(myEvent event.EventYear) (event.EventYear, error) {
	entity := event.EventYear{}
	resp, _ := api.Rest().SetBody(myEvent).Post(evenyearURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadEventYear(id string) (event.EventYear, error) {
	entity := event.EventYear{}
	resp, _ := api.Rest().Get(evenyearURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventYearWithEventId(id string) (event.EventYear, error) {
	entity := event.EventYear{}
	resp, _ := api.Rest().Get(evenyearURL + "readWithEventId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventYearWithYearId(id string) (event.EventYear, error) {
	entity := event.EventYear{}
	resp, _ := api.Rest().Get(evenyearURL + "readWithYearId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventYearsWithYearId(id string) ([]event.EventYear, error) {
	entity := []event.EventYear{}
	resp, _ := api.Rest().Get(evenyearURL + "readWithYearsId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteEventYear(id string) (event.EventYear, error) {
	entity := event.EventYear{}
	resp, _ := api.Rest().Get(evenyearURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadEventYears() ([]event.EventYear, error) {
	//fmt.Println("we are sending the requests to the following backend: ",api.BASE_URL)
	entity := []event.EventYear{}
	resp, _ := api.Rest().Get(evenyearURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CountEventYearWithYearId(yearId string) (int64, error) {
	var entity int64
	resp, _ := api.Rest().Get(evenyearURL + "countWithYearId?yearId=" + yearId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
