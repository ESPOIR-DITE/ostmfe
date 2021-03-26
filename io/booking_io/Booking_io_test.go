package booking_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/booking"
	"testing"
	"time"
)

func TestCreateBooking(t *testing.T) {
	result, err := CreateBooking(booking.Booking{"", "espoer", "0493646346", "cput", "espangola", "voleur", "congo", "kasai", "lksjdfghjfg", time.Now()})
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
