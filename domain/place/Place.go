package place

//todo each place should have image and each image should also have a description.
type Place struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
}
type PlaceImage struct {
	Id          string `json:"id"`
	PlaceId     string `json:"placeId"`
	ImageId     string `json:"imageId"`
	Description string `json:"description"`
}
type PlaceHistory struct {
	Id        string `json:"id"`
	PlaceId   string `json:"placeId"`
	HistoryId string `json:"historyId"`
}
type PlaceImageHelper struct {
	PlaceImage PlaceImage `json:"placeImage"`
	Files      [][]byte   `json:"files"`
}
