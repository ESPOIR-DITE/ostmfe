package user

import (
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	"time"
)

type Users struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
type UserAccount struct {
	Email    string    `json:"email"`
	Date     time.Time `json:"date"`
	Password string    `json:"password"`
}
type Roles struct {
	Id          string `json:"id"`
	Role        string `json:"role"`
	Description string `json:"description"`
}
type RoleOfUser struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	RoleId string `json:"roleId"`
}
type UserImage struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	HistoryId   string `json:"historyId"`
	ImageId     string `json:"imageId"`
	Description string `json:"description"`
}
type UserImageHelper struct {
	Users  Users        `json:"users"`
	Images image.Images `json:"images"`
}
type StaffImageHelper struct {
	Users     Users                    `json:"users"`
	Images    image.Images             `json:"images"`
	Histories history2.HistoriesHelper `json:"histories"`
}
