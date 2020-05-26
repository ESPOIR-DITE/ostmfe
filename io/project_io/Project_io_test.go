package project_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProject(t *testing.T) {

}
func TestDeleteProject(t *testing.T) {
	result, err := DeleteProject("PF-f860547c-1e29-46d9-b6a2-1cfac590a64c")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadProject(t *testing.T) {
	result, err := ReadProject("PF-f860547c-1e29-46d9-b6a2-1cfac590a64c")
	assert.Nil(t, err)
	fmt.Println(result)
}
