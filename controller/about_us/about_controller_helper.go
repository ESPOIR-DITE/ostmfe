package about_us

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/io/group_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
)

type GroupDataHistory struct {
	Group        group.Groups
	ProfileImage image3.Images
	Images       []image3.Images
	History      history2.HistoriesHelper
}

func GetGroupDataHistory(groupId string) GroupDataHistory {
	var groupDataHistory GroupDataHistory
	var profile image3.Images
	var images []image3.Images

	//Group
	group, err := group_io.ReadGroup(groupId)
	if err != nil {
		fmt.Println("could not read groups")
		return groupDataHistory
	}
	//images
	groupImages, err := group_io.ReadGroupImagesWithGroupId(groupId)
	if err != nil {
		fmt.Println("could not read group Image")
		return groupDataHistory
	}
	for _, groupImage := range groupImages {
		if groupImage.Description == "profile" || groupImage.Description == "1" {
			profile, err = image_io.ReadImage(groupImage.ImageId)
			if err != nil {
				fmt.Println("could not read profile Image")
				return groupDataHistory
			}
		}
		image, err := image_io.ReadImage(groupImage.ImageId)
		if err != nil {
			fmt.Println("could not read Image")

		}
		images = append(images, image)
	}
	groupHistory, err := group_io.ReadGroupHistoryWithGroupId(groupId)
	if err != nil {
		fmt.Println("could not read group history")
		return groupDataHistory
	}
	history, err := history_io.ReadHistorie(groupHistory.HistoryId)
	if err != nil {
		fmt.Println("could not read history")
		return groupDataHistory
	}
	historyhelper := history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}

	groupDataHistory = GroupDataHistory{group, profile, images, historyhelper}

	return groupDataHistory
}
