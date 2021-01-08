package misc

import (
	"fmt"
	"ostmfe/domain/event"
	"ostmfe/domain/image"
	"ostmfe/io/event_io"
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
