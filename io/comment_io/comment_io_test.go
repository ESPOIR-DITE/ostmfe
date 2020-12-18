package comment_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadCommentEvent(t *testing.T) {

}
func TestReadAllByEventId(t *testing.T) {
	result, err := ReadAllByEventId("EF-17113801-42e3-4a28-8c57-8f32acb6819b")
	assert.Nil(t, err)
	fmt.Println(" result: ", result)
}
func TestReadAllByProjectId(t *testing.T) {
	result, err := ReadAllByProjectId("PF-243e37ea-06b8-40c1-865f-c6f68cb0ba1e")
	assert.Nil(t, err)
	fmt.Println(" result: ", result)
}

func TestReadComment(t *testing.T) {
	result, err := ReadComment("CF-9a4ba387-5dc7-4ec8-b1b8-e4de6c34cfa9")
	assert.Nil(t, err)
	fmt.Println(" result: ", result)
}
