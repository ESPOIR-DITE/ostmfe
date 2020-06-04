package image

type Images struct {
	Id          string `json:"id"`
	Image       []byte `json:"image"`
	Description string `json:"description"`
}
