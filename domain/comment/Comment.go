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
type CommentProject struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	CommentId string `json:"commentId"`
}
