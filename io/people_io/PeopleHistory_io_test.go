package people_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/people"
	"testing"
)

func TestCreatePeopleHistory(t *testing.T) {
	obejct := people.PeopleHistory{"", "003003", "029342"}
	result, err := CreatePeopleHistory(obejct)
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleHistoryWithPplId(t *testing.T) {
	result, err := ReadPeopleHistoryWithPplId("HF-7c6fc5b1-259e-45f9-8c9e-f2d8d383e383")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleHistorys(t *testing.T) {
	result, err := ReadPeopleHistorys()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
