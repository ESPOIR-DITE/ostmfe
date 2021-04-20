package slider

type Slider struct {
	Id          string `json:"id"`
	SliderName  string `json:"slideName"`
	Description string `json:"description"`
	SliderImage []byte `json:"sliderImage"`
}
type SliderHelper struct {
	Id          string `json:"id"`
	SliderName  string `json:"slideName"`
	Description string `json:"description"`
	SliderImage string `json:"sliderImage"`
}
