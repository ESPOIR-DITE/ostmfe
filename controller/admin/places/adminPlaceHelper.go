package places

import (
	"fmt"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	place2 "ostmfe/domain/place"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
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
