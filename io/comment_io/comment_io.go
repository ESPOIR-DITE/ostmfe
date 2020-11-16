package comment_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/comment"
)

const commentURL = api.BASE_URL + "comment/"

func CreateComment(commentObject comment.Comment) (comment.Comment, error) {
	entity := comment.Comment{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateComment(commentObject comment.Comment) (comment.Comment, error) {
	entity := comment.Comment{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadComment(id string) (comment.Comment, error) {
	entity := comment.Comment{}
	resp, _ := api.Rest().Get(commentURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteComment(id string) (comment.Comment, error) {
	entity := comment.Comment{}
	resp, _ := api.Rest().Get(commentURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadComments() ([]comment.Comment, error) {
	entity := []comment.Comment{}
	resp, _ := api.Rest().Get(commentURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
