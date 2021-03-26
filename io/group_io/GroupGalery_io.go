package group_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/group"
)

const groupgaleryURL = api.BASE_URL + "group-galery/"

func CreateGroupGalery(myEvent group.GroupGalery) (group.GroupGalery, error) {
	entity := group.GroupGalery{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupgaleryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateGroupGalery(myEvent group.GroupGalery) (group.GroupGalery, error) {
	entity := group.GroupGalery{}
	resp, _ := api.Rest().SetBody(myEvent).Post(groupgaleryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadGroupGalery(id string) (group.GroupGalery, error) {
	entity := group.GroupGalery{}
	resp, _ := api.Rest().Get(groupgaleryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CountGroupGalery(id string) (group.GroupGalery, error) {
	entity := group.GroupGalery{}
	resp, _ := api.Rest().Get(groupgaleryURL + "countByGroupId?groupId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadGroupGaleryWithGroupId(id string) (group.GroupGalery, error) {
	entity := group.GroupGalery{}
	resp, _ := api.Rest().Get(groupgaleryURL + "readByGroupId?groupId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByGroupGalleryId(id string) ([]group.GroupGalery, error) {
	entity := []group.GroupGalery{}
	resp, _ := api.Rest().Get(groupgaleryURL + "readAllByGroupId?groupId=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteGroupGalery(id string) (group.GroupGalery, error) {
	entity := group.GroupGalery{}
	resp, _ := api.Rest().Get(groupgaleryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadGroupGalerys() ([]group.GroupGalery, error) {
	entity := []group.GroupGalery{}
	resp, _ := api.Rest().Get(groupgaleryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
