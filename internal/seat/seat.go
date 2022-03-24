package seat

import (
	"airplane-seating/internal/seat/seattype"
	"errors"
	"fmt"
)

const NO_PASSENGER = -1

type Seat struct {
	seatType    seattype.SeatType
	row         int
	column      int
	passengerId int
}

func NewSeat(seatType seattype.SeatType, row, column int) *Seat {
	return &Seat{
		seatType:    seatType,
		row:         row,
		column:      column,
		passengerId: NO_PASSENGER,
	}
}

func (s *Seat) BlockSeat(passengerId int) error {
	if !s.IsAvailable() {
		return errors.New("seat is already blocked")
	}
	s.passengerId = passengerId

	return nil
}

func (s *Seat) SeatType() seattype.SeatType {
	return s.seatType
}

func (s *Seat) Row() int {
	return s.row
}

func (s *Seat) Column() int {
	return s.column
}

func (s *Seat) PassengerId() int {
	return s.passengerId
}

func (s Seat) PrintSeat() {
	fmt.Printf("%v,%v-%v->%v\n", s.row, s.column, s.seatType, s.passengerId)
}

func (s *Seat) IsAvailable() bool {
	return s.passengerId == NO_PASSENGER
}
