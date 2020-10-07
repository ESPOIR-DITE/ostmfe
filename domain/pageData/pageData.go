package pageData

type PageData struct {
	Id          string `json:"id"`
	PageName    string `json:"pageName"`
	Description string `json:"description"`
}
type PageSection struct {
	Id        string `json:"id"`
	PageId    string `json:"pageId"`
	SectionId string `json:"sectionId"`
	Content   []byte `json:"content"`
}
type PageSectionHelper struct {
	Id          string `json:"id"`
	PageId      string `json:"pageId"`
	SectionId   string `json:"sectionId"`
	SectionName string `json:"sectionName"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
type SectionBlock struct {
	Id          string `json:"id"`
	SectionName string `json:"sectionName"`
	Description string `json:"description"`
}
