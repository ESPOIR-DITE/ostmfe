package partner_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/partner"
)

const partnerURL = api.BASE_URL + "partner/"

func CreatePartner(P partner.Partner) (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().SetBody(P).Post(partnerURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdatePartner(P partner.Partner) (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().SetBody(P).Post(partnerURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPartner(id string) (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().Get(partnerURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func DeletePartner(id string) (partner.Partner, error) {

	entity := partner.Partner{}
	resp, _ := api.Rest().Get(partnerURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadPartners() ([]partner.Partner, error) {

	entity := []partner.Partner{}
	resp, _ := api.Rest().Get(partnerURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
