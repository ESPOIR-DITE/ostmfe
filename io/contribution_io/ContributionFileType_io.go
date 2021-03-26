package contribution_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/contribution"
)

const contributionfiletypeURL = api.BASE_URL + "contribution-file-type/"

func CreateContributionFileType(contributionObject contribution.ContributionFileType) (contribution.ContributionFile, error) {
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
func UpdateContributionFileType(commentObject contribution.ContributionFileType) (contribution.ContributionFileType, error) {
	entity := contribution.ContributionFileType{}
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
func ReadContributionFileType(id string) (contribution.ContributionFileType, error) {
	entity := contribution.ContributionFileType{}
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
func DeleteContributionFileType(id string) (contribution.ContributionFileType, error) {
	entity := contribution.ContributionFileType{}
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
func ReadContributionFileTypes() ([]contribution.ContributionFileType, error) {
	entity := []contribution.ContributionFileType{}
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
