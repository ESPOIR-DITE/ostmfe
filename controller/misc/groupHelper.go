package misc

import (
	"fmt"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/image"
	"ostmfe/io/group_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
)

//Returns all Histories
type HistoryAndProfile struct {
	History history2.History
	Profile string
	Date    string
}

type GalleryImages struct {
	Gallery      image.GaleryHelper
	GroupGallery group.GroupGalery
}

func GetGroupGallery(groupId string) []GalleryImages {
	var GalleryImagesList []GalleryImages

	groupGalleryImages, err := group_io.ReadAllByGroupGalleryId(groupId)
	if err != nil {
		fmt.Println(err, "error reading groupImage")
		return GalleryImagesList
	}
	for _, groupGalleryImage := range groupGalleryImages {
		GalleryImagesList = append(GalleryImagesList, GalleryImages{GetGalleryImage(groupGalleryImage.GaleryId), groupGalleryImage})
	}
	return GalleryImagesList
}

func ReadHistoryWithImages() []HistoryAndProfile {
	var historyAndProfile []HistoryAndProfile

	histories, err := history_io.ReadHistorys()
	if err != nil {
		fmt.Println(err, " error reading all the histories")
		return historyAndProfile
	}
	for _, history := range histories {
		imageHistory, err := history_io.ReadHistoryImageWithHistoryId(history.Id)
		if err != nil {
			fmt.Println(err, " error reading all the ImageHistory")
		} else {
			historyAndProfile = append(historyAndProfile, HistoryAndProfile{history, getImage(imageHistory.ImageId), FormatDateMonth(history.Date)})
		}
	}
	return historyAndProfile
}
func getImage(imageId string) string {
	imageObject, err := image_io.ReadImage(imageId)
	if err != nil {
		fmt.Println(err, " error reading image")
		return ""
	}
	return imageObject.Id
}
