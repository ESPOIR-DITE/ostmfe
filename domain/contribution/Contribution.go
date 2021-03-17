package contribution

import (
	"mime/multipart"
	"time"
)

type Contribution struct {
	Id          string    `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	PhoneNumber string    `json:"phoneNumber"`
	Description []byte    `json:"description"`
}

type ContributionHelper struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	PhoneNumber string `json:"phoneNumber"`
	Description string `json:"description"`
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
	FileType       string `json:"fileType"`
	Description    string `json:"description"`
}

type ContributionFileTest struct {
	Id             string         `json:"id"`
	ContributionId string         `json:"contributionId"`
	File           multipart.File `json:"file"`
	FileType       string         `json:"fileType"`
	Description    string         `json:"description"`
}

type ContributionFileHelper struct {
	Id             string `json:"id"`
	ContributionId string `json:"contributionId"`
	File           string `json:"file"`
	FileType       string `json:"fileType"`
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
type ContributionHistory struct {
	Id             string `json:"id"`
	ContributionId string `json:"contributionId"`
	HistoryId      string `json:"historyId"`
	Description    string `json:"description"`
}
type EventPageFlow struct {
	Id       string `json:"id"`
	EventId  string `json:"eventId"`
	Title    string `json:"title"`
	PageFlow string `json:"pageFlow"`
}
