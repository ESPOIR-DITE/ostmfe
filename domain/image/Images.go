package image

type Images struct {
	Id          string `json:"id"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
type ImagesHelper struct {
	Id          string `json:"id"`
	Image       string `json:"image"`
	Description string `json:"description"`
	BridgeId    string `json:"bridgeId"`
}
type Gallery struct {
	Id          string `json:"id"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
type GalleryHelper struct {
	Id          string `json:"id"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Bridge      string `json:"Bridge"`
}
type ImageType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
