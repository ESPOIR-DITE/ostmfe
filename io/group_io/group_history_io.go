package group_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/group"
)

const grouphistoryURL = api.BASE_URL + "group_history/"

func CreateGroupHistory(myEvent group.GroupHistory) (group.GroupHistory, error) {
	entity := group.GroupHistory{}
	resp, _ := api.Rest().SetBody(myEvent).Post(grouphistoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGroupHistory(myEvent group.GroupHistory) (group.GroupHistory, error) {
	entity := group.GroupHistory{}
	resp, _ := api.Rest().SetBody(myEvent).Post(grouphistoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGroupHistory(id string) (group.GroupHistory, error) {
	entity := group.GroupHistory{}
	resp, _ := api.Rest().Get(grouphistoryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupHistoryWithGroupId(groupId string) (group.GroupHistory, error) {
	entity := group.GroupHistory{}
	resp, _ := api.Rest().Get(grouphistoryURL + "readWith?id=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteGroupHistory(id string) (group.GroupHistory, error) {
	entity := group.GroupHistory{}
	resp, _ := api.Rest().Get(grouphistoryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupHistorys() ([]group.GroupHistory, error) {
	entity := []group.GroupHistory{}
	resp, _ := api.Rest().Get(grouphistoryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
