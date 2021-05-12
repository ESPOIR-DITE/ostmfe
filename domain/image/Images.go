package image

type Images struct {
	Id          string `json:"id"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
type ImagesHelper struct {
	Id       string `json:"id"`
	Image    string `json:"image"`
	BridgeId string `json:"description"`
}
type Gallery struct {
	Id          string `json:"id"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
type GaleryHelper struct {
	Id          string `json:"id"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
type ImageType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
