package group_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/group"
)

const groupmemberURL = api.BASE_URL + "group_member/"

func CreateGroupMember(myEvent group.GroupMember) (group.GroupMember, error) {
	entity := group.GroupMember{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupmemberURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGroupMember(myEvent group.GroupMember) (group.GroupMember, error) {
	entity := group.GroupMember{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupmemberURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGroupMember(id string) (group.GroupMember, error) {
	entity := group.GroupMember{}
	resp, _ := api.Rest().Get(groupmemberURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupMemberAllByGroupId(groupId string) ([]group.GroupMember, error) {
	entity := []group.GroupMember{}
	resp, _ := api.Rest().Get(groupmemberURL + "read-group-id?id=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupMemberByGroupIdAndEmail(groupId, email string) ([]group.GroupMember, error) {
	entity := []group.GroupMember{}
	resp, _ := api.Rest().Get(groupmemberURL + "read-group-id?email=" + email + "groupId=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteGroupMember(id string) (group.GroupMember, error) {
	entity := group.GroupMember{}
	resp, _ := api.Rest().Get(groupmemberURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupMembers() ([]group.GroupMember, error) {
	entity := []group.GroupMember{}
	resp, _ := api.Rest().Get(groupmemberURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
