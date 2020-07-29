package misc

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"ostmfe/domain/collection"
	"ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/member"
	"ostmfe/domain/partner"
	people2 "ostmfe/domain/people"
	"ostmfe/domain/place"
	project2 "ostmfe/domain/project"
	user2 "ostmfe/domain/user"
	"ostmfe/io/collection_io"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/member_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
	"ostmfe/io/project_io"
	"ostmfe/io/user_io"
	"strings"
	"time"
)

const (
	YYYYMMDD_FORMAT    = "2006-01-02"
	YYYMMDDTIME_FORMAT = "2006-01-02 15:04:05"
	layoutUS           = "January 2, 2006"
	layoutISO          = "2006-01-02"
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

/****
formating date with month name
*/
func FormatingDateMonth(date string) string {
	t, _ := time.Parse(YYYMMDDTIME_FORMAT, date)
	return t.Format(layoutUS)
}

/***
Converting String to []byte
*/
func ConvertToByteArray(valeu string) []byte {
	toreturn := []byte(valeu)
	return toreturn
}

func ConvertingToString(bytes []byte) string {
	toreturn := string(bytes)
	return toreturn
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
	fmt.Println(parts)
	latitudeString := parts[0]
	longitudeString := parts[1]
	//latitude, _ = strconv.ParseFloat(latitudeString, 64)
	//longitude, _ = strconv.ParseFloat(longitudeString, 64)
	return latitudeString, longitudeString
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
	History  history2.HistoriesHelper
	Partners []partner.Partner
	Members  []member.Member
}

func GetProjectEditable(projectId string) ProjectEditable {
	var Images []image3.Images
	var Parteners []partner.Partner
	var Member []member.Member
	var historyToreturn history2.HistoriesHelper
	projectEditable := ProjectEditable{}
	//Checking if the projectId is empty
	if projectId != "" {
		fmt.Println(" can not read this people, because the projectId is empty or null")
		return projectEditable
	}
	projectObject, err := project_io.ReadProject(projectId)
	if err != nil {
		fmt.Println(err, " can not read project")
		return projectEditable
	}
	//IMAGES
	projectImage, err := project_io.ReadAllOfProjectImage(projectObject.Id)
	if err != nil {
		fmt.Println(err, " error read project Image")
	} else {
		for _, image := range projectImage {
			ImageObejct, err := image_io.ReadImage(image.ImageId)
			if err != nil {
				fmt.Println(err, " error read Image")
			}
			Images = append(Images, ImageObejct)
		}
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
	//History
	projectHistory, err := project_io.ReadProjectHistoryOf(projectObject.Id)
	if err != nil {
		fmt.Println(err, " error read project History of id: ", projectObject.Id)
	} else {
		history, err := history_io.ReadHistorie(projectHistory.HistoryId)
		if err != nil {
			fmt.Println(err, " error read history of id: ", projectHistory.HistoryId)
		}
		historyToreturn = history2.HistoriesHelper{history.Id, ConvertingToString(history.History)}
	}
	projectEditable = ProjectEditable{projectObject, Images, historyToreturn, Parteners, Member}
	return projectEditable
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

//Help to get people with correct dates

type PeopleWithStringdate struct {
	Id         string
	Name       string
	Surname    string
	BirthDate  string
	DeathDate  string
	Origin     string
	Profession string
}

func GetPeopleWithStringdate() []PeopleWithStringdate {
	var peopelToreturn []PeopleWithStringdate
	peoples, err := people_io.ReadPeoples()
	if err != nil {
		fmt.Println(err, "error reading peoples")
		return peopelToreturn
	}
	for _, people := range peoples {

		dateOfBirth := people.BirthDate.Format("2006-01-02")
		fmt.Println(people.BirthDate)
		fmt.Println(dateOfBirth)
		fmt.Println("yyyy-mm-dd : ", people.BirthDate.Format("2006-01-02"))
		dateOfDearth := people.DeathDate.Format("2006-01-02")
		peopleObject := PeopleWithStringdate{people.Id, people.Name, people.Surname, dateOfBirth, dateOfDearth, people.Origin, people.Profession}
		peopelToreturn = append(peopelToreturn, peopleObject)
		peopleObject = PeopleWithStringdate{}
	}
	return peopelToreturn
}

//Get people data
type PeopleData struct {
	People     people2.People
	Images     []image3.Images
	Profession []people2.Profession
	History    history2.History
	Category   []people2.Category
}

// GET ALL PEOPLE
func GetPeopleDataList() []PeopleData {
	var peopleData []PeopleData
	var professions []people2.Profession
	var categoryList []people2.Category
	var imageList []image3.Images
	var history history2.History
	peoples, err := people_io.ReadPeoples()
	if err != nil {
		fmt.Println(err, "error reading peoples")
		return peopleData
	}
	for _, people := range peoples {
		//getting the Images od this person
		peopleImages, err := people_io.ReadPeopleImagewithPeopleId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading peopleImage for peopleId: ", people.Id)
		} else {
			for _, peopleImage := range peopleImages {
				images, err := image_io.ReadImage(peopleImage.ImageId)
				if err != nil {
					fmt.Println(err, "error reading peopleImage for peopleImageId: ", peopleImage.Id)
				} else {
					imageList = append(imageList, images)
				}
			}
		}

		//Getting People proffesions
		peopleProfession, err := people_io.ReadPeopleProfessionWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading peopleProfession for peopleId: ", people.Id)
		} else {
			for _, peopleProfession := range peopleProfession {
				profession, err := people_io.ReadProfession(peopleProfession.Profession)
				if err != nil {
					fmt.Println(err, "error reading Profession for peopleId: ", people.Id)
				} else {
					professions = append(professions, profession)
				}
			}
		}

		//Getting People history
		peopleHistory, err := people_io.ReadPeopleHistoryWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading people history for peopleId: ", people.Id)
		} else {
			history, err = history_io.ReadHistory(peopleHistory.HistoryId)
			if err != nil {
				fmt.Println(err, "error reading history: ", people.Id)
			}
		}

		//Getting People category
		peoplecategories, err := people_io.ReadPeopleCategoryWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading people category for peopleId: ", people.Id)
		} else {
			for _, peoplecategory := range peoplecategories {
				category, err := people_io.ReadCategory(peoplecategory.Id)
				if err != nil {
					fmt.Println(err, "error reading category for peopleId: ", peoplecategory.Id)
				} else {
					categoryList = append(categoryList, category)
				}
			}
		}
		peopleDataObject := PeopleData{people, imageList, professions, history, categoryList}
		peopleData = append(peopleData, peopleDataObject)
		peopleDataObject = PeopleData{}
	}
	return peopleData
}

