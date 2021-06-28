package misc

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	museum "ostmfe/domain"
	"ostmfe/domain/collection"
	"ostmfe/domain/event"
	history2 "ostmfe/domain/history"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/member"
	"ostmfe/domain/pageData"
	"ostmfe/domain/partner"
	people2 "ostmfe/domain/people"
	"ostmfe/domain/place"
	project2 "ostmfe/domain/project"
	user2 "ostmfe/domain/user"
	io2 "ostmfe/io"
	"ostmfe/io/collection_io"
	"ostmfe/io/contribution_io"
	"ostmfe/io/event_io"
	"ostmfe/io/history_io"
	"ostmfe/io/image_io"
	"ostmfe/io/member_io"
	"ostmfe/io/pageData_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/people_io"
	"ostmfe/io/place_io"
	"ostmfe/io/project_io"
	"ostmfe/io/user_io"
	"path/filepath"
	"strings"
	"time"
)

const (
	YYYYMMDD_FORMAT    = "2006-01-02"
	YYYMMDDTIME_FORMAT = "2006-01-02 15:04:05"
	layoutUS           = "January 2, 2006"
	RFC3339Nano        = "2006-01-02T15:00:00.000+0000"
)

/**
Format date in yyyy-MM-dd HH:mm:ss
*/

func FormatDateTime(date time.Time) string {
	return date.Format(YYYYMMDD_FORMAT)
}

//For sorting event slices
func ParseDateTime(date string) time.Time {
	result, err := time.Parse(YYYYMMDD_FORMAT, date)
	if err != nil {
		return time.Time{}
	}
	return result
}

/****
formating date with month name
*/
func FormatDateMonth(date string) string {
	t, _ := time.Parse(RFC3339Nano, date)
	return t.Format(YYYYMMDD_FORMAT)
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
	for index, project := range projects {
		//fmt.Println(project.Title)
		projectImage, err := project_io.ReadWithProjectIdProjectImage(project.Id)
		if err != nil {
			fmt.Println(err, " Can not find the following project in project image table: " /***, project.Title**/)
		} else {
			image, err = image_io.ReadImage(projectImage.ImageId)
			//fmt.Println(image.Image)
			if err != nil {
				fmt.Println(err, " Can not find the following project image Id in Image table: ")
			}
		}
		projectObject := ProjectContentsHome{project.Id, project.Title, image.Id, project.Description}
		projectContentsHomeObject = append(projectContentsHomeObject, projectObject)
		projectObject = ProjectContentsHome{}
		if index == 2 {
			break
		}
	}
	return projectContentsHomeObject
}

/****
This struct will return all the picture
*/
type ProjectEditable struct {
	Project  project2.Project
	Images   []ProjectImageHelperEditable
	History  history2.HistoriesHelper
	Partners []partner.Partner
	Members  []member.Member
}
type ProjectImageHelperEditable struct {
	Id             string
	ImageId        string
	ProjectImageId string
}

