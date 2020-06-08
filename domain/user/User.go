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
type Roles struct {
	Id          string `json:"id"`
	Role        string `json:"role"`
	Description string `json:"description"`
}
type UserRole struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	RoleId string `json:"roleId"`
}
