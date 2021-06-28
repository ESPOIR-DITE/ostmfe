package peoples

import (
	"fmt"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	"ostmfe/domain/people"
	"ostmfe/io/comment_io"
	"ostmfe/io/event_io"
	"ostmfe/io/people_io"
)

/****
This struct will return all the picture and HistoryId objects of a person
*/
//type PeopleEditable struct {
//	People  people.People
//	Images  PeopleImageHelper
//	History history2.HistoriesHelper
//}

type EventPeopleData struct {
	Id        string
	EventName string
}
type NewPeopleEditable struct {
	People            people.People
	ProfileImage      image.Images
	DescriptiveImages []image.Images
	History           history2.HistoriesHelper
}

func NewGetPeopleEditable(peopleId string) NewPeopleEditable {
	var peopleEditable NewPeopleEditable
	var historiesHelper history2.HistoriesHelper

	people, err := people_io.ReadPeople(peopleId)
	if err != nil {
		fmt.Println(err, " can not read this people")
		return peopleEditable
	}
	profileImage, err := people_io.ReadPeopleProfileImage(peopleId)
	if err != nil {
		fmt.Println(err, " can not read this people profile Image")
		//return peopleEditable
	}
	descriptiveImage, err := people_io.ReadPeopleDescriptiveImage(peopleId)
	if err != nil {
		fmt.Println(err, " can not read this people descriptive Image")
	}
	history, err := people_io.ReadPeopleHistories(peopleId)
	if err != nil {
		fmt.Println(err, " can not read this people History")
	} else {
		historiesHelper = history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}
	}
	return NewPeopleEditable{people, profileImage, descriptiveImage, historiesHelper}
}

// Deprecated
//func GetPeopleEditable(peopleId string) PeopleEditable {
//	var Images PeopleImageHelper
//	var historyToreturn history2.HistoriesHelper
//	var peopleEditable PeopleEditable
//	var deathDate string
//
//	peopleObject, err := people_io.ReadPeople(peopleId)
//	if err != nil {
//		fmt.Println(err, " can not read this people")
//		return peopleEditable
//	}
//	if misc.FormatDateMonth(peopleObject.BirthDate) == misc.FormatDateMonth(peopleObject.DeathDate) {
//		deathDate = "living"
//	} else {
//		deathDate = misc.FormatDateMonth(peopleObject.DeathDate)
//	}
//	peopleToReturn := people.People{peopleObject.Id, peopleObject.Name, peopleObject.Surname, misc.FormatDateMonth(peopleObject.BirthDate), deathDate, peopleObject.Origin, peopleObject.Profession, peopleObject.Brief}
//	//Reading Image
//	fmt.Println("peopleId", peopleId)
//	peopleImages, err := people_io.ReadPeopleImageWithPeopleId(peopleId)
//	if err != nil {
//		fmt.Println(err, "  error reading peopleImages")
//	} else {
//		ImageObejct, err := image_io.ReadImage(peopleImages.ImageId)
//		if err != nil {
//			fmt.Println(err, " error read Image")
//		}
//		//I am replacing description with peopleImageId to facilitate image update process.
//		imageObject := PeopleImageHelper{ImageObejct.Id, misc.ConvertingToString(ImageObejct.Image), peopleImages.Id}
//		Images = imageObject
//	}
//	//HistoryId
//	peopleHistory, err := people_io.ReadPeopleHistoryWithPplId(peopleId)
//	if err != nil {
//		fmt.Println(err, "  error reading peopleHistory")
//	} else {
//		history, err := history_io.ReadHistorie(peopleHistory.HistoryId)
//		if err != nil {
//			fmt.Println(err, " error read history of id: ", peopleHistory.HistoryId)
//		}
//		historyToreturn = history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}
//	}
//	peopleEditable = PeopleEditable{peopleToReturn, Images, historyToreturn}
//	return peopleEditable
//}

//With peopleId, you get the commentNumber, pending, active.
func peopleCommentCalculation(peopleIs string) (commentNumber int64, pending int64, active int64) {
	var commentNumbers int64 = 0
	var pendings int64 = 0
	var actives int64 = 0
	peopleComments, err := comment_io.ReadAllCommentPeopleWithPeopleId(peopleIs)
	if err != nil {
		fmt.Println(err, " error reading People comment")
		return commentNumbers, pendings, actives
	} else {
		for _, peopleComment := range peopleComments {
			comments, err := comment_io.ReadComment(peopleComment.CommentId)
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

func GetPeopleEvents(peopleId string) []EventPeopleData {
	var eventPeopleObject []EventPeopleData

	peopleEvents, err := event_io.ReadEventPeopleWithPeopleId(peopleId)
	if err != nil {
		fmt.Println(err, " error reading event people")
		return eventPeopleObject
	}
	for _, peopleEvent := range peopleEvents {
		event, err := event_io.ReadEvent(peopleEvent.EventId)
		if err != nil {
			fmt.Println(err, " error reading")
		}
		eventPeopleObject = append(eventPeopleObject, EventPeopleData{peopleEvent.Id, event.Name})
	}
	return eventPeopleObject
}

func GetPeopleCategory(peopleId string) (people.Category, error) {
	peopleCategory, err := people_io.ReadPeopleCategoryWithPplId(peopleId)
	if err != nil {
		fmt.Println(err, " error reading people category")
		return people.Category{}, err
	}
	category, err := people_io.ReadCategory(peopleCategory.Category)
	if err != nil {
		fmt.Println(err, " error reading category")
		return people.Category{}, err
	}
	return category, nil
}
func getPeopleObjectReady(peopleObject people.People) people.People {
	var deathDate string
	if misc.FormatDateMonth(peopleObject.BirthDate) == misc.FormatDateMonth(peopleObject.DeathDate) {
		deathDate = "living"
	} else {
		deathDate = misc.FormatDateMonth(peopleObject.DeathDate)
	}
	return people.People{peopleObject.Id,
		peopleObject.Name,
		peopleObject.Surname,
		misc.FormatDateMonth(peopleObject.BirthDate),
		deathDate,
		peopleObject.Origin,
		peopleObject.Profession,
		peopleObject.Brief,
		peopleObject.HistoriesId}
}
