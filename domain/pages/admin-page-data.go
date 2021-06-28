package pages

import (
	museum "ostmfe/domain"
	"ostmfe/domain/comment"
	"ostmfe/domain/contribution"
	"ostmfe/domain/event"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	"ostmfe/domain/pageData"
	"ostmfe/domain/partner"
	"ostmfe/domain/people"
	"ostmfe/domain/place"
	"ostmfe/domain/project"
	"ostmfe/domain/user"
)

type AdminLandingPageData struct {
	Projects            int                       `json:"projects"`
	Events              int                       `json:"events"`
	Users               int                       `json:"users"`
	Comments            int                       `json:"comments"`
	ContributionEvent   int                       `json:"contributionEvent"`
	ContributionHistory int                       `json:"contributionHistory"`
	ContributionProject int                       `json:"contributionProject"`
	GroupMemberHelpers  []group.GroupMenberHelper `json:"groupMemberHelpers"`
	UserImageHelper     user.UserImageHelper      `json:"userImageHelper"`
}

type UserPageData struct {
	UserAndRole     []UserAndRole        `json:"userAndRole"`
	Role            []user.Roles         `json:"roles"`
	UserImageHelper user.UserImageHelper `json:"userImageHelper"`
}
type UserAndRole struct {
	Users user.Users `json:"users"`
	Role  user.Roles `json:"roles"`
}
type EventViewData struct {
	Profile          image.ImagesHelper           `json:"profileImage"`
	DescriptiveImage []image.ImagesHelper         `json:"descriptiveImage"`
	Event            event.Event                  `json:"event"`
	History          history2.HistoriesHelper     `json:"histories"`
	EventYear        museum.Years                 `json:"eventYear"`
	EventProject     project.Project              `json:"eventProject"`
	EventPlace       place.Place                  `json:"eventPlace"`
	EventPeoples     []people.People              `json:"eventPeoples"`
	EventPartners    []partner.Partner            `json:"eventPartners"`
	EventGroups      []group.Groupes              `json:"eventGroups"`
	Comments         []comment.CommentHelper      `json:"comments"`
	EventPageFlow    []contribution.EventPageFlow `json:"eventPageFlows"`
	GalleryHelper    []image.GalleryHelper        `json:"galleryHelpers"`
	Projects         []project.Project            `json:"projects"`
	Partners         []partner.Partner            `json:"partners"`
	People           []people.People              `json:"people"`
	Places           []place.Place                `json:"places"`
	Years            []museum.Years               `json:"years"`
	Groupes          []group.Groupes              `json:"groupes"`
}

type GroupAdminEditPresentation struct {
	Group             group.Groupes            `json:"groupes"`
	Histories         history2.HistoriesHelper `json:"histories"`
	DescriptiveImages []image.ImagesHelper     `json:"descriptiveImages"`
	ProfileImage      image.ImagesHelper       `json:"profileImage"`
	Partners          partner.Partner          `json:"partners"`
	Project           project.Project          `json:"projects"`
}

type GroupClientSingleData struct {
	Banner             pageData.BannerImageHelper  `json:"banner"`
	Group              group.Groupes               `json:"groupes"`
	ProfileImage       image.ImagesHelper          `json:"profileImage"`
	DescriptionImages  []image.GalleryHelper       `json:"descriptionImages"`
	History            history2.HistoriesHelper    `json:"historiesHelper"`
	PeoplePresentation []people.PeoplePresentation `json:"peoplePresentationList"`
	Gallery            []image.GalleryHelper       `json:"galleryHelpers"`
}
