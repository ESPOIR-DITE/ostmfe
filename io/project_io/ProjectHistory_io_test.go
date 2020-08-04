package project_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProjectHistory(t *testing.T) {
	//object:= project2.ProjectHistory{""}
	result, err := ReadProjectHistories()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadProjectHistoryOf(t *testing.T) {
	result, err := ReadProjectHistoryOf("PF-243e37ea-06b8-40c1-865f-c6f68cb0ba1e")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
