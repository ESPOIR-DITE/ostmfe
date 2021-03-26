package misc

import (
	"fmt"
	"ostmfe/domain/comment"
	"ostmfe/domain/image"
	"ostmfe/domain/people"
	"ostmfe/io/comment_io"
	"ostmfe/io/people_io"
)

func GetPeopleComments(peopleId string) []comment.CommentHelper2 {
	var commentList []comment.CommentHelper2
	peopleComments, err := comment_io.ReadAllByPeopleId(peopleId)
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

type PeopleGalleryImages struct {
	Gallery       image.GaleryHelper
	PeopleGallery people.PeopleGalery
}

func GetPeopleGallery(placeId string) []PeopleGalleryImages {
	var GalleryImagesList []PeopleGalleryImages
	peopleGaleries, err := people_io.ReadAllByPeopleIdGalery(placeId)
	if err != nil {
		fmt.Println(err, "error reading People Gallery")
		return GalleryImagesList
	}
	for _, peopleGallerie := range peopleGaleries {
		GalleryImagesList = append(GalleryImagesList, PeopleGalleryImages{GetGalleryImage(peopleGallerie.Galery), peopleGallerie})
	}
	return GalleryImagesList
}
