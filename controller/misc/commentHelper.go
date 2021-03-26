package misc

import (
	"fmt"
	"ostmfe/domain/comment"
	"ostmfe/io/comment_io"
)

//Calling this method makes the comment to be visible on the front page.
func ActivateComment(commentId string) bool {
	commentObject, err := comment_io.ReadComment(commentId)
	if err != nil {
		fmt.Print(err, " error reading comment")
		return false
	}
	_, err = comment_io.UpdateComment(comment.Comment{commentObject.Id, commentObject.Email, commentObject.Name, commentObject.Date, commentObject.Comment, commentObject.ParentCommentId, true})
	if err != nil {
		fmt.Print(err, " error reading comment")
		return false
	}
	return true
}

func DeleteComment(commentId string) bool {
	commentObject, err := comment_io.ReadComment(commentId)
	if err != nil {
		fmt.Print(err, " error reading comment")
		return false
	}
	_, err = comment_io.DeleteComment(commentObject.Id)
	if err != nil {
		fmt.Print(err, " error reading comment")
		return false
	}
	return true
}
