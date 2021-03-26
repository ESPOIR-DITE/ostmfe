package user_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	user3 "ostmfe/domain/user"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := user3.Users{"espoioer", "espoir", "ditekemena"}
	result, err := CreateUser(user)
	assert.Nil(t, err)
	fmt.Println("result", result)
}
func TestDeleteUser(t *testing.T) {
	result, err := DeleteUser("espoioer")
	assert.Nil(t, err)
	fmt.Println("result", result)
}
func TestReadUser(t *testing.T) {
	result, err := ReadUser("user")
	assert.Nil(t, err)
	fmt.Println("result", result)
}
func TestUpdateUser(t *testing.T) {
	user := user3.Users{"espoir@gmail.com", "espoire", "ditekemena"}
	result, err := UpdateUser(user)
	assert.Nil(t, err)
	fmt.Println("result", result)
}
func TestReadUsers(t *testing.T) {
	result, err := ReadUsers()
	assert.Nil(t, err)
	fmt.Println("result", result)
}
