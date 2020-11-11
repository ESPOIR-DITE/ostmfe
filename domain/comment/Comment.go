package comment

import "time"

type Comment struct {
	Id              string    `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Date            time.Time `json:"date"`
	Comment         []byte    `json:"comment"`
	ParentCommentId string    `json:"parentCommentId"`
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