//GET A PEOPLE WITH ALL HIS DATA
func GetPeopleData(peopleId string) PeopleData {
	var peopleData PeopleData
	var professions []people2.Profession
	var categoryList []people2.Category
	var imageList []image3.Images
	var history history2.History
	people, err := people_io.ReadPeople(peopleId)
	if err != nil {
		fmt.Println(err, "error reading peoples")
		return peopleData
	} else {
		//getting the Images od this person
		peopleImages, err := people_io.ReadPeopleImagewithPeopleId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading peopleImage for peopleId: ", people.Id)
		} else {
			for _, peopleImage := range peopleImages {
				images, err := image_io.ReadImage(peopleImage.ImageId)
				if err != nil {
					fmt.Println(err, "error reading peopleImage for peopleImageId: ", peopleImage.Id)
				} else {
					imageList = append(imageList, images)
				}
			}
		}

		//Getting People proffesions
		peopleProfession, err := people_io.ReadPeopleProfessionWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading peopleProfession for peopleId: ", people.Id)
		} else {
			for _, peopleProfession := range peopleProfession {
				profession, err := people_io.ReadProfession(peopleProfession.Profession)
				if err != nil {
					fmt.Println(err, "error reading Profession for peopleId: ", people.Id)
				} else {
					professions = append(professions, profession)
				}
			}
		}

		//Getting People history
		peopleHistory, err := people_io.ReadPeopleHistoryWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading people history for peopleId: ", people.Id)
		} else {
			history, err = history_io.ReadHistory(peopleHistory.HistoryId)
			if err != nil {
				fmt.Println(err, "error reading history: ", people.Id)
			}
		}

		//Getting People category
		peoplecategories, err := people_io.ReadPeopleCategoryWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading people category for peopleId: ", people.Id)
		} else {
			for _, peoplecategory := range peoplecategories {
				category, err := people_io.ReadCategory(peoplecategory.Id)
				if err != nil {
					fmt.Println(err, "error reading category for peopleId: ", peoplecategory.Id)
				} else {
					categoryList = append(categoryList, category)
				}
			}
		}
		peopleDataObject := PeopleData{people, imageList, professions, history, categoryList}
		return peopleDataObject
	}
	return peopleData
}

