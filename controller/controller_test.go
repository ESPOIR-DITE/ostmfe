package controller

import (
	"fmt"
	"testing"
)

func TestCheckEventAndOdd(t *testing.T) {
	result := CheckEventAndOdd(6)
	fmt.Println(result)
}
