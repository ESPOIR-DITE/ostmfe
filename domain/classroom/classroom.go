package classroom

type Classroom struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Details     []byte `json:"details"`
	Icon        []byte `json:"icon"`
}

type ClassroomHelper struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Details     string `json:"details"`
	Icon        string `json:"icon"`
}