//GET ALL PEOPLE OF CATEGORY
func GetPeopleDataListOfCategory(category string) []PeopleData {
	var peopleData []PeopleData
	var professions []people2.Profession
	var categoryList []people2.Category
	var imageList []image3.Images
	var history history2.History

	//Getting People category
	peoplecategories, err := people_io.ReadPeopleCategoryWithCategoryId(category)
	if err != nil {
		fmt.Println(err, "error reading people category for peopleId: ", category)
	}
	for _, peopleCategory := range peoplecategories {

		people, err := people_io.ReadPeople(peopleCategory.PeopleId)
		if err != nil {
			fmt.Println(err, "error reading peoples")
			return peopleData
		}
		//getting the Images od this person
		peopleImages, err := people_io.ReadPeopleImagewithPeopleId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading peopleImage for peopleId: ", people.Id)
		} else {
			for _, peopleImage := range peopleImages {
				images, err := image_io.ReadImage(peopleImage.ImageId)
				if err != nil {
					fmt.Println(err, "error reading peopleImage for peopleImageId: ", peopleImage.Id)
				} else {
					imageList = append(imageList, images)
				}
			}
		}

		//Getting People proffesions
		peopleProfession, err := people_io.ReadPeopleProfessionWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading peopleProfession for peopleId: ", people.Id)
		} else {
			for _, peopleProfession := range peopleProfession {
				profession, err := people_io.ReadProfession(peopleProfession.Profession)
				if err != nil {
					fmt.Println(err, "error reading Profession for peopleId: ", people.Id)
				} else {
					professions = append(professions, profession)
				}
			}
		}

		//Getting People history
		peopleHistory, err := people_io.ReadPeopleHistoryWithPplId(people.Id)
		if err != nil {
			fmt.Println(err, "error reading people history for peopleId: ", people.Id)
		} else {
			history, err = history_io.ReadHistory(peopleHistory.HistoryId)
			if err != nil {
				fmt.Println(err, "error reading history: ", people.Id)
			}
		}
		peopleDataObject := PeopleData{people, imageList, professions, history, categoryList}
		peopleData = append(peopleData, peopleDataObject)
		peopleDataObject = PeopleData{}
	}
	return peopleData
}

//Event data type
type EventData struct {
	Event    event.Event
	Images   []EventImageHelperEditable
	Partners []partner.Partner
	Projects []project2.Project
	Place    place.Place
	History  history2.HistoriesHelper
}
type EventImageHelperEditable struct {
	Id           string
	ImageId      string
	EventImageId string
}

func GetEventDate(eventId string) EventData {
	var eventData EventData
	//var images []image3.Images
	var partners []partner.Partner
	var projects []project2.Project
	var place place.Place
	var imageHelper []EventImageHelperEditable
	var historyHelper history2.HistoriesHelper

	event, err := event_io.ReadEvent(eventId)
	if err != nil {
		fmt.Println(err, "error reading event: ", eventId)
		return eventData
	}
	if event.Id != "" {
		//First let's get the Images
		eventImages, err := event_io.ReadEventImgOf(event.Id)
		if err != nil {
			fmt.Println(err, "error reading eventImages: ", eventId)
		} else {
			for _, eventImage := range eventImages {
				image, err := image_io.ReadImage(eventImage.ImageId)
				if err != nil {
					fmt.Println(err, "error reading Images for eventImageId: ", eventImage.ImageId)
				} else {
					imageHelperObject := EventImageHelperEditable{image.Id, ConvertingToString(image.Image), eventImage.Id}
					imageHelper = append(imageHelper, imageHelperObject)
				}
			}
		}

		//Second, Partners
		eventPartners, err := event_io.ReadEventPartenerOf(event.Id)
		if err != nil {
			fmt.Println(err, "error reading eventPartners: ", eventId)
		} else {
			for _, eventPartner := range eventPartners {
				partner, err := partner_io.ReadPartner(eventPartner.PartenerId)
				if err != nil {
					fmt.Println(err, "error reading Partners, PartnerId: ", eventPartner.PartenerId)
				} else {
					partners = append(partners, partner)
				}
			}
		}

		//thirdly, Projects
		eventProjects, err := event_io.ReadEventProjectOf(event.Id)
		if err != nil {
			fmt.Println(err, "error reading eventProjects: ", eventId)
		} else {
			for _, eventProject := range eventProjects {
				project, err := project_io.ReadProject(eventProject.ProjectId)
				if err != nil {
					fmt.Println(err, "error reading Projects: ", eventId)
				} else {
					projects = append(projects, project)
				}
			}
		}
		//Fourth, Places
		eventplace, err := event_io.ReadEventPlaceOf(event.Id)
		if err != nil {
			fmt.Println(err, "error reading event Place: ", eventId)
		} else {
			place, err = place_io.ReadPlace(eventplace.PlaceId)
			if err != nil {
				fmt.Println(err, "error reading Place EventPlaceId : ")
			}
		}
		eventHistory, err := event_io.ReadEventHistoryWithEventId(eventId)
		if err != nil {
			fmt.Println(err, "error reading eventHistory ")
		} else {
			history, err := history_io.ReadHistorie(eventHistory.HistoryId)
			if err != nil {
				fmt.Println(err, "error reading History")
			}
			historyHelper = history2.HistoriesHelper{history.Id, ConvertingToString(history.History)}
		}
	}
	eventDataObject := EventData{event, imageHelper, partners, projects, place, historyHelper}
	return eventDataObject
}
