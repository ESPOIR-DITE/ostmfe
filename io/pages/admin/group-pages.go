package admin

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pages"
)

const adminGroupPage = api.BASE_URL + "admin-pages/group/"

func GetGroupAdminEditPageData(groupId string) (pages.GroupAdminEditPresentation, error) {
	entity := pages.GroupAdminEditPresentation{}
	resp, _ := api.Rest().Get(adminGroupPage + "read?groupId=" + groupId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
