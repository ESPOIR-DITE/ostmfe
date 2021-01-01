package comment

type Comment struct {
	Id              string `json:"id"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Date            string `json:"date"`
	Comment         []byte `json:"comment"`
	ParentCommentId string `json:"parentCommentId"`
}

type CommentHelper struct {
	Id              string `json:"id"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Date            string `json:"date"`
	Comment         string `json:"comment"`
	ParentCommentId string `json:"parentCommentId"`
}
type CommentHelper2 struct {
	Id              string `json:"id"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Date            string `json:"date"`
	Comment         string `json:"comment"`
	ParentCommentId CommentHelper
	BridgeId        string
}
type CommentStack struct {
	ParentComment CommentHelper
	SubComment    []CommentHelper
}
type CommentEvent struct {
	Id        string `json:"id"`
	EventId   string `json:"eventId"`
	CommentId string `json:"commentId"`
}
type CommentPlace struct {
	Id        string `json:"id"`
	PlaceId   string `json:"placeId"`
	CommentId string `json:"commentId"`
}
type CommentGroup struct {
	Id        string `json:"id"`
	GroupId   string `json:"groupId"`
	CommentId string `json:"commentId"`
}
type CommentHistory struct {
	Id        string `json:"id"`
	HistoryId string `json:"historyId"`
	CommentId string `json:"commentId"`
}
type CommentPeople struct {
	Id        string `json:"id"`
	PeopleId  string `json:"peopleId"`
	CommentId string `json:"commentId"`
}
type CommentProject struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	CommentId string `json:"commentId"`
}
