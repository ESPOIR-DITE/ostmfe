package image_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/image"
	"testing"
)

func TestCreateImageType(t *testing.T) {
	object := image.ImageType{"", "other"}
	result, err := CreateImageType(object)
	assert.NotNil(t, err)
	fmt.Println(result)
}
func TestReadImageTypeWithName(t *testing.T) {
	result, err := ReadImageTypeWithName("other")
	assert.Nil(t, err)
	fmt.Println(result)
}
