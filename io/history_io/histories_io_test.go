package history_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/history"
	"testing"
)

func TestReadHistorie(t *testing.T) {
	result, err := ReadHistorie("HF-f1f35cb8-459e-4cda-a61e-480febf0b871")
	fmt.Println(" result: ", result)
	assert.Nil(t, err)
}
func TestCreateHistory(t *testing.T) {
	mybite := ConvertToByteArray("<meta charset=\"utf-8\">\n  <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">")
	object := history.Histories{"", mybite}
	result, err := CreateHistorie(object)
	fmt.Println(" result: ", result)
	assert.Nil(t, err)
}
func TestUpdateHistorie(t *testing.T) {

	object := history.Histories{"HF-a33026da-4059-41e7-aa5a-4f6b6f1bc0d1", []byte{12}}
	result, err := CreateHistorie(object)
	fmt.Println(" result: ", result)
	assert.Nil(t, err)
}
func TestReadHistories(t *testing.T) {
	result, err := ReadHistories()
	fmt.Println(" result: ", result)
	assert.Nil(t, err)
}
func ConvertToByteArray(valeu string) []byte {
	toreturn := []byte(valeu)
	return toreturn
}
