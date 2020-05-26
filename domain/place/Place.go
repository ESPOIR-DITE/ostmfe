package place

//todo each place should have image and each image should also have a description.
type Place struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Description string `json:"description"`
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
