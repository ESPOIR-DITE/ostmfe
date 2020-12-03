package slider

type Slider struct {
	Id            string `json:"id"`
	SliderName    string `json:"slideName"`
	Description   string `json:"description"`
	SliderMessage []byte `json:"sliderMessage"`
	SliderImage   []byte `json:"sliderImage"`
}
type SliderHelper struct {
	Id            string `json:"id"`
	SliderName    string `json:"slideName"`
	Description   string `json:"description"`
	SliderMessage string `json:"sliderMessage"`
	SliderImage   string `json:"sliderImage"`
}
