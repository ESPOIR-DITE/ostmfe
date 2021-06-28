package admin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHomeAdminPage(t *testing.T) {
	result, err := GetHomeAdminPage("")
	fmt.Println("result : ", result)
	assert.NotNil(t, err)
}

//Event
func TestGetEventEditData(t *testing.T) {
	result, err := GetEventEditData("EF-4893ea0f-6b04-4ac8-b664-367c16bd31c3")
	fmt.Println("result : ", result)
	assert.NotNil(t, err)
}
