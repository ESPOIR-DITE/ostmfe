package comment_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/comment"
)

const commentplaceURL = api.BASE_URL + "comment-place/"

func CreateCommentPlace(commentObject comment.CommentEvent) (comment.CommentEvent, error) {
	entity := comment.CommentEvent{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentplaceURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCommentPlace(commentObject comment.CommentEvent) (comment.CommentEvent, error) {
	entity := comment.CommentEvent{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentplaceURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentPlace(id string) (comment.CommentEvent, error) {
	entity := comment.CommentEvent{}
	resp, _ := api.Rest().Get(commentplaceURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CountCommentPlace(eventId string) (int64, error) {
	var entity int64
	resp, _ := api.Rest().Get(commentplaceURL + "count?eventId=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCommentPlace(id string) (comment.CommentEvent, error) {
	entity := comment.CommentEvent{}
	resp, _ := api.Rest().Get(commentplaceURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentPlaces() ([]comment.CommentEvent, error) {
	entity := []comment.CommentEvent{}
	resp, _ := api.Rest().Get(commentplaceURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByPlaceId(placeId string) ([]comment.CommentEvent, error) {
	entity := []comment.CommentEvent{}
	resp, _ := api.Rest().Get(commentplaceURL + "readAllByPlaceId?eventId=" + placeId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllbyPlaceId(projectId string) ([]comment.CommentEvent, error) {
	entity := []comment.CommentEvent{}
	resp, _ := api.Rest().Get(commentplaceURL + "readAllByCommentId")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