func GetProjectEditable(projectId string) ProjectEditable {
	var Images []ProjectImageHelperEditable
	var Parteners []partner.Partner
	var Member []member.Member
	var historyToreturn history2.HistoriesHelper
	projectEditable := ProjectEditable{}
	//Checking if the projectId is empty we should't event border to read....
	if projectId == "" {
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
			//Placing project image in image object at the place of description
			imageMakeUpObejct := ProjectImageHelperEditable{ImageObejct.Id, ConvertingToString(ImageObejct.Image), image.Id}
			Images = append(Images, imageMakeUpObejct)
			imageMakeUpObejct = ProjectImageHelperEditable{}
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
	//HistoryId
	projectHistory, err := project_io.ReadProjectHistoryOf(projectObject.Id)
	if err != nil {
		fmt.Println(err, " error read project HistoryId of id: ", projectObject.Id)
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
	for _, collectio := range collections {
		collection_type, err := collection_io.ReadWithCollectionId(collectio.Id)
		if err != nil {
			fmt.Println(err, " error reading collection_type")
		} else {
			collectionType, err = collection_io.ReadCollectionTyupe(collection_type.CollectionType)
		}
		collectionBridgeObject := CollectionBridge{collectio, collectionType}
		collectionBridge = append(collectionBridge, collectionBridgeObject)
		collectionBridgeObject = CollectionBridge{}
		collectionType = collection.CollectionTypes{}
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
		//fmt.Println("User Role",userRole)
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

		peopleObject := PeopleWithStringdate{people.Id, people.Name, people.Surname, FormatDateMonth(people.BirthDate), FormatDateMonth(people.DeathDate), people.Origin, people.Profession}
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
		peoplecategories, err := people_io.ReadPeopleCategoryWithCategoryId(people.Id)
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
		peoplecategories, err := people_io.ReadPeopleCategoryWithCategoryId(people.Id)
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
	Projects project2.Project
	Place    place.Place
	History  history2.HistoriesHelper
	Peoples  []people2.People
	Year     museum.Years
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
	var project project2.Project
	var place place.Place
	var imageHelper []EventImageHelperEditable
	var historyHelper history2.HistoriesHelper
	var peoples []people2.People
	var year museum.Years

	theEvent, err := event_io.ReadEvent(eventId)
	if err != nil {
		fmt.Println(err, "error reading event: ", eventId)
		return eventData
	}

	if theEvent.Id != "" {
		//First let's get the Images
		eventImages, err := event_io.ReadEventImgOf(theEvent.Id)
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
		eventPartners, err := event_io.ReadEventPartenerOf(theEvent.Id)
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
		eventYear, err := event_io.ReadEventYearWithEventId(eventId)
		if err != nil {
			fmt.Println(err, "error reading eventYear of: ", eventId)
		} else {
			year, err = io2.ReadYear(eventYear.YearId)
			if err != nil {
				fmt.Println(err, "error reading Year of: ", eventId)
			}

		}

		//thirdly, Projects
		eventProject, err := event_io.ReadEventProjectWithEventId(theEvent.Id)
		if err != nil {
			fmt.Println(err, "error reading eventProjects: ", eventId)
		} else {
			project, err = project_io.ReadProject(eventProject.ProjectId)
			if err != nil {
				fmt.Println(err, "error reading Project: ", eventId)
			}
		}
		//Fourth, Places
		eventplace, err := event_io.ReadEventPlaceOf(theEvent.Id)
		if err != nil {
			fmt.Println(err, "error reading event Place: ", eventId)
		} else {
			place, err = place_io.ReadPlace(eventplace.PlaceId)
			if err != nil {
				fmt.Println(err, "error reading Place EventPlaceId : ")
			}
		}
		//History
		eventHistory, err := event_io.ReadEventHistoryWithEventId(eventId)
		if err != nil {
			fmt.Println(err, "error reading eventHistory ")
		} else {
			history, err := history_io.ReadHistorie(eventHistory.HistoryId)
			if err != nil {
				fmt.Println(err, "error reading HistoryId")
			}
			historyHelper = history2.HistoriesHelper{history.Id, ConvertingToString(history.History)}
		}

		//People
		eventPeoples, err := event_io.ReadEventPeopleOf(eventId)
		if err != nil {
			fmt.Println(err, "error reading eventPEople ")
		} else {
			for _, eventPeople := range eventPeoples {
				people, err := people_io.ReadPeople(eventPeople.PeopleId)
				if err != nil {
					fmt.Println(err, "error reading people")
				} else {
					peoples = append(peoples, people)
				}
			}
		}

	}
	eventObejct := event.Event{theEvent.Id, theEvent.Name, FormatDateMonth(theEvent.Date), theEvent.IsPast, theEvent.Description}
	eventDataObject := EventData{eventObejct, imageHelper, partners, project, place, historyHelper, peoples, year}
	return eventDataObject
}

//Client Events
type SimpleEventData struct {
	Event        event.Event
	ProfileImage image3.Images
}

//Client Events
type SimpleEventDataLeft struct {
	Event        event.Event
	ProfileImage image3.Images
}

func GetSimpleEventData(limit int) []SimpleEventData {
	var profileImage image3.Images
	var eventDataList []SimpleEventData

	//Here we are reading all upcoming the events
	events, err := UpComingEvents()
	if err != nil {
		return eventDataList
	}
	for index, eventEntity := range events {
		eventImages, err := event_io.ReadEventImgOf(eventEntity.Id)
		if err != nil {
			fmt.Println(err, " error reading events Images")
		} else {
			//fmt.Println(" Looping eventImages")
			for _, eventImage := range eventImages {
				//fmt.Println(" eventImage.Description: ", eventImage.Description)
				if eventImage.Description == "1" || eventImage.Description == "profile" {
					//fmt.Println(" We have a profile Image")
					profileImage, err = image_io.ReadImage(eventImage.ImageId)
					if err != nil {
						fmt.Println(err, " error reading profile event image")
					}
				}
			}
		}
		//we need to make sure that profileImage is not empty
		if profileImage.Id != "" {
			eventObject := event.Event{eventEntity.Id, eventEntity.Name, FormatDateMonth(eventEntity.Date), eventEntity.IsPast, eventEntity.Description}
			eventData := SimpleEventData{eventObject, profileImage /** images**/}
			eventDataList = append(eventDataList, eventData)
			eventData = SimpleEventData{}
			profileImage = image3.Images{}
			eventObject = event.Event{}
		}
		//fmt.Println("This error may occur if there is no events created error:  profileImage is empty")

		// we are putting limit here so that the loop should exit if the index reach the limited number
		if index == limit {
			break
		}
	}

	return eventDataList
}

//This method like the top one, it returns all the events of a particular year.
func GetSimpleEventDataOfYear(yearId string) []SimpleEventData {
	var profileImage image3.Images
	var eventDataList []SimpleEventData

	eventYears, err := event_io.ReadEventYearsWithYearId(yearId)
	if err != nil {
		fmt.Println(err, " error reading event years")
	} else {
		for _, eventYear := range eventYears {

			eventEntity, err := event_io.ReadEvent(eventYear.EventId)
			if err != nil {
				fmt.Println(err, " error reading events")
			}

			eventImages, err := event_io.ReadEventImgOf(eventEntity.Id)
			if err != nil {
				fmt.Println(err, " error reading events Images")
			} else {
				fmt.Println(" Looping eventImages")
				for _, eventImage := range eventImages {
					//fmt.Println(" eventImage.Description: ", eventImage.Description)
					if eventImage.Description == "1" || eventImage.Description == "profile" {
						fmt.Println(" We have a profile Image")
						profileImage, err = image_io.ReadImage(eventImage.ImageId)
						if err != nil {
							fmt.Println(err, " error reading profile event image")
						}
					}
				}
			}
			//we need to make sure that profileImage is not empty
			if profileImage.Id != "" {
				//fmt.Println(" profileImage.Id: ", profileImage.Id)
				eventObject := event.Event{eventEntity.Id, eventEntity.Name, FormatDateMonth(eventEntity.Date), eventEntity.IsPast, eventEntity.Description}
				eventData := SimpleEventData{eventObject, profileImage /** images**/}
				eventDataList = append(eventDataList, eventData)
				eventData = SimpleEventData{}
				profileImage = image3.Images{}
				eventObject = event.Event{}

				//adding data to the correct list
				//if CheckEventAndOdd(index)
			}
			fmt.Println("This error may occur if there is no events created error:  profileImage is empty")

			// we are putting limit here so that the loop should exit if the index reach the limited number

		}
	}

	return eventDataList
}

//dealing with sidebar data
type SidebarData struct {
	PageData []pageData.PageData
	Menu     string
	Submenu  string
}

func GetSideBarData(menu, submenu string) SidebarData {
	var pageDataObject []pageData.PageData
	//Reading all the Pages
	pageData, err := pageData_io.ReadPageDatas()
	if err != nil {
		fmt.Println(err, " error reading all the PageData")
		return SidebarData{pageDataObject, menu, submenu}
	}
	return SidebarData{pageData, menu, submenu}
}

//This method help to get a contributor file type
func GetFileExtension(fileData *multipart.FileHeader) (bool, string) {
	var extension = filepath.Ext(fileData.Filename)
	contributionFileTypes, err := contribution_io.ReadContributionFileTypes()
	if err != nil {
		fmt.Println("error reading contributionFileType")
		return true, ""
	} else {
		for _, contributionFileType := range contributionFileTypes {
			fmt.Println("extension: " + extension + " file extension: " + contributionFileType.FileType)
			//t := strings.Trim(extension, ".")
			t := strings.Replace(extension, ".", "", -1)
			fmt.Println("extension2: " + t + " file extension: " + contributionFileType.FileType)
			if t == contributionFileType.FileType {
				return true, contributionFileType.Id
			}
		}
	}
	return false, ""
}

func GetGalleryImage(galleryId, bridgeId string) image3.GalleryHelper {
	var galleryImage image3.GalleryHelper
	gallery, err := image_io.ReadGalleryH(galleryId)
	if err != nil {
		fmt.Println(err, " error reading image")
		return galleryImage
	}
	return image3.GalleryHelper{gallery.Id, gallery.Image, gallery.Description, bridgeId}
}

func GetBanner(pageName string) (pageData.BannerImageHelper, error) {
	page, err := pageData_io.ReadPageDataWIthName(pageName)
	if err != nil {
		return pageData.BannerImageHelper{}, err
	}
	banner, err := pageData_io.ReadBannerN(page.BannerId)
	if err != nil {
		return pageData.BannerImageHelper{}, err
	}
	return banner, nil
}
