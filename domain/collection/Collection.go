package collection

type Collection struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	ProfileDescription string `json:"profile_description"`
	History            []byte `json:"history"`
}

type Collection_image struct {
	Id           string `json:"id"`
	ImageId      string `json:"image_id"`
	CollectionId string `json:"collectionId"`
	Description  string `json:"description"`
}
type CollectionTypes struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Collection_type struct {
	Id             string `json:"Id"`
	CollectionId   string `json:"collectionId"`
	CollectionType string `json:"collectionType"`
}
type CollectionImageHelper struct {
	Collection_image Collection_image `json:"collection_image"`
	Files            [][]byte         `json:"files"`
}
