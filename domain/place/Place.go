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
type PlaceHistories struct {
	Id        string `json:"id"`
	PlaceId   string `json:"placeId"`
	HistoryId string `json:"historyId"`
}

type PlaceGallery struct {
	Id        string `json:"id"`
	PlaceId   string `json:"placeId"`
	GalleryId string `json:"galleryId"`
}

type PlaceImageHelper struct {
	PlaceImage PlaceImage `json:"placeImage"`
	Files      [][]byte   `json:"files"`
}
type PlacePlace struct {
	Id          string `json:"id"`
	PlaceId     string `json:"placeId"`
	HistoryId   string `json:"historyId"`
	Description string `json:"description"`
}
