package seat

import (
	"airplane-seating/internal/seatType"
	"errors"
)

type Seat struct {
	seatType    seatType.SeatType
	row         int
	column      int
	isAvailable bool
	passengerId int
}

func (s *Seat) BlockSeat(passengerId int) error {
	if !s.isAvailable {
		return errors.New("seat is already blocked")
	}
	s.passengerId = passengerId
	s.isAvailable = false

	return nil
}

func NewSeat(seatType seatType.SeatType, row, column int) *Seat {
	return &Seat{
		seatType,
		row,
		column,
		true,
		-1,
	}
}