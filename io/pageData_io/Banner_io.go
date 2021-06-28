package pageData_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/pageData"
)

const bannerURL = api.BASE_URL + "banner/"

func CreateBanner(P pageData.Banner) (pageData.Banner, error) {

	entity := pageData.Banner{}
	resp, _ := api.Rest().SetBody(P).Post(bannerURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBanner(P pageData.Banner) (pageData.Banner, error) {

	entity := pageData.Banner{}
	resp, _ := api.Rest().SetBody(P).Post(bannerURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadBanner(id string) (pageData.Banner, error) {
	entity := pageData.Banner{}
	resp, _ := api.Rest().Get(bannerURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBannerN(id string) (pageData.BannerImageHelper, error) {
	entity := pageData.BannerImageHelper{}
	resp, _ := api.Rest().Get(bannerURL + "readN?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteBanner(id string) (pageData.Banner, error) {

	entity := pageData.Banner{}
	resp, _ := api.Rest().Get(bannerURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadBanners() ([]pageData.Banner, error) {

	entity := []pageData.Banner{}
	resp, _ := api.Rest().Get(bannerURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
