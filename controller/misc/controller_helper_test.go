package misc

import (
	"fmt"
	"testing"
)

func TestSeparateLatLng(t *testing.T) {
	longitude, latitude := SeparateLatLng("-34.195430264889836, 18.429533444457995")
	fmt.Println(longitude, " ", latitude)
}
