package group

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/comment"
	"ostmfe/domain/group"
	history2 "ostmfe/domain/history"
	image2 "ostmfe/domain/image"
	partner2 "ostmfe/domain/partner"
	project2 "ostmfe/domain/project"
	"ostmfe/io/comment_io"
	"ostmfe/io/group_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/project_io"
)

type GroupData struct {
	Group   group.Groupes
	History history2.HistoriesHelper
	Profile image2.Images
	Images  []image2.ImagesHelper
	Partner []partner2.Partner
	Project []project2.Project
}

func GetGroupData(groupId string) GroupData {
	var groupData GroupData
	var profileImage image2.Images
	var imageList []image2.ImagesHelper
	var projectList []project2.Project
	var partnerList []partner2.Partner
	var historyObejct history2.HistoriesHelper
	//Checking the group
	group, err := group_io.ReadGroup(groupId)
	if err != nil {
		fmt.Println(err, " error reading group")
		return groupData
	}
	//image
	groupImages, err := group_io.ReadGroupImagesWithGroupId(group.Id)
	if err != nil {
		fmt.Println(err, " error reading groupImage")
	} else {
		for _, groupImage := range groupImages {
			if groupImage.Description == "1" || groupImage.Description == "profile" {
				profileImage, err = image_io.ReadImage(groupImage.ImageId)
				if err != nil {
					fmt.Println(err, " error reading groupImage")
				}
			}
			imageObject, err := image_io.ReadImage(groupImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading groupImage")
			}
			imageObejectHelper := image2.ImagesHelper{imageObject.Id, misc.ConvertingToString(imageObject.Image), groupImage.Id}
			imageList = append(imageList, imageObejectHelper)
		}
	}
	//project
	groupProjects, err := group_io.ReadGroupProjectWithGroupId(groupId)
	if err != nil {
		fmt.Println(err, " error reading groupProject")
	} else {
		for _, groupProject := range groupProjects {
			project, err := project_io.ReadProject(groupProject.ProjectId)
			if err != nil {
				fmt.Println(err, " error reading project")
			}
			projectList = append(projectList, project)
		}
	}
	//partners
	groupPartners, err := group_io.ReadGroupPartnerWithGroupId(groupId)
	if err != nil {
		fmt.Println(err, " error reading groupPartner")
	} else {
		for _, groupPartner := range groupPartners {
			partners, err := partner_io.ReadPartner(groupPartner.PartenerId)
			if err != nil {
				fmt.Println(err, " error reading partner")
			}
			partnerList = append(partnerList, partners)
		}
	}
	//History
	groupHistory, err := group_io.ReadGroupHistoryWithGroupId(groupId)
	if err != nil {
		fmt.Println(err, " error reading groupHistory")
	} else {
		history, err := history_io.ReadHistorie(groupHistory.HistoryId)
		if err != nil {
			fmt.Println(err, " error reading History")
		}
		historyObejct = history2.HistoriesHelper{history.Id, misc.ConvertingToString(history.History)}
	}

	groupData = GroupData{group, historyObejct, profileImage, imageList, partnerList, projectList}

	return groupData
}

func GetGroupCommentsWithEventId(eventId string) []comment.CommentHelper2 {
	var commentList []comment.CommentHelper2
	eventComments, err := comment_io.ReadAllByGroupId(eventId)
	if err != nil {
		fmt.Println(err, " error reading all the eventContribution")
	} else {
		for _, eventComment := range eventComments {
			commentObject, err := comment_io.ReadComment(eventComment.CommentId)
			if err != nil {
				fmt.Println(err, " error reading all the Contribution")
			}
			commentObject2 := comment.CommentHelper2{commentObject.Id, commentObject.Email, commentObject.Name, misc.FormatDateMonth(commentObject.Date), misc.ConvertingToString(commentObject.Comment), getParentDeatils(commentObject.ParentCommentId), eventComment.Id}
			commentList = append(commentList, commentObject2)
		}
	}
	return commentList
}
func getParentDeatils(commentId string) comment.CommentHelper {
	commentObject, err := comment_io.ReadComment(commentId)
	if err != nil {
		fmt.Println(err, " error reading all the Contribution")
		return comment.CommentHelper{}
	}
	return comment.CommentHelper{commentObject.Id, commentObject.Email, commentObject.Name, misc.FormatDateMonth(commentObject.Date), misc.ConvertingToString(commentObject.Comment), commentObject.ParentCommentId}
}
