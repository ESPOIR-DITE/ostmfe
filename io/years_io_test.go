package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadYears(t *testing.T) {
	result, err := ReadYears()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
