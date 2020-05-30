package people

import "time"

type People struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	BirthDate  time.Time `json:"birth_date"`
	DeathDate  time.Time `json:"deathdate"`
	Origine    string    `json:"origine"`
	Profession string    `json:"profession"`
}

type People_image struct {
	Id        string `json:"id"`
	PeopleId  string `json:"people_id"`
	ImageId   string `json:"image_id"`
	ImageType string `json:"image_type"`
}

type People_profession struct {
	Profession  string `json:"profession"`
	People_id   string `json:"people_id"`
	Description string `json:"description"`
}
type PlaceImageHelper struct {
	People_image People_image `json:"people_image"`
	Files        [][]byte     `json:"files"`
}

type PeoplePlace struct {
	Id       string `json:"id"`
	PlaceId  string `json:"place_id"`
	PeopleId string `json:"people_id"`
}

type Profession struct {
	Id          string `json:"id"`
	Profession  string `json:"profession"`
	Description string `json:"description"`
}

type Profession_image struct {
	ProfessionId string `json:"professionId"`
	ImageId      string `json:"image_id"`
	Description  string `json:"description"`
}
type PeopleHistory struct {
	Id        string `json:"id"`
	PeopleId  string `json:"people_id"`
	HistoryId string `json:"history_id"`
}
type PeopleCategory struct {
	Id       string `json:"id"`
	Category string `json:"category"`
}
