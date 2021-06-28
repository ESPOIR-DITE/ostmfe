package project

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/io/comment_io"
)

func GetProjectCommentsWithProjectId(projectId string) []comment.CommentHelper2 {
	var commentList []comment.CommentHelper2
	projectComments, err := comment_io.ReadAllByProjectId(projectId)
	if err != nil {
		fmt.Println(err, " error reading all the eventContribution")
	} else {
		for _, projectComment := range projectComments {
			commentObject, err := comment_io.ReadComment(projectComment.CommentId)
			if err != nil {
				fmt.Println(err, " error reading all the Contribution")
			}
			fmt.Println("commentObject: ", commentObject)
			commentObject2 := comment.CommentHelper2{commentObject.Id, commentObject.Email, commentObject.Name, misc.FormatDateMonth(commentObject.Date), misc.ConvertingToString(commentObject.Comment), getParentDeatils(commentObject.ParentCommentId, projectComment.Id), projectComment.Id}
			commentList = append(commentList, commentObject2)
		}
	}
	return commentList
}
func getParentDeatils(commentId, bridgeId string) comment.CommentHelper {
	commentObject, err := comment_io.ReadComment(commentId)
	if err != nil {
		fmt.Println(err, " error reading all the Contribution")
		return comment.CommentHelper{}
	}
	return comment.CommentHelper{commentObject.Id, commentObject.Email, commentObject.Name, misc.FormatDateMonth(commentObject.Date), misc.ConvertingToString(commentObject.Comment), commentObject.ParentCommentId, commentObject.Stat, bridgeId}
}

func projectCommentCalculation(projectId string) (commentNumber int64, pending int64, active int64) {
	var commentNumbers int64 = 0
	var pendings int64 = 0
	var actives int64 = 0
	historyComments, err := comment_io.ReadAllByProjectId(projectId)
	if err != nil {
		fmt.Println(err, " error reading Project comment")
		return commentNumbers, pendings, actives
	} else {
		for _, historyComment := range historyComments {
			comments, err := comment_io.ReadComment(historyComment.CommentId)
			if err != nil {
				fmt.Println(err, " error reading comment")
			} else {
				if comments.Stat == true {
					actives++
				} else {
					pending++
				}
				commentNumber++
			}
		}
	}
	return commentNumbers, pendings, actives
}
