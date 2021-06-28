package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHomeClientPage(t *testing.T) {
	result, err := HomeClientPage()
	fmt.Println("Result: ", result)
	assert.NotNil(t, err)
}
func TestAboutClientPage(t *testing.T) {
	result, err := AboutClientPage()
	fmt.Println("Result: ", result)
	assert.NotNil(t, err)
}
