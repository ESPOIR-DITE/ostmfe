package histories

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
)

type HistorySimpleData struct {
	History      history.HistoryHelper
	Images       []image3.ImagesHelper
	ProfileImage image3.Images
	Histories    history.HistoriesHelper
}

func GetHistorySimpleData(historyId string) HistorySimpleData {
	var historySimpleData HistorySimpleData
	var myHistories history.HistoriesHelper
	var profileImage image3.Images
	var images []image3.ImagesHelper

	ourHistory, err := history_io.ReadHistory(historyId)
	if err != nil {
		fmt.Println(err, " error reading History")
		return historySimpleData
	}

	//reading the histories
	historyHistories, err := history_io.ReadHistoryHistoriesWithHistoryId(historyId)
	if err != nil {
		fmt.Println(err, " error reading HistoryHistories")
	} else {
		histories, err := history_io.ReadHistorie(historyHistories.HistoriesId)
		if err != nil {
			fmt.Println(err, " error reading Histories")
		} else {
			myHistories = history.HistoriesHelper{histories.Id, misc.ConvertingToString(histories.History)}
		}
	}

	//Images
	historyImages, err := history_io.ReadHistoryImagesWithHistoryId(historyId)
	if err != nil {
		fmt.Println(err, " error reading HistoryImages")
	} else {
		for _, historyImage := range historyImages {
			if historyImage.Description == "1" || historyImage.Description == "profile" {
				profileImage, err = image_io.ReadImage(historyImage.ImageId)
				if err != nil {
					fmt.Println(err, " error reading ProfileImages")
				}
			}
			ImageObejct, err := image_io.ReadImage(historyImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading ProfileImages")
			}
			imageObject := image3.ImagesHelper{ImageObejct.Id, misc.ConvertingToString(ImageObejct.Image), historyImage.Id}
			images = append(images, imageObject)
		}
	}
	historyObject := history.HistoryHelper{ourHistory.Id, ourHistory.Title, ourHistory.Description, misc.FormatDateMonth(ourHistory.Date)}
	historySimpleData = HistorySimpleData{historyObject, images, profileImage, myHistories}
	return historySimpleData
}
