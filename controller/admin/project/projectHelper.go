package project

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/io/comment_io"
)

func GetProjectCommentsWithProjectId(eventId string) []comment.CommentHelper2 {
	var commentList []comment.CommentHelper2
	projectComments, err := comment_io.ReadAllByProjectId(eventId)
	if err != nil {
		fmt.Println(err, " error reading all the eventContribution")
	} else {
		for _, projectComment := range projectComments {
			commentObject, err := comment_io.ReadComment(projectComment.CommentId)
			if err != nil {
				fmt.Println(err, " error reading all the Contribution")
			}
			fmt.Println("commentObject: ", commentObject)
			commentObject2 := comment.CommentHelper2{commentObject.Id, commentObject.Email, commentObject.Name, misc.FormatDateMonth(commentObject.Date), misc.ConvertingToString(commentObject.Comment), getParentDeatils(commentObject.ParentCommentId), projectComment.Id}
			commentList = append(commentList, commentObject2)
		}
	}
	return commentList
}
func getParentDeatils(commentId string) comment.CommentHelper {
	commentObject, err := comment_io.ReadComment(commentId)
	if err != nil {
		fmt.Println(err, " error reading all the Contribution")
		return comment.CommentHelper{}
	}
	return comment.CommentHelper{commentObject.Id, commentObject.Email, commentObject.Name, misc.FormatDateMonth(commentObject.Date), misc.ConvertingToString(commentObject.Comment), commentObject.ParentCommentId, commentObject.Stat}
}
