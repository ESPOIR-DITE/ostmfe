package contribution_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/contribution"
)

const contributionURL = api.BASE_URL + "contribution/"

func CreateContribution(contributionObject contribution.Contribution) (contribution.Contribution, error) {
	entity := contribution.Contribution{}
	resp, _ := api.Rest().SetBody(contributionObject).Post(contributionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateContribution(commentObject contribution.Contribution) (contribution.Contribution, error) {
	entity := contribution.Contribution{}
	resp, _ := api.Rest().SetBody(commentObject).Post(contributionURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContribution(id string) (contribution.Contribution, error) {
	entity := contribution.Contribution{}
	resp, _ := api.Rest().Get(contributionURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteContribution(id string) (contribution.Contribution, error) {
	entity := contribution.Contribution{}
	resp, _ := api.Rest().Get(contributionURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributions() ([]contribution.Contribution, error) {
	entity := []contribution.Contribution{}
	resp, _ := api.Rest().Get(contributionURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
