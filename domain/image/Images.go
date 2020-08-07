package image

type Images struct {
	Id          string `json:"id"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}

type ImagesHelper struct {
	Id       string `json:"id"`
	ImageId  string `json:"image"`
	BridgeId string `json:"description"`
}
