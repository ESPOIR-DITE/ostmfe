package places

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	"ostmfe/io/comment_io"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
)

type PlaceImageHelperEditable struct {
	Id           string
	ImageId      string
	PlaceImageId string
}
type PlaceDataEditable struct {
	Place        place2.Place
	Images       []PlaceImageHelperEditable
	ProfileImage PlaceImageHelperEditable
	History      history2.HistoriesHelper
}

func GetPlaceEditable(placeId string) PlaceDataEditable {
	var place place2.Place
	var profileImage PlaceImageHelperEditable
	var images []PlaceImageHelperEditable
	var historyhelper history2.HistoriesHelper

	var placeDataEditable PlaceDataEditable

	place, err := place_io.ReadPlace(placeId)
	if err != nil {
		fmt.Println(err, " error reading place")
		return placeDataEditable
	}

	//Images
	placeImages, err := place_io.ReadPlaceImageAllOf(placeId)
	if err != nil {
		fmt.Println(err, " error reading place Image. This place may not have images")
	} else {
		fmt.Println("looping PlaceImages ", placeImages)
		for _, placeImage := range placeImages {
			if placeImage.Description == "1" || placeImage.Description == "profile" {
				profImage, err := image_io.ReadImage(placeImage.ImageId)
				if err != nil {
					fmt.Println(err, " error reading Image")
				}
				profileImage = PlaceImageHelperEditable{profImage.Id, misc.ConvertingToString(profImage.Image), placeImage.Id}
			}
			image, err := image_io.ReadImage(placeImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading Image")
			}
			imageObject := PlaceImageHelperEditable{image.Id, misc.ConvertingToString(image.Image), placeImage.Id}
			images = append(images, imageObject)
		} //end looping placeImages
	} //end reading Place image

	//History
	placeHistory, err := place_io.ReadPlaceHistporyOf(placeId)
	if err != nil {
		fmt.Println(err, " error reading placeHistory. This place may not have a History yet")
	} else {
		history, err := history_io.ReadHistorie(placeHistory.HistoryId)
		if err != nil {
			fmt.Println(err, " error reading History")
		} else {
			historyhelper = history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}
		}

	}
	placeDataEditable = PlaceDataEditable{place, images, profileImage, historyhelper}
	return placeDataEditable
}

//PlaceComment and Place Gallery.

func GetPlaceCommentsWithEventId(placeId string) []comment.CommentHelper2 {
	var commentList []comment.CommentHelper2
	placeComments, err := comment_io.ReadAllByPlaceId(placeId)
	if err != nil {
		fmt.Println(err, " error reading all the eventContribution")
	} else {
		for _, placeComment := range placeComments {
			commentObject, err := comment_io.ReadComment(placeComment.CommentId)
			if err != nil {
				fmt.Println(err, " error reading all the Contribution")
			}
			commentObject2 := comment.CommentHelper2{commentObject.Id, commentObject.Email, commentObject.Name, misc.FormatDateMonth(commentObject.Date), misc.ConvertingToString(commentObject.Comment), getParentDeatils(commentObject.ParentCommentId), placeComment.Id}
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

//Getting all the PlacePeople Of a particular place.
func GetAllPeoplePlace(placeId string) []people.People {
	var peopleList []people.People
	PeoplePlaces, err := people_io.ReadPeoplePlaceAllByPlaceId(placeId)
	if err != nil {
		fmt.Println(err, " error reading all PeoplePlace")
		return peopleList
	}
	for _, PeoplePlace := range PeoplePlaces {
		PeopleObejct, err := people_io.ReadPeople(PeoplePlace.PeopleId)
		if err != nil {
			fmt.Println(err, " error reading all People")
		} else {
			//We are putting the Id of the PeoplePlace into people id field so that we can be able to delete peoplePlace
			peopleList = append(peopleList, people.People{PeoplePlace.Id, PeopleObejct.Name, PeopleObejct.Surname, misc.FormatDateMonth(PeopleObejct.BirthDate), PeopleObejct.DeathDate, PeopleObejct.Origin, PeopleObejct.Profession, PeopleObejct.Brief})
		}
	}
	return peopleList
}

//Get All the event Of a Place
func GetEventPlace(placeId string) []event.Event {
	var events []event.Event
	var eventIds []string
	eventsPlace, err := event_io.ReadEventFindByPlaceId(placeId)
	if err != nil {
		fmt.Println(err, " error reading all eventPlaces")
		return events
	}
	for _, eventPlace := range eventsPlace {
		eventIds = append(eventIds, eventPlace.EventId)
	}
	return misc.GetEventListOfEventIdList(eventIds)
}
