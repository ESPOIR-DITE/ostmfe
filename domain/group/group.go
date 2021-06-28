package group

import image3 "ostmfe/domain/image"

type Groupes struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HistoryId   string `json:"historyId"`
}
type GroupGalery struct {
	Id       string `json:"id"`
	GroupId  string `json:"groupId"`
	GaleryId string `json:"galeryId"`
}

type GroupHistory struct {
	Id        string `json:"id"`
	GroupId   string `json:"groupId"`
	HistoryId string `json:"historyId"`
}
type GroupImage struct {
	Id          string `json:"id"`
	ImageId     string `json:"imageId"`
	GroupId     string `json:"groupId"`
	ImageTypeId string `json:"imageTypeId"`
	Description string `json:"description"`
}
type GroupImageHelper struct {
	Groupes Groupes       `json:"groupes"`
	Images  image3.Images `json:"images"`
}
type GroupMember struct {
	Id       string `json:"id"`
	MemberId string `json:"userId"`
	Date     string `json:"date"`
	GroupId  string `json:"groupId"`
}
type GroupPartener struct {
	Id          string `json:"id"`
	PartenerId  string `json:"partenerId"`
	GroupId     string `json:"groupId"`
	Description string `json:"description"`
}
type GroupProject struct {
	Id          string `json:"id"`
	ProjectId   string `json:"projectId"`
	GroupId     string `json:"groupId"`
	Description string `json:"description"`
}
type GroupMenberHelper struct {
	Groupes Groupes `json:"groupes"`
	Members int     `json:"members"`
}
