package contribution_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/contribution"
)

const contributioneventURL = api.BASE_URL + "contribution-event/"

func CreateContributionEvent(contributionObject contribution.ContributionEvent) (contribution.ContributionEvent, error) {
	entity := contribution.ContributionEvent{}
	resp, _ := api.Rest().SetBody(contributionObject).Post(contributioneventURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateContributionEvent(commentObject contribution.ContributionEvent) (contribution.ContributionEvent, error) {
	entity := contribution.ContributionEvent{}
	resp, _ := api.Rest().SetBody(commentObject).Post(contributioneventURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionEvent(id string) (contribution.ContributionEvent, error) {
	entity := contribution.ContributionEvent{}
	resp, _ := api.Rest().Get(contributioneventURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteContributionEvent(id string) (contribution.ContributionEvent, error) {
	entity := contribution.ContributionEvent{}
	resp, _ := api.Rest().Get(contributioneventURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionEvents() ([]contribution.ContributionEvent, error) {
	entity := []contribution.ContributionEvent{}
	resp, _ := api.Rest().Get(contributioneventURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllByContributionId(contributionString string) ([]contribution.ContributionEvent, error) {
	entity := []contribution.ContributionEvent{}
	resp, _ := api.Rest().Get(contributioneventURL + "findAllByContributionId?contributionId=" + contributionString)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllByEventId(eventId string) ([]contribution.ContributionEvent, error) {
	entity := []contribution.ContributionEvent{}
	resp, _ := api.Rest().Get(contributioneventURL + "findAllByEventId?eventId=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
