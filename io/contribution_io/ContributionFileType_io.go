package contribution_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/contribution"
)

const contributionfiletypeURL = api.BASE_URL + "contribution-event/"

func CreateContributionFileType(contributionObject contribution.ContributionType) (contribution.ContributionFile, error) {
	entity := contribution.ContributionFile{}
	resp, _ := api.Rest().SetBody(contributionObject).Post(contributionfiletypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateContributionFileType(commentObject contribution.ContributionType) (contribution.ContributionType, error) {
	entity := contribution.ContributionType{}
	resp, _ := api.Rest().SetBody(commentObject).Post(contributionfiletypeURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionFileType(id string) (contribution.ContributionType, error) {
	entity := contribution.ContributionType{}
	resp, _ := api.Rest().Get(contributionfiletypeURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteContributionFileType(id string) (contribution.ContributionType, error) {
	entity := contribution.ContributionType{}
	resp, _ := api.Rest().Get(contributionfiletypeURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionFileTypes() ([]contribution.ContributionType, error) {
	entity := []contribution.ContributionType{}
	resp, _ := api.Rest().Get(contributionfiletypeURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
