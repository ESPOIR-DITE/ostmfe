package peoples

import (
	"fmt"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	"ostmfe/domain/people"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/people_io"
)

/****
This struct will return all the picture and History objects of a person
*/
type PeopleEditable struct {
	People  people.People
	Images  []PeopleImageHelperEditable
	History history2.HistoriesHelper
}
type PeopleImageHelperEditable struct {
	Id            string
	ImageId       string
	PeopleImageId string
}

func GetPeopleEditable(peopleId string) PeopleEditable {
	var Images []PeopleImageHelperEditable
	var historyToreturn history2.HistoriesHelper
	var peopleEditable PeopleEditable

	people, err := people_io.ReadPeople(peopleId)
	if err != nil {
		fmt.Println(err, " can not read this people")
		return peopleEditable
	}
	//Reading Image
	peopleImages, err := people_io.ReadPeopleImagewithPeopleId(peopleId)
	if err != nil {
		fmt.Println(err, "  error reading peopleImages")
	} else {
		for _, peopleImage := range peopleImages {
			ImageObejct, err := image_io.ReadImage(peopleImage.ImageId)
			if err != nil {
				fmt.Println(err, " error read Image")
			}
			//I am replacing description with peopleImageId to facilitate image update process.
			imageObject := PeopleImageHelperEditable{ImageObejct.Id, misc.ConvertingToString(ImageObejct.Image), peopleImage.Id}
			Images = append(Images, imageObject)
		}
	}
	//History
	peopleHistory, err := people_io.ReadPeopleHistoryWithPplId(peopleId)
	if err != nil {
		fmt.Println(err, "  error reading peopleHistory")
	} else {
		history, err := history_io.ReadHistorie(peopleHistory.HistoryId)
		if err != nil {
			fmt.Println(err, " error read history of id: ", peopleHistory.HistoryId)
		}
		historyToreturn = history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}
	}
	peopleEditable = PeopleEditable{people, Images, historyToreturn}
	return peopleEditable
}
