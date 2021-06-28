package pageData

type PageData struct {
	Id          string `json:"id"`
	PageName    string `json:"pageName"`
	BannerId    string `json:"bannerId"`
	Description string `json:"description"`
}
type PageSection struct {
	Id        string `json:"id"`
	PageId    string `json:"pageId"`
	SectionId string `json:"sectionId"`
	//PageName    string `json:"pageName"`
	//Description string `json:"description"`
	Content []byte `json:"content"`
}
type ReadPageSectionHelper struct {
	SectionName string `json:"sectionName"`
	Content     string `json:"content"`
}
type ReadPageSection struct {
	SectionName string `json:"sectionName"`
	Content     string `json:"content"`
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
type PageBanner struct {
	Id          string `json:"id"`
	PageName    string `json:"pageName"`
	Description string `json:"description"`
	BannerId    string `json:"bannerId"`
}
type Banner struct {
	Id    string `json:"id"`
	Image []byte `json:"image"`
}
type BannerImageHelper struct {
	Id    string `json:"id"`
	Image string `json:"image"`
}
