package comment_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/comment"
)

const commentpeopleURL = api.BASE_URL + "comment-people/"

func CreateCommentPeople(commentObject comment.CommentPeople) (comment.CommentPeople, error) {
	entity := comment.CommentPeople{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentpeopleURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCommentPeople(commentObject comment.CommentPeople) (comment.CommentPeople, error) {
	entity := comment.CommentPeople{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentpeopleURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentPeople(id string) (comment.CommentPeople, error) {
	entity := comment.CommentPeople{}
	resp, _ := api.Rest().Get(commentpeopleURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CountCommentPeople(eventId string) (int64, error) {
	var entity int64
	resp, _ := api.Rest().Get(commenteventURL + "count?eventId=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCommentPeople(id string) (comment.CommentPeople, error) {
	entity := comment.CommentPeople{}
	resp, _ := api.Rest().Get(commentpeopleURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentPeoples() ([]comment.CommentPeople, error) {
	entity := []comment.CommentPeople{}
	resp, _ := api.Rest().Get(commentpeopleURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByPeopleId(eventId string) ([]comment.CommentPeople, error) {
	entity := []comment.CommentPeople{}
	resp, _ := api.Rest().Get(commentpeopleURL + "readAllByPeopleId?peopleId=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllbyPeopleId(projectId string) ([]comment.CommentPeople, error) {
	entity := []comment.CommentPeople{}
	resp, _ := api.Rest().Get(commentpeopleURL + "readAllByCommentId")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
