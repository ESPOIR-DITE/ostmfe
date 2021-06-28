package image_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/image"
	"ostmfe/utile"
	"testing"
)

func TestCreateImageType(t *testing.T) {
	object := image.ImageType{"", "other"}
	result, err := CreateImageType(object)
	assert.NotNil(t, err)
	fmt.Println(result)
}
func TestReadImageTypeWithName(t *testing.T) {
	result, err := ReadImageTypeWithName(utile.PROFILE)
	assert.Nil(t, err)
	fmt.Println(result)
}
