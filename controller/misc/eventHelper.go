package misc

import (
	"fmt"
	"ostmfe/domain/event"
	"ostmfe/domain/image"
	"ostmfe/io/event_io"
	"time"
)

type EventGalleryImages struct {
	Gallery      image.GaleryHelper
	EventGallery event.EventGalery
}

func GetEventGallery(eventId string) []EventGalleryImages {
	var GalleryImagesList []EventGalleryImages

	groupGalleryImages, err := event_io.ReadAllEventGalleryWithEventId(eventId)
	if err != nil {
		fmt.Println(err, "error reading groupImage")
		return GalleryImagesList
	}
	for _, groupGalleryImage := range groupGalleryImages {
		GalleryImagesList = append(GalleryImagesList, EventGalleryImages{GetGalleryImage(groupGalleryImage.GaleryId), groupGalleryImage})
	}
	return GalleryImagesList
}

//flexible for any request that needs a list of events by sending a list of eventIds
func GetEventListOfEventIdList(eventIds []string) []event.Event {
	var events []event.Event
	for _, eventId := range eventIds {
		eventObject, err := event_io.ReadEvent(eventId)
		if err != nil {
			fmt.Println(err, " error reading event")
		}
		events = append(events, event.Event{eventObject.Id, eventObject.Name, FormatDateMonth(eventObject.Date), eventObject.IsPast, eventObject.Description})
	}
	return events
}

//This method help to get upcoming event
func UpComingEvents() ([]event.Event, error) {
	var myList []event.Event
	var myReserveList []event.Event
	events, err := event_io.ReadEvents()
	if err != nil {
		fmt.Println(err, " error reading events")
		return myList, err
	}
	for _, event := range events {
		DateString := FormatDateMonth(event.Date)
		dateI := ParseDateTime(DateString)
		nowYear, nowMonth, _ := time.Now().Date()
		if dateI.Month() == nowMonth && nowYear == dateI.Year() {
			myList = append(myList, event)
			fmt.Println("this months events: ", event.Name)
		} else {
			myReserveList = append(myReserveList, event)
		}
	}
	myList = append(myReserveList)
	return myList, nil
}
