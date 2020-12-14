package event

type Event struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	IsPast      string `json:"isPast"`
	Description string `json:"description"`
}

type EventImage struct {
	Id          string `json:"id"`
	ImageId     string `json:"imageId"`
	EventId     string `json:"eventId"`
	Description string `json:"description"`
}

type EventPartener struct {
	Id          string `json:"id"`
	PartenerId  string `json:"partenerId"`
	EventId     string `json:"eventId"`
	Description string `json:"description"`
}

type EventPlace struct {
	Id          string `json:"id"`
	PlaceId     string `json:"placeId"`
	EventId     string `json:"eventId"`
	Description string `json:"description"`
}

type EventProject struct {
	Id          string `json:"id"`
	ProjectId   string `json:"projectId"`
	EventId     string `json:"eventId"`
	Description string `json:"description"`
}
type EventHistory struct {
	Id        string `json:"id"`
	HistoryId string `json:"historyId"`
	EventId   string `json:"eventId"`
}
type EventImageHelper struct {
	EventImage EventImage `json:"eventImage"`
	Files      [][]byte   `json:"files"`
}
type EventPeople struct {
	Id       string `json:"id"`
	EventId  string `json:"eventId"`
	PeopleId string `json:"peopleId"`
}
type EventYear struct {
	Id      string `json:"id"`
	EventId string `json:"eventId"`
	YearId  string `json:"yearId"`
}
type EventGroup struct {
	Id      string `json:"id"`
	EventId string `json:"eventId"`
	GroupId string `json:"groupId"`
}
type EventGalery struct {
	Id       string `json:"id"`
	EventId  string `json:"eventId"`
	GaleryId string `json:"galeryId"`
}
