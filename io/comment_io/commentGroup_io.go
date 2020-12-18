package comment_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/comment"
)

const commentgroupURL = api.BASE_URL + "comment-group/"

func CreateCommentGroup(commentObject comment.CommentGroup) (comment.CommentGroup, error) {
	entity := comment.CommentGroup{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentgroupURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCommentGroup(commentObject comment.CommentGroup) (comment.CommentGroup, error) {
	entity := comment.CommentGroup{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentgroupURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentGroup(id string) (comment.CommentGroup, error) {
	entity := comment.CommentGroup{}
	resp, _ := api.Rest().Get(commentgroupURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CountCommentGroup(eventId string) (int64, error) {
	var entity int64
	resp, _ := api.Rest().Get(commentgroupURL + "count?eventId=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCommentGroup(id string) (comment.CommentGroup, error) {
	entity := comment.CommentGroup{}
	resp, _ := api.Rest().Get(commentgroupURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentGroups() ([]comment.CommentGroup, error) {
	entity := []comment.CommentGroup{}
	resp, _ := api.Rest().Get(commentgroupURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByGroupId(eventId string) ([]comment.CommentGroup, error) {
	entity := []comment.CommentGroup{}
	resp, _ := api.Rest().Get(commentgroupURL + "readAllByEventId?eventId=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllbyGroupId(projectId string) ([]comment.CommentGroup, error) {
	entity := []comment.CommentGroup{}
	resp, _ := api.Rest().Get(commentgroupURL + "readAllByCommentId")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
