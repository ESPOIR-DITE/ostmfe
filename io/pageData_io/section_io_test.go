package pageData_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/pageData"
	"testing"
)

func TestCreateSection(t *testing.T) {
	sectionObejct := pageData.SectionBlock{"", "Notification", "test"}
	result, err := CreateSection(sectionObejct)
	assert.Nil(t, err)
	fmt.Println("Result: ", result)
}
func TestReadSection(t *testing.T) {

	result, err := ReadSection("SF-fe539a9d-6b95-4f3c-b041-503ea721c0e5")
	assert.Nil(t, err)
	fmt.Println("Result: ", result)
}
func TestReadPageDataWIthName(t *testing.T) {
	result, err := ReadPageDatas()
	assert.Nil(t, err)
	fmt.Println("Result: ", result)
}
func TestReadPageDataWIthName2(t *testing.T) {
	result, err := ReadPageDataWIthName("HomePage")
	assert.Nil(t, err)
	fmt.Println("Result: ", result)
}
func TestReadPageSection(t *testing.T) {
	result, err := ReadPageSection("PS-00155c2d-7ee6-4944-8e27-53d0030a86d6")
	assert.Nil(t, err)
	fmt.Println("Result: ", result)
}
