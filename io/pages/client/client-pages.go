package client

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pages"
)

const clientPage = api.BASE_URL + "client-pages/"

func HomeClientPage() (pages.ClientLandingPageData, error) {
	entity := pages.ClientLandingPageData{}
	resp, _ := api.Rest().Get(clientPage + "home")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func AboutClientPage() (pages.AboutUsPageData, error) {
	entity := pages.AboutUsPageData{}
	resp, _ := api.Rest().Get(clientPage + "aboutUs")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func EVentClientPage() (pages.EventPageData, error) {
	entity := pages.EventPageData{}
	resp, _ := api.Rest().Get(clientPage + "event")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func PeopleClientPage(peopleId string) (pages.PeoplePageData, error) {
	entity := pages.PeoplePageData{}
	resp, _ := api.Rest().Get(clientPage + "people-single?peopleId=" + peopleId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func PeopleHomePage() (pages.PeopleHomePage, error) {
	entity := pages.PeopleHomePage{}
	resp, _ := api.Rest().Get(clientPage + "people-home")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetGroupClientSingleData(groupId string) (pages.GroupClientSingleData, error) {
	entity := pages.GroupClientSingleData{}
	resp, _ := api.Rest().Get(clientPage + "group-single?groupId=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
