package login

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdminLogin(t *testing.T) {
	result := AdminLogin("espoir@gmail.com", "0001")
	assert.True(t, result)
	fmt.Println("result: ", result)
}
