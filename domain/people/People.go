package people

import (
	"ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	"ostmfe/domain/place"
)

type People struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	BirthDate  string `json:"birth_date"`
	DeathDate  string `json:"deathdate"`
	Origin     string `json:"origin"`
	Profession string `json:"profession"`
	Brief      string `json:"brief"`
}

type PeopleImage struct {
	Id        string `json:"id"`
	PeopleId  string `json:"peopleId"`
	ImageId   string `json:"imageId"`
	ImageType string `json:"imageTypeId"`
}

type People_profession struct {
	Profession  string `json:"profession"`
	People_id   string `json:"people_id"`
	Description string `json:"description"`
}
type PeopleImageHelper struct {
	People            People                   `json:"people"`
	ProfileImage      image.Images             `json:"profileImage"`
	DescriptiveImages []image.Images           `json:"descriptiveImages"`
	History           history2.HistoriesHelper `json:"history"`
}

type PeoplePlace struct {
	Id       string `json:"id"`
	PlaceId  string `json:"placeId"`
	PeopleId string `json:"peopleId"`
}

type Profession struct {
	Id          string `json:"id"`
	Profession  string `json:"profession"`
	Description string `json:"description"`
}

type Profession_image struct {
	ProfessionId string `json:"professionId"`
	ImageId      string `json:"imageId"`
	Description  string `json:"description"`
}
type PeopleHistory struct {
	Id        string `json:"id"`
	PeopleId  string `json:"peopleId"`
	HistoryId string `json:"historyId"`
}
type PeopleCategory struct {
	Id          string `json:"id"`
	Category    string `json:"categoryId"`
	PeopleId    string `json:"peopleId"`
	Description string `json:"description"`
}
type PeopleGalery struct {
	Id       string `json:"id"`
	PeopleId string `json:"peopleId"`
	Galery   string `json:"galery"`
}

type Category struct {
	Id       string `json:"id"`
	Category string `json:"category"`
}

type PeopleAggregate struct {
	People           People                   `json:"people"`
	Category         Category                 `json:"category"`
	ProfileImage     image.ImagesHelper       `json:"profileImage"`
	History          history2.HistoriesHelper `json:"histories"`
	Gallery          []image.Gallery          `json:"gallery"`
	DescriptionImage []image.ImagesHelper     `json:"descriptiveImage"`
	Profession       []Profession             `json:"profession"`
	Places           []place.Place            `json:"places"`
	Events           []event.Event            `json:"events"`
}
