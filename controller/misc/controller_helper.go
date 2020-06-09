package misc

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"ostmfe/domain/collection"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/member"
	"ostmfe/domain/partner"
	project2 "ostmfe/domain/project"
	user2 "ostmfe/domain/user"
	"ostmfe/io/collection_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/member_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/project_io"
	"ostmfe/io/user_io"
	"strings"
	"time"
)

const (
	YYYYMMDD_FORMAT    = "2006-01-02"
	YYYMMDDTIME_FORMAT = "2006-01-02 15:04:05"
)

/**
Format date in yyyy-MM-dd HH:mm:ss
*/

func FormatDateTime(date time.Time) string {
	return date.Format(YYYMMDDTIME_FORMAT)
}

/**
format date in yyyy-MM-dd
*/
func FormatDate(date time.Time) string {
	return date.Format(YYYYMMDD_FORMAT)
}

/***
this method should separates longitude and latitude
*/
func SeparateLatLng(latlng string) (string, string) {
	var longitude string
	var latitude string
	val := strings.TrimSuffix(latlng, ")")
	val2 := strings.TrimPrefix(val, "(")
	parts := strings.Split(val2, ",")
	if len(parts) != 2 {
		return latitude, longitude
	}
	latitude = parts[0]
	longitude = parts[1]
	return latitude, longitude
}

func CheckFiles(files []io.Reader) [][]byte {
	var bytelist [][]byte
	for index, reablefile := range files {
		if reablefile != nil {
			reader := bufio.NewReader(reablefile)
			content, _ := ioutil.ReadAll(reader)
			bytelist = append(bytelist, content)
			fmt.Println(" done with file: ", index)
		}
	}
	return bytelist
}

type ProjectContentsHome struct {
	ProjectId   string
	Title       string
	Picture     string
	Description string
}

func GetProjectContentsHomes() []ProjectContentsHome {
	projectContentsHomeObject := []ProjectContentsHome{}
	image := image3.Images{}
	projects, err := project_io.ReadProjects()
	if err != nil {
		fmt.Println(err, " Error reading all the projects")
		return projectContentsHomeObject
	}
	for _, project := range projects {
		//fmt.Println(project.Title)
		projectImage, err := project_io.ReadWithProjectIdProjectImage(project.Id)
		if err != nil {
			fmt.Println(err, " Can not find the following project in project image table: ", project.Title)
		} else {
			image, err = image_io.ReadImage(projectImage.ImageId)
			//fmt.Println(image.Image)
			if err != nil {
				fmt.Println(err, " Can not find the following project image Id in Image table: ", projectImage.ImageId)
			}
		}
		projectObject := ProjectContentsHome{project.Id, project.Title, image.Id, project.Description}
		projectContentsHomeObject = append(projectContentsHomeObject, projectObject)
		projectObject = ProjectContentsHome{}
	}
	return projectContentsHomeObject
}

/****
This struct will return all the picture
*/
type ProjectEditable struct {
	Project  project2.Project
	Images   []image3.Images
	History  history2.History
	Partners []partner.Partner
	Members  []member.Member
}

func GetProjectEditable(projectId string) ProjectEditable {
	var Images []image3.Images
	var Parteners []partner.Partner
	var Member []member.Member
	var historyToreturn history2.History
	projectEditable := ProjectEditable{}
	projectObject, err := project_io.ReadProject(projectId)
	if err != nil {
		fmt.Println(err, " can not read project")
		return projectEditable
	}
	//IMAGES
	projectImage, err := project_io.ReadAllOfProjectImage(projectObject.Id)
	if err != nil {
		fmt.Println(err, " error read project Image")
	}
	for _, image := range projectImage {
		ImageObejct, err := image_io.ReadImage(image.ImageId)
		if err != nil {
			fmt.Println(err, " error read Image")
		}
		Images = append(Images, ImageObejct)
	}
	//PARTNERS
	projectPartener, err := project_io.ReadAllOfProjectPartner(projectObject.Id)
	if err != nil {
		fmt.Println(err, " error read projectPartener")
	}
	for _, partners := range projectPartener {
		partner, err := partner_io.ReadPartner(partners.PartenerID)
		if err != nil {
			fmt.Println(err, " error read partner of id: ", partners.PartenerID)
		}
		Parteners = append(Parteners, partner)
	}
	//MEMBERS
	projectMemebers, err := project_io.ReadAllOfProjectMembers(projectObject.Id)
	if err != nil {
		fmt.Println(err, " error read projectMemeber")
	}
	for _, projectMember := range projectMemebers {
		mamber, err := member_io.ReadMember(projectMember.MemberId)
		if err != nil {
			fmt.Println(err, " error read mamber of id: ", projectMember.MemberId)
		}
		Member = append(Member, mamber)
	}
	projectHistory, err := project_io.ReadProjectHistoryOf(projectObject.Id)
	if err != nil {
		fmt.Println(err, " error read project History of id: ", projectObject.Id)
	} else {
		history, err := history_io.ReadHistory(projectHistory.HistoryId)
		if err != nil {
			fmt.Println(err, " error read history of id: ", projectHistory.HistoryId)
		}
		historyToreturn = history2.History{history.Id, history.Title, ConvertingToString(history.Content), history.Content, history.Date}
	}
	projectEditable = ProjectEditable{projectObject, Images, historyToreturn, Parteners, Member}
	return projectEditable
}
func ConvertingToString(byte []byte) string {
	s := string(byte)
	return s
}

//Help getting a collection Object that has collection Object and collection Type

type CollectionBridge struct {
	Collection     collection.Collection
	CollectionType collection.CollectionTypes
}

func GetCollectionBridge() []CollectionBridge {
	var collectionBridge []CollectionBridge
	var collectionType collection.CollectionTypes
	collections, err := collection_io.ReadCollections()
	if err != nil {
		fmt.Println(err, " error reading Collections")
	}
	for _, collection := range collections {
		collection_type, err := collection_io.ReadWithCollectionId(collection.Id)
		if err != nil {
			fmt.Println(err, " error reading collection_type")
		} else {
			collectionType, err = collection_io.ReadCollectionTyupe(collection_type.CollectionType)
		}
		collectionBridgeObject := CollectionBridge{collection, collectionType}
		collectionBridge = append(collectionBridge, collectionBridgeObject)
		collectionBridgeObject = CollectionBridge{}
	}
	return collectionBridge
}

//Users and thier Roles
type UsersAndRoles struct {
	User user2.Users
	Role string
}

func GetUserAndRole() []UsersAndRoles {
	var usersAndRoles []UsersAndRoles

	users, err := user_io.ReadUsers()
	if err != nil {
		fmt.Println(err, " Error reading Users")
		return usersAndRoles
	}
	for _, user := range users {
		userRole, err := user_io.ReadUserRoleWithEmail(user.Email)
		if err != nil {
			fmt.Println(err, " error reading user role of: "+user.Email)
		}
		role, err := user_io.ReadRole(userRole.RoleId)
		if err != nil {
			fmt.Println(err, " error reading role of the following role: "+userRole.RoleId)
		}
		user_role := UsersAndRoles{user, role.Role}
		usersAndRoles = append(usersAndRoles, user_role)
		user_role = UsersAndRoles{}
	}
	return usersAndRoles
}
