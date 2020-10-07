package group_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/group"
)

const grouppartnerURL = api.BASE_URL + "group_partner/"

func CreateGroupPartner(myEvent group.GroupPartener) (group.GroupPartener, error) {
	entity := group.GroupPartener{}
	resp, _ := api.Rest().SetBody(myEvent).Post(grouppartnerURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGroupPartner(myEvent group.GroupPartener) (group.GroupPartener, error) {
	entity := group.GroupPartener{}
	resp, _ := api.Rest().SetBody(myEvent).Post(grouppartnerURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGroupPartner(id string) (group.GroupPartener, error) {
	entity := group.GroupPartener{}
	resp, _ := api.Rest().Get(grouppartnerURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupPartnerWithGroupId(groupId string) ([]group.GroupPartener, error) {
	entity := []group.GroupPartener{}
	resp, _ := api.Rest().Get(grouppartnerURL + "readWith?id=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteGroupPartner(id string) (group.GroupPartener, error) {
	entity := group.GroupPartener{}
	resp, _ := api.Rest().Get(grouppartnerURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupPartners() ([]group.GroupPartener, error) {
	entity := []group.GroupPartener{}
	resp, _ := api.Rest().Get(grouppartnerURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
