package group

type Groupes struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
	Description string `json:"description"`
}
type GroupImageHelper struct {
	GroupImage GroupImage `json:"groupImage"`
	Files      [][]byte   `json:"files"`
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
