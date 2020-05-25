package user

type Users struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
type UserAccount struct {
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
