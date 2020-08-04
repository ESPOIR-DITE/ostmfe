package people

import (
	"fmt"
	"ostmfe/controller/misc"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/people"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/people_io"
)

/****
This object will be used to display a person on the person page.
*/
type PeopleBriefData struct {
	People people.People
	Image  image3.Images
}

/**
This method returns all the people's data only for those that have pictures.
*/
func GetPeopleBriefData() []PeopleBriefData {
	var peopleBriefDatas []PeopleBriefData
	var image image3.Images

	peoples, err := people_io.ReadPeoples()
	if err != nil {
		fmt.Println(err, " couldn't read people")
		return peopleBriefDatas
	}
	for _, people := range peoples {
		peopleImage, err := people_io.ReadPeopleImageWithPeopleId(people.Id)
		if err != nil {
			fmt.Println(err, " couldn't read peopleImage")
		}
		image, err = image_io.ReadImage(peopleImage.ImageId)
		if err != nil {
			fmt.Println(err, " couldn't read image")
		} else {
			peopleBriefData := PeopleBriefData{people, image}
			peopleBriefDatas = append(peopleBriefDatas, peopleBriefData)
		}
	}
	return peopleBriefDatas
}

//This struct make Up a people data with history
type PeopleDataHistory struct {
	People       people.People
	ProfileImage image3.Images
	Images       []image3.Images
	History      history2.HistoriesHelper
}

/****
This is the function used to display the content of a people in his page on the client side.
*/
func GetPeopleDataHistory(id string) PeopleDataHistory {
	var peopleDataHistory PeopleDataHistory
	var profileImage image3.Images
	var images []image3.Images
	//People
	people, err := people_io.ReadPeople(id)
	if err != nil {
		fmt.Println("could not read people")
		return peopleDataHistory
	}
	//Images
	peopleImages, err := people_io.ReadPeopleImagewithPeopleId(id)
	if err != nil {
		fmt.Println("could not read people Image")
		return peopleDataHistory
	}
	for _, peopleImage := range peopleImages {
		if peopleImage.ImageType == "profile" || peopleImage.ImageType == "1" {
			profileImage, err = image_io.ReadImage(peopleImage.ImageId)
			if err != nil {
				fmt.Println("could not read profile Image")
				//return peopleDataHistory;
			}
		}
		image, err := image_io.ReadImage(peopleImage.ImageId)
		if err != nil {
			fmt.Println("could not read Image")
			//return peopleDataHistory;
		}
		images = append(images, image)
	}
	//HistoryId
	peopleHistory, err := people_io.ReadPeopleHistoryWithPplId(id)
	if err != nil {
		fmt.Println("could not read people history")
		return peopleDataHistory
	}
	history, err := history_io.ReadHistorie(peopleHistory.HistoryId)
	if err != nil {
		fmt.Println("could not read history")
		return peopleDataHistory
	}
	historyhelper := history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}

	peopleDataHistory = PeopleDataHistory{people, profileImage, images, historyhelper}

	return peopleDataHistory

}
