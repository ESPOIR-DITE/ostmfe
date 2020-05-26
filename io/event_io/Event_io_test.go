package event_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	event2 "ostmfe/domain/event"
	"testing"
)

func TestCreateEventPlace(t *testing.T) {
	eventPlace := event2.EventPlace{"", "08783748", "description"}
	result, err := CreateEventPlace(eventPlace)
	assert.Nil(t, err)
	fmt.Println(result)
}
