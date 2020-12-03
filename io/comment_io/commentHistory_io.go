package comment_io

import (
	"errors"
	"ostmfe/api"
	"ostmfe/domain/comment"
)

const commenthistoryURL = api.BASE_URL + "comment-history/"

func CreateCommentHistory(commentObject comment.CommentHistory) (comment.CommentHistory, error) {
	entity := comment.CommentHistory{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commenthistoryURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateCommentHistory(commentObject comment.CommentHistory) (comment.CommentHistory, error) {
	entity := comment.CommentHistory{}
	resp, _ := api.Rest().SetBody(commentObject).Post(commenthistoryURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentHistory(id string) (comment.CommentHistory, error) {
	entity := comment.CommentHistory{}
	resp, _ := api.Rest().Get(commenthistoryURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CountCommentHistory(eventId string) (int64, error) {
	var entity int64
	resp, _ := api.Rest().Get(commenthistoryURL + "count?eventId=" + eventId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteCommentHistory(id string) (comment.CommentHistory, error) {
	entity := comment.CommentHistory{}
	resp, _ := api.Rest().Get(commenthistoryURL + "delete?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadCommentHistorys() ([]comment.CommentHistory, error) {
	entity := []comment.CommentHistory{}
	resp, _ := api.Rest().Get(commenthistoryURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllByHistoryId(historyId string) ([]comment.CommentHistory, error) {
	entity := []comment.CommentHistory{}
	resp, _ := api.Rest().Get(commenthistoryURL + "readAllByHistoryId?historyId=" + historyId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadAllbyHistoryId(projectId string) ([]comment.CommentHistory, error) {
	entity := []comment.CommentHistory{}
	resp, _ := api.Rest().Get(commenthistoryURL + "readAllByCommentId")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
