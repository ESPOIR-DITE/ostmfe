package history

type History struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type HistoryHelper struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type HistoryImage struct {
	Id          string `json:"id"`
	ImageId     string `json:"imageId"`
	HistoryId   string `json:"historyId"`
	Description string `json:"description"`
}

type HistoryHistories struct {
	Id          string `json:"id"`
	HistoryId   string `json:"historyId"`
	HistoriesId string `json:"historiesId"`
}

type HistoryImageHelper struct {
	HistoryImage HistoryImage `json:"historyImage"`
	Files        [][]byte     `json:"files"`
}

type HistoryPeople struct {
	Id        string `json:"id"`
	HistoryId string `json:"historyId"`
	PeopleId  string `json:"peopleId"`
}

type HistoryGalery struct {
	Id        string `json:"id"`
	HistoryId string `json:"historyId"`
	GaleryId  string `json:"galeryId"`
}

type Histories struct {
	Id      string `json:"id"`
	History []byte `json:"history"`
}

type HistoriesHelper struct {
	Id      string `json:"id"`
	History string `json:"history"`
}

type CategoryH struct {
	Id          string `json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
type HistoryCategory struct {
	Id        string `json:"id"`
	Category  string `json:"category"`
	HistoryId string `json:"historyId"`
}
