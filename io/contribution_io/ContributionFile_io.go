package contribution_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/contribution"
)

const contributionfileURL = api.BASE_URL + "contribution-event/"

func CreateContributionFile(contributionObject contribution.ContributionFile) (contribution.ContributionFile, error) {
	entity := contribution.ContributionFile{}
	resp, _ := api.Rest().SetBody(contributionObject).Post(contributionfileURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateContributionFile(commentObject contribution.ContributionFile) (contribution.ContributionFile, error) {
	entity := contribution.ContributionFile{}
	resp, _ := api.Rest().SetBody(commentObject).Post(contributionfileURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionFile(id string) (contribution.ContributionFile, error) {
	entity := contribution.ContributionFile{}
	resp, _ := api.Rest().Get(contributionfileURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteContributionFile(id string) (contribution.ContributionFile, error) {
	entity := contribution.ContributionFile{}
	resp, _ := api.Rest().Get(contributionfileURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadContributionFiles() ([]contribution.ContributionFile, error) {
	entity := []contribution.ContributionFile{}
	resp, _ := api.Rest().Get(contributionfileURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
