package contribution_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/contribution"
)

const contributionprojectURL = api.BASE_URL + "contribution-event/"

func CreateContributionProject(contributionObject contribution.ContributionProject) (contribution.ContributionProject, error) {
	entity := contribution.ContributionProject{}
	resp, _ := api.Rest().SetBody(contributionObject).Post(contributionprojectURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateContributionProject(commentObject contribution.ContributionProject) (contribution.ContributionProject, error) {
	entity := contribution.ContributionProject{}
	resp, _ := api.Rest().SetBody(commentObject).Post(contributionprojectURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionProject(id string) (contribution.ContributionProject, error) {
	entity := contribution.ContributionProject{}
	resp, _ := api.Rest().Get(contributionprojectURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteContributionProject(id string) (contribution.ContributionProject, error) {
	entity := contribution.ContributionProject{}
	resp, _ := api.Rest().Get(contributionprojectURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionProjects() ([]contribution.ContributionProject, error) {
	entity := []contribution.ContributionProject{}
	resp, _ := api.Rest().Get(contributionprojectURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllByProjectId(projectId string) ([]contribution.ContributionProject, error) {
	entity := []contribution.ContributionProject{}
	resp, _ := api.Rest().Get(contributionprojectURL + "readAllByProjectId?projectId=" + projectId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllByContributionProjectId(contributionId string) ([]contribution.ContributionProject, error) {
	entity := []contribution.ContributionProject{}
	resp, _ := api.Rest().Get(contributionprojectURL + "readByContributionId?contributionId=" + contributionId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
