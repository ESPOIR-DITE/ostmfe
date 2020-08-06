package place_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	place2 "ostmfe/domain/place"
	"testing"
)

func TestCreatePlace(t *testing.T) {
	long := "-47.0983474323422"
	lat := "-47.0983474323422"
	placeObeject := place2.Place{"", "example", long, lat, "test1"}
	place, err := CreatePlace(placeObeject)
	assert.Nil(t, err)
	fmt.Println("result ", place)
}
