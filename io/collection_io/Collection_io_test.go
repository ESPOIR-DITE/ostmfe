package collection_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	collection2 "ostmfe/domain/collection"
	"testing"
)

func TestCreateCollection(t *testing.T) {
	hystory := []byte{10}
	collection := collection2.Collection{"", "00039484", "tst", hystory}
	result, err := CreateCollection(collection)
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestDeleteCollection(t *testing.T) {
	result, err := DeleteCollection("CF-02eb58c7-a2c6-4f8d-8942-b7c9882c9e2d")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
