package contribution_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/contribution"
)

const contributionhistoryURL = api.BASE_URL + "contribution-history/"

func CreateContributionHistory(contributionObject contribution.ContributionHistory) (contribution.ContributionHistory, error) {
	entity := contribution.ContributionHistory{}
	resp, _ := api.Rest().SetBody(contributionObject).Post(contributionhistoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateContributionHistory(commentObject contribution.ContributionHistory) (contribution.ContributionHistory, error) {
	entity := contribution.ContributionHistory{}
	resp, _ := api.Rest().SetBody(commentObject).Post(contributionhistoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionHistory(id string) (contribution.ContributionHistory, error) {
	entity := contribution.ContributionHistory{}
	resp, _ := api.Rest().Get(contributionhistoryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteContributionHistory(id string) (contribution.ContributionHistory, error) {
	entity := contribution.ContributionHistory{}
	resp, _ := api.Rest().Get(contributionhistoryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionHistorys() ([]contribution.ContributionHistory, error) {
	entity := []contribution.ContributionHistory{}
	resp, _ := api.Rest().Get(contributionhistoryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllByHistoryId(projectId string) ([]contribution.ContributionHistory, error) {
	entity := []contribution.ContributionHistory{}
	resp, _ := api.Rest().Get(contributionhistoryURL + "readAllByProjectId?projectId=" + projectId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllByContributionHistoryId(contributionId string) ([]contribution.ContributionHistory, error) {
	entity := []contribution.ContributionHistory{}
	resp, _ := api.Rest().Get(contributionhistoryURL + "readByContributionId?contributionId=" + contributionId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
