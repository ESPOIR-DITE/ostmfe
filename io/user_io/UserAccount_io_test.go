package user_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	user2 "ostmfe/domain/user"
	"testing"
	"time"
)

/***
The test class are create in many ways but, I prefer just to create a go file with the same name with the file that
I want to test and add a _test. example: UserAccount.go it test file should be UserAccount_test.go
And make sure that the file you are testing and the test file should be in the folder.
*/

func TestCreateUserAccount(t *testing.T) {
	// Creating an Object of type UserAccount
	useAccountObject := user2.UserAccount{"espoirditekemena@gmail.com", time.Now(), "xxxyyy"}
	/***
	Sending the Object to the back end
	if err is not nil, that means something is wrong. the test should fail.
	if err is nil that means everything is right and the test should pass
	*/
	result, err := CreateUserAccount(useAccountObject)
	assert.Nil(t, err)
	// we just displaying the result if all went well.
	fmt.Println("Result :", result)
}

func TestDeleteUserAccount(t *testing.T) {
	//Here we dont need an Object, we send only an Id of a user that we want to delete.
	result, err := DeleteUserAccount("espoirditekemena@gmail.com")
	assert.Nil(t, err)
	fmt.Println("Result :", result)
}

func TestReadUserAccount(t *testing.T) {
	result, err := ReadUserAccount("espoirditekemena@gmail.com")
	assert.Nil(t, err)
	fmt.Println("Result :", result)
}

func TestReadUserAccounts(t *testing.T) {
	result, err := ReadUserAccounts()
	assert.Nil(t, err)
	fmt.Println("Result :", result)
}

func TestUpdateUserAccount(t *testing.T) {
	useAccountObject := user2.UserAccount{"espoirditekemena@gmail.com", time.Now(), "xxxyyy"}
	result, err := UpdateUserAccount(useAccountObject)
	assert.Nil(t, err)
	fmt.Println("Result :", result)
}
