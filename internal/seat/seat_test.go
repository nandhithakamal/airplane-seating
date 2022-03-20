package seat

import (
	"airplane-seating/internal/seat/seatType"
	"github.com/stretchr/testify/assert"
	"testing"
)

//todo: idempotent calls on BlockSeat
func TestSeat_BlockSeat(t *testing.T) {
	seat := NewSeat(seatType.WINDOW, 1, 2)
	passengerId := 1

	err := seat.BlockSeat(passengerId)

	assert.NoError(t, err)
}

func TestSeat_BlockSeat_ShouldReturnErrorIfSeatIsUnavailable(t *testing.T) {
	seat := NewSeat(seatType.MIDDLE, 1, 2)
	passengerId := 1
	anotherPassengerId := 2
	setupErr := seat.BlockSeat(passengerId)
	assert.NoError(t, setupErr)

	err := seat.BlockSeat(anotherPassengerId)

	assert.Error(t, err)
}

func TestSeat_SeatType_ShouldReturnSeatType(t *testing.T) {
	windowSeat := NewSeat(seatType.WINDOW, 1, 1)
	middleSeat := NewSeat(seatType.MIDDLE, 1, 2)
	aisleSeat := NewSeat(seatType.AISLE, 1, 3)

	assert.Equal(t, seatType.WINDOW, windowSeat.SeatType())
	assert.Equal(t, seatType.MIDDLE, middleSeat.SeatType())
	assert.Equal(t, seatType.AISLE, aisleSeat.SeatType())
}

func TestSeat_Row_ShouldReturnRowOfSeat(t *testing.T) {
	frontSeat := NewSeat(seatType.WINDOW, 1, 1)
	backSeat := NewSeat(seatType.AISLE, 10, 1)

	assert.Equal(t, 1, frontSeat.Row())
	assert.Equal(t, 10, backSeat.Row())
}

func TestSeat_Column_ShouldReturnColumnOfSeat(t *testing.T) {
	windowSeat := NewSeat(seatType.WINDOW, 1, 1)
	middleSeat := NewSeat(seatType.MIDDLE, 1, 2)
	aisleSeat := NewSeat(seatType.AISLE, 1, 3)

	assert.Equal(t, 1, windowSeat.Column())
	assert.Equal(t, 2, middleSeat.Column())
	assert.Equal(t, 3, aisleSeat.Column())
}

func TestSeat_PassengerID_ShouldReturnPassengerAssignedToSeat(t *testing.T) {
	firstPassenger := 1
	secondPassenger := 2
	aisleSeat := NewSeat(seatType.AISLE, 1, 3)
	aisleSeat.BlockSeat(firstPassenger)

	secondAisleSeat := NewSeat(seatType.AISLE, 1, 4)
	secondAisleSeat.BlockSeat(secondPassenger)

	assert.Equal(t, firstPassenger, aisleSeat.PassengerId())
	assert.Equal(t, secondPassenger, secondAisleSeat.PassengerId())
}
