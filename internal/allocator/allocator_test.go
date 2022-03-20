package allocator

import (
	"airplane-seating/internal/seat"
	"airplane-seating/internal/seat/seatType"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllocator_AllocatePassengersToSeats_ShouldReturnErrorIfNotEnoughSeats(t *testing.T) {
	allocator := NewAllocator([][]int{{2, 2}, {2, 2}}, 10)

	allocatedSeats, err := allocator.AllocatePassengersToSeats()

	assert.Nil(t, allocatedSeats)
	assert.Error(t, err)
}

func TestAllocator_AllocatePassengersToSeats_ShouldNotReturnErrorIfEnoughSeatsAreAvailable(t *testing.T) {
	allocator := NewAllocator([][]int{{2, 2}, {2, 2}}, 6)

	allocatedSeats, err := allocator.AllocatePassengersToSeats()

	assert.NotNil(t, allocatedSeats)
	assert.NoError(t, err)
}

func TestAllocator_sortSeatsForAllocation_ShouldReturnSortedSeats(t *testing.T) {
	unsortedSeats := []*seat.Seat{
		seat.NewSeat(seatType.WINDOW, 1, 1),
		seat.NewSeat(seatType.AISLE, 1, 2),
		seat.NewSeat(seatType.WINDOW, 2, 1),
		seat.NewSeat(seatType.AISLE, 2, 2),
		seat.NewSeat(seatType.AISLE, 1, 3),
		seat.NewSeat(seatType.WINDOW, 1, 4),
		seat.NewSeat(seatType.AISLE, 2, 3),
		seat.NewSeat(seatType.WINDOW, 2, 4),
	}
	expectedSortedSeats := []*seat.Seat{
		seat.NewSeat(seatType.AISLE, 1, 2),
		seat.NewSeat(seatType.AISLE, 1, 3),
		seat.NewSeat(seatType.AISLE, 2, 2),
		seat.NewSeat(seatType.AISLE, 2, 3),
		seat.NewSeat(seatType.WINDOW, 1, 1),
		seat.NewSeat(seatType.WINDOW, 1, 4),
		seat.NewSeat(seatType.WINDOW, 2, 1),
		seat.NewSeat(seatType.WINDOW, 2, 4),
	}

	actualSortedSeats := sortSeatsForAllocation(unsortedSeats)

	assert.EqualValues(t, expectedSortedSeats, actualSortedSeats)
}

func TestAllocator_sortSeatsForAllocation_ShouldReturnSortedSeatsWithMiddleSeat(t *testing.T) {
	unsortedSeats := []*seat.Seat{
		seat.NewSeat(seatType.WINDOW, 1, 1),
		seat.NewSeat(seatType.MIDDLE, 1, 2),
		seat.NewSeat(seatType.AISLE, 1, 3),
		seat.NewSeat(seatType.WINDOW, 2, 1),
		seat.NewSeat(seatType.MIDDLE, 2, 2),
		seat.NewSeat(seatType.AISLE, 2, 3),

		seat.NewSeat(seatType.AISLE, 1, 4),
		seat.NewSeat(seatType.MIDDLE, 1, 5),
		seat.NewSeat(seatType.AISLE, 1, 6),
		seat.NewSeat(seatType.AISLE, 2, 4),
		seat.NewSeat(seatType.MIDDLE, 2, 5),
		seat.NewSeat(seatType.AISLE, 2, 6),
		seat.NewSeat(seatType.AISLE, 3, 4),
		seat.NewSeat(seatType.MIDDLE, 3, 5),
		seat.NewSeat(seatType.AISLE, 3, 6),

		seat.NewSeat(seatType.AISLE, 1, 7),
		seat.NewSeat(seatType.MIDDLE, 1, 8),
		seat.NewSeat(seatType.WINDOW, 1, 9),
		seat.NewSeat(seatType.AISLE, 2, 7),
		seat.NewSeat(seatType.MIDDLE, 2, 8),
		seat.NewSeat(seatType.WINDOW, 2, 9),
	}
	expectedSortedSeats := []*seat.Seat{
		seat.NewSeat(seatType.AISLE, 1, 3),
		seat.NewSeat(seatType.AISLE, 1, 4),
		seat.NewSeat(seatType.AISLE, 1, 6),
		seat.NewSeat(seatType.AISLE, 1, 7),
		seat.NewSeat(seatType.AISLE, 2, 3),
		seat.NewSeat(seatType.AISLE, 2, 4),
		seat.NewSeat(seatType.AISLE, 2, 6),
		seat.NewSeat(seatType.AISLE, 2, 7),
		seat.NewSeat(seatType.AISLE, 3, 4),
		seat.NewSeat(seatType.AISLE, 3, 6),

		seat.NewSeat(seatType.WINDOW, 1, 1),
		seat.NewSeat(seatType.WINDOW, 1, 9),
		seat.NewSeat(seatType.WINDOW, 2, 1),
		seat.NewSeat(seatType.WINDOW, 2, 9),

		seat.NewSeat(seatType.MIDDLE, 1, 2),
		seat.NewSeat(seatType.MIDDLE, 1, 5),
		seat.NewSeat(seatType.MIDDLE, 1, 8),
		seat.NewSeat(seatType.MIDDLE, 2, 2),
		seat.NewSeat(seatType.MIDDLE, 2, 5),
		seat.NewSeat(seatType.MIDDLE, 2, 8),
		seat.NewSeat(seatType.MIDDLE, 3, 5),
	}

	actualSortedSeats := sortSeatsForAllocation(unsortedSeats)

	assert.EqualValues(t, expectedSortedSeats, actualSortedSeats)
}
