package member

type Member struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	InterestedIn string `json:"interestedIn"`
	Message      string `json:"message"`
	Address      string `json:"address"`
}
