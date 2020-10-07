package event

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/event"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/people"
	place2 "ostmfe/domain/place"
	"ostmfe/io/event_io"
	"ostmfe/io/group_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
)

/****
This struct is responsible to have all the necessary objects to present an event on single-event page.
*/
type EventData struct {
	Event        event.Event
	ProfileImage image3.Images
	Images       []image3.Images
	History      history2.HistoriesHelper
}

func GetEventData(eventId string) EventData {
	var eventData EventData
	var profileImage image3.Images
	var images []image3.Images
	var history history2.HistoriesHelper
	eventobj, err := event_io.ReadEvent(eventId)
	if err != nil {
		fmt.Println(err, " error reading eventobj")
		return eventData
	}
	eventImages, err := event_io.ReadEventImgOf(eventId)
	if err != nil {
		fmt.Println(err, " error reading Event image")
		return eventData
	}
	for _, eventImage := range eventImages {
		if eventImage.Description == "1" || eventImage.Description == "profile" {
			profileImage, err = image_io.ReadImage(eventImage.ImageId)
		}
		image, err := image_io.ReadImage(eventImage.ImageId)
		if err != nil {
			fmt.Println(err, " error reading image")
		} else {
			images = append(images, image)
		}
		//HistoryId
		eventHistory, err := event_io.ReadEventHistoryWithEventId(eventId)
		if err != nil {
			fmt.Println(err, " error reading eventobj HistoryId")
		} else {
			histor, err := history_io.ReadHistorie(eventHistory.HistoryId)
			if err != nil {
				fmt.Println(err, " error reading  HistoryId")
			} else {
				history = history2.HistoriesHelper{histor.Id, misc.ConvertingToString(histor.History)}
			}
		}
		eventObject := event.Event{eventobj.Id, eventobj.Name, misc.FormatDateMonth(eventobj.Date), eventobj.IsPast, eventobj.Description}
		eventData = EventData{eventObject, profileImage, images, history}
	}
	return eventData
}

//EventPlace
func GetEnventPlaceData(eventId string) place2.Place {
	var place place2.Place
	eventPlace, err := event_io.ReadEventPlaceOf(eventId)
	if err != nil {
		fmt.Println(err, "error reading eventPlace")
		return place
	} else {
		place, err = place_io.ReadPlace(eventPlace.PlaceId)
		if err != nil {
			fmt.Println(err, "error reading Place")
		}
	}
	return place
}

//EventPeople
func GetEventPeopleData(eventId string) []people.People {
	var peoples []people.People
	var profileImage image3.Images

	eventPeoples, err := event_io.ReadEventPeopleOf(eventId)
	if err != nil {
		fmt.Println(err, " error reading eventPeople")
		return peoples
	} else {
		for _, eventPeople := range eventPeoples {
			peopleObej, err := people_io.ReadPeople(eventPeople.PeopleId)
			if err != nil {
				fmt.Println(err, " error reading People")
			} else {
				peopleImages, err := people_io.ReadPeopleImagewithPeopleId(peopleObej.Id)
				if err != nil {
					fmt.Println(err, " error reading PeopleImage")
				} else {
					for _, peopleImage := range peopleImages {
						if peopleImage.ImageType == "profile" || peopleImage.ImageType == "1" {
							profileImage, err = image_io.ReadImage(peopleImage.ImageId)
							if err != nil {
								fmt.Println("could not read profile Image")
							}
						}
					}
				}
				//I am adding the image in deathdate variale
				peopleObject := people.People{peopleObej.Id, peopleObej.Name, peopleObej.Surname, misc.FormatDateMonth(peopleObej.BirthDate), profileImage.Id, peopleObej.Origin, peopleObej.Profession, peopleObej.Brief}
				peoples = append(peoples, peopleObject)
			}
		}
	}
	return peoples
}

type GroupData struct {
	Group group.Groups
	Image image3.Images
}

func GetGroupsData(eventId string) []GroupData {
	var goupDatas []GroupData

	groupEvents, err := event_io.ReadEventGroupAllOfs(eventId)
	if err != nil {
		fmt.Println(err, " error reading EventGroups")
		return goupDatas
	}
	for _, groupEvent := range groupEvents {
		group, err := group_io.ReadGroup(groupEvent.GroupId)
		if err != nil {
			fmt.Println(err, " error reading groups")
			return goupDatas
		}
		groupImage, err := group_io.ReadGroupImageWithGroupId(group.Id)
		if err != nil {
			fmt.Println(err, " error reading groups image")
		} else {
			image, err := image_io.ReadImage(groupImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading groups image")
			} else {
				goupDatas = append(goupDatas, GroupData{group, image})
			}
		}
	}
	return goupDatas
}
