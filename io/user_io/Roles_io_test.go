package user_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/user"
	"testing"
)

func TestCreateRole(t *testing.T) {
	role := user.Roles{"", "user 1", "minor privilages"}
	result, err := CreateRole(role)
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestDeleteRole(t *testing.T) {
	result, err := DeleteRole("RF-ceebde57-bf65-4131-9a24-00bcc36f2ce5")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadRole(t *testing.T) {
	result, err := ReadRoles()
	assert.Nil(t, err)
	fmt.Println(result)
}
