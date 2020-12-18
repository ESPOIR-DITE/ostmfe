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
