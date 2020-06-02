package misc

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	image3 "ostmfe/domain/image"
	"ostmfe/domain/member"
	"ostmfe/domain/partner"
	project2 "ostmfe/domain/project"
	"ostmfe/io/image_io"
	"ostmfe/io/member_io"
	"ostmfe/io/partner_io"
	"ostmfe/io/project_io"
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
		fmt.Println(project.Title)
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
	Project        project2.Project
	Images         []image3.Images
	ProjectHistory project2.ProjectHistory
	Parteners      []partner.Partner
	Member         []member.Member
}

func GetProjectEditable(projectId string) ProjectEditable {
	var Images []image3.Images
	var Parteners []partner.Partner
	var Member []member.Member
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
	history, err := project_io.ReadProjectHistoryOf(projectObject.Id)
	if err != nil {
		fmt.Println(err, " error read project History of id: ", projectObject.Id)
	}
	projectEditable = ProjectEditable{projectObject, Images, history, Parteners, Member}
	return projectEditable
}
