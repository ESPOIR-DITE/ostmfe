package event_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	event2 "ostmfe/domain/event"
	"testing"
)

func TestCreateEventPlace(t *testing.T) {
	eventPlace := event2.EventPlace{"", "0000", "1111", ""}
	result, err := CreateEventPlace(eventPlace)
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventPlace(t *testing.T) {
	result, err := ReadEventPlace("EPF-61b377e0-7dee-4bc0-944b-25146be982f2")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventPlaceOf(t *testing.T) {
	result, err := ReadEventPlaceOf("EF-17113801-42e3-4a28-8c57-8f32acb6819b")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestUpdateEventPlace(t *testing.T) {
	eventPlace := event2.EventPlace{"EPF-61b377e0-7dee-4bc0-944b-25146be982f2", "1011", "1111", ""}
	result, err := UpdateEventPlace(eventPlace)
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestDeleteEventPlace(t *testing.T) {
	result, err := DeleteEventPlace("EPF-ae1d15f6-f1bb-4fdb-8b21-b993f9171848")
	assert.Nil(t, err)
	fmt.Println(result)
}

//Testing Event Image
func TestReadEventImgOf(t *testing.T) {
	result, err := ReadEventImgOf("EF-6e05e64c-34aa-460b-a742-529528eb4cdb")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventImg(t *testing.T) {
	result, err := ReadEventmgs()
	assert.Nil(t, err)
	fmt.Println(result)
}

func TestReadEventPeopleWithBoth(t *testing.T) {
	result, err := ReadEventPeopleWithBoth("EF-17113801-42e3-4a28-8c57-8f32acb6819b", "PF-40dcefba-0673-4f0e-bb95-9c14eeb40f79")
	assert.Nil(t, err)
	fmt.Println(result)
}

func TestCreateEventPlace2(t *testing.T) {

}

func TestCreateEventProject(t *testing.T) {
	EventProjectObejct := event2.EventProject{"", "1223", "00003", "test"}
	result, err := CreateEventProject(EventProjectObejct)
	assert.Nil(t, err)
	fmt.Println(result)
}

//event Group
func TestCreateEventGroup(t *testing.T) {
	object := event2.EventGroup{"", "0001", "11111"}
	result, err := CreateEventGroup(object)
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventGroup(t *testing.T) {
	result, err := ReadEventGroup("EGF-d55144e3-332a-4262-9f72-f577adf3df1b")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventGroupOf(t *testing.T) {
	result, err := ReadEventGroupOf("0001")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventGroupWithGroupId(t *testing.T) {
	result, err := ReadEventGroupWithGroupId("11111")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventGroupWithBoth(t *testing.T) {
	result, err := ReadEventGroupWithBoth("EF-4470cdf3-88d6-4f1f-a3ac-05151efedbb5", "GF-68e6d302-488f-450f-b196-5d5d3bf444c5")
	assert.Nil(t, err)
	fmt.Println(result)
}
