package seat

import (
	"airplane-seating/internal/seatType"
	"github.com/stretchr/testify/assert"
	"testing"
)

//todo: idempotent calls on BlockSeat
func TestBlockSeat_ShouldReturnNilErrorIfSeatIsAvailable(t *testing.T) {
	seat := NewSeat(seatType.WINDOW, 1, 2)
	passengerId := 1

	err := seat.BlockSeat(passengerId)

	assert.NoError(t, err)
}

func TestBlockSeat_ShouldReturnErrorIfSeatIsUnavailable(t *testing.T) {
	seat := NewSeat(seatType.MIDDLE, 1, 2)
	passengerId := 1
	anotherPassengerId := 2
	setupErr := seat.BlockSeat(passengerId)
	assert.NoError(t, setupErr)

	err := seat.BlockSeat(anotherPassengerId)

	assert.Error(t, err)
}
