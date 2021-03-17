package misc

import (
	"fmt"
	"ostmfe/domain/comment"
	"ostmfe/domain/history"
	"ostmfe/domain/image"
	"ostmfe/io/comment_io"
	"ostmfe/io/history_io"
)

func GetHistoryComments(historyId string) []comment.CommentHelper2 {
	var commentList []comment.CommentHelper2
	peopleComments, err := comment_io.ReadAllByHistoryId(historyId)
	if err != nil {
		fmt.Println(err, " error reading all the eventContribution")
	} else {
		for _, peopleComment := range peopleComments {
			commentObject, err := comment_io.ReadComment(peopleComment.CommentId)
			if err != nil {
				fmt.Println(err, " error reading all the Contribution")
			}
			commentObject2 := comment.CommentHelper2{commentObject.Id, commentObject.Email, commentObject.Name, FormatDateMonth(commentObject.Date), ConvertingToString(commentObject.Comment), getParentDeatils(commentObject.ParentCommentId), peopleComment.Id}
			commentList = append(commentList, commentObject2)
		}
	}
	return commentList
}

type HistoryGalleryImages struct {
	Gallery        image.GaleryHelper
	HistoryGallery history.HistoryGalery
}

func GetHistoryGallery(historyId string) []HistoryGalleryImages {
	var GalleryImagesList []HistoryGalleryImages
	peopleGalleries, err := history_io.ReadAllHistoryGallery(historyId)
	if err != nil {
		fmt.Println(err, "error reading People Gallery")
		return GalleryImagesList
	}
	for _, peopleGallerie := range peopleGalleries {
		GalleryImagesList = append(GalleryImagesList, HistoryGalleryImages{GetGalleryImage(peopleGallerie.GaleryId), peopleGallerie})
	}
	return GalleryImagesList
}
