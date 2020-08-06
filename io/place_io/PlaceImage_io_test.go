package place_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadPlaceImage(t *testing.T) {
	place, err := ReadPlaceImages()
	assert.Nil(t, err)
	fmt.Println("result ", place)
}
func TestReadPlaceImageAllOf(t *testing.T) {
	place, err := ReadPlaceImageAllOf("PF-c0bb4284-a07c-4cfd-aee7-f7e6a8406796")
	assert.Nil(t, err)
	fmt.Println("result ", place)
}
func TestReadPlaceImageWithImageId(t *testing.T) {
	place, err := ReadPlaceImageWithImageId("IF-00c1e970-8f52-4750-882f-7d8fdeb13e8c") //href="/admin_user/place/delete_image/IF-00c1e970-8f52-4750-882f-7d8fdeb13e8c/PIF-7ca8e485-a27a-4ca8-9f55-68b6380a4531"
	assert.Nil(t, err)
	fmt.Println("result ", place)
}
