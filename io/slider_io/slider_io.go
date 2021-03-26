package slider_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/slider"
)

const sliderURL = api.BASE_URL + "slider/"

func CreateSlider(M slider.Slider) (slider.Slider, error) {

	entity := slider.Slider{}
	resp, _ := api.Rest().SetBody(M).Post(sliderURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateSlider(M slider.Slider) (slider.Slider, error) {
	entity := slider.Slider{}
	resp, _ := api.Rest().SetBody(M).Post(sliderURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadSlider(id string) (slider.Slider, error) {
	entity := slider.Slider{}
	resp, _ := api.Rest().Get(sliderURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteSlider(id string) (slider.Slider, error) {
	entity := slider.Slider{}
	resp, _ := api.Rest().Get(sliderURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadSliders() ([]slider.Slider, error) {
	entity := []slider.Slider{}
	resp, _ := api.Rest().Get(sliderURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
