package people_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/people"
	"testing"
)

func TestCreatePeopleHistory(t *testing.T) {
	obejct := people.PeopleHistory{"", "003003", "029342"}
	result, err := CreatePeopleHistory(obejct)
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleHistoryWithPplId(t *testing.T) {
	result, err := ReadPeopleHistoryWithPplId("HF-7c6fc5b1-259e-45f9-8c9e-f2d8d383e383")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleHistorys(t *testing.T) {
	result, err := ReadPeopleHistorys()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleImage(t *testing.T) {
	result, err := ReadPeopleImageWithPeopleId("PF-35736764-ef06-46e0-afcb-dc17363cc1e0")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleImages(t *testing.T) {
	result, err := ReadPeopleImages()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}

func TestReadCategories(t *testing.T) {
	result, err := ReadCategories()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestDeleteCategory(t *testing.T) {
	result, err := DeleteCategory("CF-38fb1583-5246-45b9-b575-43698e345d12")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}

//Place

func TestCreatePeoplePlace(t *testing.T) {
	peoplePlace := people.PeoplePlace{"", "00001", "99383"}
	result, err := CreatePeoplePlace(peoplePlace)
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestDeletePeoplePlace(t *testing.T) {
	result, err := DeletePeoplePlace("PF-f75dba7a-2942-4c65-96dc-fada02d54c7f")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeoplePlaceAllByPlaceId(t *testing.T) {
	result, err := ReadPeoplePlaceAllByPlaceId("PF-35736764-ef06-46e0-afcb-dc17363cc1e0")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
