package comment_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/comment"
)

const commentProjectURL = api.BASE_URL + "comment-project/"

func CreateCommentProject(commentObject comment.CommentProject) (comment.CommentProject, error) {
	entity := comment.CommentProject{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentProjectURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCommentProject(commentObject comment.CommentProject) (comment.CommentProject, error) {
	entity := comment.CommentProject{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commentProjectURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentProject(id string) (comment.CommentProject, error) {
	entity := comment.CommentProject{}
	resp, _ := api.Rest().Get(commentProjectURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCommentProject(id string) (comment.CommentProject, error) {
	entity := comment.CommentProject{}
	resp, _ := api.Rest().Get(commentProjectURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadCommentProjects() ([]comment.CommentProject, error) {
	entity := []comment.CommentProject{}
	resp, _ := api.Rest().Get(commentProjectURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllByProjectId(projectId string) ([]comment.CommentProject, error) {
	entity := []comment.CommentProject{}
	resp, _ := api.Rest().Get(commentProjectURL + "readAllbyProjectId?projectId=" + projectId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllbyCommentId(projectId string) ([]comment.CommentProject, error) {
	entity := []comment.CommentProject{}
	resp, _ := api.Rest().Get(commentProjectURL + "readAllbyCommentId")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CountProjectComment(projectId string) (int64, error) {
	var entity int64
	resp, _ := api.Rest().Get(commentProjectURL + "count?projectId=" + projectId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
