package project_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadAllOfProjectMembers(t *testing.T) {
	pMember, err := ReadAllOfProjectMembers("PF-1af5ee3f-f2a7-4f87-9de8-e1f915f75260")
	assert.Nil(t, err)
	fmt.Println(pMember)
}
