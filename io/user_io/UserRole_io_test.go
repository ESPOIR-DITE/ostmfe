package user_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	user2 "ostmfe/domain/user"
	"testing"
)

func TestReadUserRoles(t *testing.T) {
	result, err := ReadUserRoles()
	assert.Nil(t, err)
	fmt.Println("result", result)
}
func TestCreateUserRole(t *testing.T) {
	object := user2.RoleOfUser{"URF-af345a91-35a7-4244-8f95-a8b3ebc31eeb", "espoirditekemena@gmail.com", "RF-d8bfa4aa-8281-4f63-9e4b-5a9b635db6f0"}
	result, err := CreateUserRole(object)
	assert.Nil(t, err)
	fmt.Println("result", result)
}
func TestDeleteUserRole(t *testing.T) {
	object := user2.RoleOfUser{"URF-5fc58579-f910-4de2-b96e-1bd51ae617cf", "ephrahimassani@gmail.com", "RF-b7172e40-bf23-4898-86ad-2451980b730a"}
	result, err := UpdateUserRole(object)
	assert.Nil(t, err)
	fmt.Println("result", result)
}
