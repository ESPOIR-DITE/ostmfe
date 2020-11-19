package comment

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/io/comment_io"
)

func GetProjectComment(projectId string) []comment.CommentStack {
	var parentCommentStack []comment.CommentStack
	var parentComment []comment.CommentHelper
	var subComment []comment.CommentHelper

	for _, commentObject := range getProjectComment(projectId) {
		if commentObject.ParentCommentId != "" {
			subComment = append(subComment, commentObject)
		} else {
			parentComment = append(parentComment, commentObject)
			parentCommentStack = append(parentCommentStack, comment.CommentStack{commentObject, subComment})
		}
	}
	return parentCommentStack
}

func getProjectComment(projectId string) []comment.CommentHelper {
	var myCommentObject []comment.CommentHelper
	projectComments, err := comment_io.ReadAllByProjectId(projectId)
	if err != nil {
		fmt.Println("error reading event Comment")
		return myCommentObject
	}
	for _, projectComment := range projectComments {
		myComment, err := comment_io.ReadComment(projectComment.CommentId)
		if err != nil {
			fmt.Println("error reading Comment")
		}
		if myComment.ParentCommentId != "" && myComment.Comment != nil {
			commentHelper := comment.CommentHelper{myComment.Id, myComment.Email, myComment.Name, misc.FormatDateMonth(myComment.Date), misc.ConvertingToString(myComment.Comment), myComment.ParentCommentId}
			myCommentObject = append(myCommentObject, commentHelper)
		}
	}
	return myCommentObject
}

func GetEventComments(eventId string) []comment.CommentStack {
	var parentCommentStack []comment.CommentStack
	var subCommentStack []comment.CommentHelper

	for _, commentObject := range getComments(eventId) {
		myComment, err := comment_io.ReadComment(commentObject.Id)
		if err != nil {
			fmt.Println("error reading Comment")
		}
		if myComment.ParentCommentId != "" && myComment.Comment != nil {
			subCommentStack = getSubComment(commentObject.Id)
		}
		parentCommentStack = append(parentCommentStack, comment.CommentStack{commentObject, subCommentStack})
	}

	fmt.Println("parentStack ", parentCommentStack)

	return parentCommentStack
}

func getSubComment(parentComment string) []comment.CommentHelper {
	var myComments []comment.CommentHelper
	subComments, err := comment_io.ReadCommentWithParentId(parentComment)
	if err != nil {
		return myComments
	}
	for _, eventComment := range subComments {
		if eventComment.ParentCommentId == parentComment && eventComment.Comment != nil {
			commentHelper := comment.CommentHelper{eventComment.Id, eventComment.Email, eventComment.Name, misc.FormatDateMonth(eventComment.Date), misc.ConvertingToString(eventComment.Comment), eventComment.ParentCommentId}
			myComments = append(myComments, commentHelper)
		}
	}
	return myComments
}

//This method returns a list of either parent or subcomment
func getComments(eventId string) []comment.CommentHelper {
	var myCommentObject []comment.CommentHelper
	eventComments, err := comment_io.ReadAllByEventId(eventId)
	if err != nil {
		fmt.Println("error reading event Comment")
		return myCommentObject
	}
	for _, eventComment := range eventComments {
		myComment, err := comment_io.ReadComment(eventComment.CommentId)
		if err != nil {
			fmt.Println("error reading Comment")
		} else {
			commentHelper := comment.CommentHelper{myComment.Id, myComment.Email, myComment.Name, misc.FormatDateMonth(myComment.Date), misc.ConvertingToString(myComment.Comment), myComment.ParentCommentId}
			myCommentObject = append(myCommentObject, commentHelper)
		}
	}
	return myCommentObject
}

func checkIfIsBelongingToParent(commentId string, comment2 []comment.Comment) (bool, comment.Comment) {
	var commentObject comment.Comment
	if comment2 == nil {
		return false, commentObject
	} else {
		for _, commentValeu := range comment2 {
			if commentId == commentValeu.ParentCommentId {
				return true, commentValeu
			}
		}
		return false, commentObject
	}
}
