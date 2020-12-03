package contribution

type Contribution struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	PhoneNumber string `json:"phoneNumber"`
	Description []byte `json:"description"`
}
type ContributionHelper struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	PhoneNumber string `json:"phoneNumber"`
	Description string `json:"description"`
	Image       string
}
type ContributionEvent struct {
	Id             string `json:"id"`
	ContributionId string `json:"contributionId"`
	EventId        string `json:"eventId"`
	Description    string `json:"description"`
}
type ContributionFile struct {
	Id             string `json:"id"`
	ContributionId string `json:"contributionId"`
	File           []byte `json:"file"`
	Description    string `json:"description"`
}
type ContributionFileType struct {
	Id       string `json:"id"`
	FileType string `json:"fileType"`
}
type ContributionProject struct {
	Id             string `json:"id"`
	ProjectId      string `json:"projectId"`
	ContributionId string `json:"contributionId"`
	Description    string `json:"description"`
}
