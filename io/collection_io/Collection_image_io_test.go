package collection_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadCollectionImgs(t *testing.T) {
	result, err := ReadCollectionImgs()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadCollectionImg(t *testing.T) {
	result, err := ReadCollectionImgs()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadCollectionImgsWithCollectionId(t *testing.T) {
	result, err := ReadCollectionImgsWithCollectionId("CF-31e713ee-fa92-42bd-afc9-2d84aca82314")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
