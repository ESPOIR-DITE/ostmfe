package event_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	event2 "ostmfe/domain/event"
	"testing"
)

func TestCreateEventPlace(t *testing.T) {
	eventPlace := event2.EventPlace{"", "08783748", "description"}
	result, err := CreateEventPlace(eventPlace)
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventPlace(t *testing.T) {
	result, err := ReadEventPlace("EPF-9b93e869-ef34-41b2-8eca-68600dcc4b73")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadEventPlaceOf(t *testing.T) {
	result, err := ReadEventPlaceOf("EF-17113801-42e3-4a28-8c57-8f32acb6819b")
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
