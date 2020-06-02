package project

type Project struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProjectMember struct {
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

type ProjectPartner struct {
	ProjectId   string `json:"projectId"`
	PartenerID  string `json:"partenerId"`
	Description string `json:"description"`
}
type ProjectImageHelper struct {
	Files        [][]byte     `json:"files"`
	ProjectImage ProjectImage `json:"projectImage"`
}
type ProjectHistory struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	HistoryId string `json:"historyId"`
}
