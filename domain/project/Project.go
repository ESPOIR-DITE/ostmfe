package project

import "ostmfe/domain/image"

type Project struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProjectMember struct {
	Id          string `json:"id"`
	ProjectId   string `json:"project_id"`
	MemberId    string `json:"member_id"`
	Description string `json:"description"`
}
type ProjectImage struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	ImageId   string `json:"imageId"`
	ImageType string `json:"imageType"`
}
type ProjectPageFlow struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	ProjectId  string `json:"projectId"`
	PageFlowId string `json:"pageFLowId"`
}

type ProjectPartner struct {
	Id          string `json:"id"`
	ProjectId   string `json:"projectId"`
	PartenerID  string `json:"partenerId"`
	Description string `json:"description"`
}
type ProjectImageHelper struct {
	Image   image.Images `json:"images"`
	Project Project      `json:"project"`
}
type ProjectHistory struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	HistoryId string `json:"historyId"`
}
type ProjectGallery struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	GalleryId string `json:"galleryId"`
}
