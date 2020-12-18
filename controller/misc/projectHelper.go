package misc

import (
	"fmt"
	"ostmfe/domain/comment"
	"ostmfe/domain/image"
	"ostmfe/domain/project"
	"ostmfe/io/comment_io"
	"ostmfe/io/project_io"
)

type ProjectGalleryImages struct {
	Gallery        image.GaleryHelper
	ProjectGallery project.ProjectGallery
}

func GetProjectGallery(projectId string) []ProjectGalleryImages {
	var GalleryImagesList []ProjectGalleryImages

	projectGalleryImages, err := project_io.ReadAllProjectGalleryWithProjectId(projectId)
	if err != nil {
		fmt.Println(err, "error reading groupImage")
		return GalleryImagesList
	}
	for _, projectGalleryImage := range projectGalleryImages {
		GalleryImagesList = append(GalleryImagesList, ProjectGalleryImages{GetGalleryImage(projectGalleryImage.GalleryId), projectGalleryImage})
	}
	return GalleryImagesList
}

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
			commentObject2 := comment.CommentHelper2{commentObject.Id, commentObject.Email, commentObject.Name, FormatDateMonth(commentObject.Date), ConvertingToString(commentObject.Comment), getParentDeatils(commentObject.ParentCommentId), projectComment.Id}
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
	return comment.CommentHelper{commentObject.Id, commentObject.Email, commentObject.Name, FormatDateMonth(commentObject.Date), ConvertingToString(commentObject.Comment), commentObject.ParentCommentId}
}
