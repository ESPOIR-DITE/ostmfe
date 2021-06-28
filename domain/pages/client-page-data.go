package pages

import (
	"ostmfe/domain/comment"
	"ostmfe/domain/event"
	"ostmfe/domain/group"
	"ostmfe/domain/history"
	"ostmfe/domain/image"
	"ostmfe/domain/pageData"
	"ostmfe/domain/people"
	"ostmfe/domain/place"
	"ostmfe/domain/project"
	"ostmfe/domain/slider"
	"ostmfe/domain/user"
)

type ClientLandingPageData struct {
	Sliders            []slider.Slider              `json:"sliders"`
	ProjectImageHelper []project.ProjectImageHelper `json:"projectImageHelper"`
	EventImageHelper   []event.EventImageHelper     `json:"eventImageHelper"`
	Projects           []project.Project            `json:"projects"`
	Histories          []history.History            `json:"histories"`
	Places             []place.Place                `json:"places"`
	ReadPageSection    []pageData.ReadPageSection   `json:"readPageSectionList"`
	PeoplePresentation []people.PeoplePresentation  `json:"peoplePresentationList"`
}
type AboutUsPageData struct {
	BannerImageHelper pageData.BannerImageHelper `json:"banner"`
	PageSection       []pageData.ReadPageSection `json:"pageSection"`
	StaffImageHelper  []user.StaffImageHelper    `json:"staffImageHelper"`
	GroupImageHelper  []group.GroupImageHelper   `json:"groupImageHelper"`
	GroupGallery      []image.GalleryHelper      `json:"groupGallery"`
}
type EventPageData struct {
	BannerImageHelper pageData.Banner          `json:"banner"`
	EventImageHelder  []event.EventImageHelper `json:"eventImageHelperList"`
	EventYearHelper   []event.EventYearHelper  `json:"eventYearHelperList"`
	Project           []project.Project        `json:"projectList"`
}
type PeoplePageData struct {
	People            people.People           `json:"people"`
	ProfileImage      image.ImagesHelper      `json:"profileImage"`
	DescriptiveImage  []image.ImagesHelper    `json:"descriptiveImage"`
	HistoriesHelper   history.HistoriesHelper `json:"historiesHelper"`
	PlaceList         []place.Place           `json:"placeList"`
	Events            []event.Event           `json:"events"`
	Gallery           []image.GalleryHelper   `json:"galleryHelperList"`
	BannerImageHelper pageData.Banner         `json:"banner"`
	Comments          []comment.CommentHelper `json:"comments"`
	NumberOfComments  int64                   `json:"numberOfComments"`
}

type PeopleHomePage struct {
	Categories        []people.Category          `json:"categories"`
	PeopleImageHelper []people.PeopleImageHelper `json:"peopleImageHelpers"`
	BannerImageHelper pageData.Banner            `json:"banner"`
	PageSection       []pageData.ReadPageSection `json:"pageSection"`
}
