package project_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	project2 "ostmfe/domain/project"
	"testing"
)

func TestCreateProjectImage(t *testing.T) {
	var bytes [][]byte
	projectImage := project2.ProjectImage{"", "10000", "00001", "1"}
	projectImageHelper := project2.ProjectImageHelper{bytes, projectImage}
	result, err := CreateProjectImage(projectImageHelper)
	assert.NotNil(t, err)
	fmt.Println(err)
	fmt.Println(result)
}
