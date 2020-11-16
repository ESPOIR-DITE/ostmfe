package comment_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/comment"
)

const commenteventURL = api.BASE_URL + "comment-event/"

func CreateCommentEvent(commentObject comment.CommentEvent) (comment.CommentEvent, error) {
	entity := comment.CommentEvent{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commenteventURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCommentEvent(commentObject comment.CommentEvent) (comment.CommentEvent, error) {
	entity := comment.CommentEvent{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commenteventURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentEvent(id string) (comment.CommentEvent, error) {
	entity := comment.CommentEvent{}
	resp, _ := api.Rest().Get(commenteventURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCommentEvent(id string) (comment.CommentEvent, error) {
	entity := comment.CommentEvent{}
	resp, _ := api.Rest().Get(commenteventURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentEvents() ([]comment.CommentEvent, error) {
	entity := []comment.CommentEvent{}
	resp, _ := api.Rest().Get(commenteventURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByEventId(projectId string) ([]comment.CommentEvent, error) {
	entity := []comment.CommentEvent{}
	resp, _ := api.Rest().Get(commenteventURL + "readAllByEventId")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllbyEventId(projectId string) ([]comment.CommentEvent, error) {
	entity := []comment.CommentEvent{}
	resp, _ := api.Rest().Get(commenteventURL + "readAllByCommentId")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
