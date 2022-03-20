package seat

import (
	"airplane-seating/internal/seat/seatType"
	"errors"
	"fmt"
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
		seatType:    seatType,
		row:         row,
		column:      column,
		isAvailable: true,
		passengerId: -1,
	}
}

func (s Seat) PrintSeat() {
	fmt.Printf("%v,%v-%v\n", s.row, s.column, s.seatType)
}
