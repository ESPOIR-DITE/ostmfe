package people_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/image"
	"ostmfe/domain/people"
)

const peopleImg = api.BASE_URL + "people_image/"

func CreatePeopleImage(pI people.PeopleImage) (people.PeopleImage, error) {

	entity := people.PeopleImage{}
	resp, _ := api.Rest().SetBody(pI).Post(peopleImg + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreatePeopleImageHere(pI people.PeopleImage) (people.PeopleImage, error) {

	entity := people.PeopleImage{}
	resp, _ := api.Rest().SetBody(pI).Post(peopleImg + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreatePeopleImageX(pI people.PeopleImage) (people.PeopleImage, error) {

	entity := people.PeopleImage{}
	resp, _ := api.Rest().SetBody(pI).Post(peopleImg + "createx")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePeopleImage(pI people.PeopleImage) (people.PeopleImage, error) {

	entity := people.PeopleImage{}
	resp, _ := api.Rest().SetBody(pI).Post(peopleImg + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPeopleImage(id string) (people.PeopleImage, error) {

	entity := people.PeopleImage{}
	resp, _ := api.Rest().Get(peopleImg + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPeopleImagewithPeopleId(id string) ([]people.PeopleImage, error) {
	entity := []people.PeopleImage{}
	resp, _ := api.Rest().Get(peopleImg + "read_people?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadPeopleImageWithPeopleId(id string) (people.PeopleImage, error) {
	entity := people.PeopleImage{}
	resp, _ := api.Rest().Get(peopleImg + "read_people_image?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePeopleImage(id string) (people.PeopleImage, error) {

	entity := people.PeopleImage{}
	resp, _ := api.Rest().Get(peopleImg + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPeopleImages() ([]people.PeopleImage, error) {
	entity := []people.PeopleImage{}
	resp, _ := api.Rest().Get(peopleImg + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPeopleProfileImage(peopleId string) (image.Images, error) {
	entity := image.Images{}
	resp, _ := api.Rest().Get(peopleImg + "getProfileImage?peopleId=" + peopleId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPeopleDescriptiveImage(peopleId string) ([]image.Images, error) {
	entity := []image.Images{}
	resp, _ := api.Rest().Get(peopleImg + "getDescriptiveImages?peopleId=" + peopleId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
