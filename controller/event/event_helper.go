package event

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
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
	event, err := event_io.ReadEvent(eventId)
	if err != nil {
		fmt.Println(err, " error reading event")
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
			fmt.Println(err, " error reading event HistoryId")
		} else {
			histor, err := history_io.ReadHistorie(eventHistory.HistoryId)
			if err != nil {
				fmt.Println(err, " error reading  HistoryId")
			} else {
				history = history2.HistoriesHelper{histor.Id, misc.ConvertingToString(histor.History)}
			}
		}
		eventData = EventData{event, profileImage, images, history}
	}
	return eventData
}
