package member_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/member"
)

const memberURL = api.BASE_URL + "member/"

func CreateMember(M member.Member) (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().SetBody(M).Post(memberURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func UdateMember(M member.Member) (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().SetBody(M).Post(memberURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadMember(id string) (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().Get(memberURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeleteMember(id string) (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().Get(memberURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadMembers() (member.Member, error) {
	entity := member.Member{}
	resp, _ := api.Rest().Get(memberURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
