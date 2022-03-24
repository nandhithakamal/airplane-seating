package allocator

import (
	"airplane-seating/internal/seat"
	"airplane-seating/internal/seat/seatmap"
	"airplane-seating/internal/seat/seattype"
	"airplane-seating/internal/util"
	"errors"
	"fmt"
	"sort"
)

type Allocator struct {
	seatLayout                   seatmap.Layout
	numberOfPassengersToBeSeated int
}

func NewAllocator(l seatmap.Layout, n int) *Allocator {
	return &Allocator{
		seatLayout:                   l,
		numberOfPassengersToBeSeated: n,
	}
}

func (a *Allocator) AllocatePassengersToSeats() (seatMap *seatmap.SeatMap, err error) {
	m := seatmap.NewSeatMap(a.seatLayout)
	seats, err := m.Initialise()
	if err != nil {
		return nil, err
	}
	if a.numberOfPassengersToBeSeated > len(seats) {
		return nil, errors.New(fmt.Sprintf("Not enough seats for %v passengers", a.numberOfPassengersToBeSeated))
	}

	seatsSortedForAllocation := sortSeatsForAllocation(seats)
	a.blockSeatsForPassengers(seatsSortedForAllocation)

	return m, nil
}

func (a *Allocator) blockSeatsForPassengers(seats []*seat.Seat) {
	for i := 0; i < a.numberOfPassengersToBeSeated; i++ {
		util.HandleError(seats[i].BlockSeat(i + 1))
	}
}

func sortSeatsForAllocation(seats []*seat.Seat) []*seat.Seat {
	aisleSeats := filterSeatsByType(seats, seattype.AISLE)
	sortedAisleSeats := sortByRowAndColumn(aisleSeats)

	windowSeats := filterSeatsByType(seats, seattype.WINDOW)
	sortedWindowSeats := sortByRowAndColumn(windowSeats)

	middleSeats := filterSeatsByType(seats, seattype.MIDDLE)
	sortedMiddleSeats := sortByRowAndColumn(middleSeats)

	seatsSortedForAllocation := append(append(sortedAisleSeats, sortedWindowSeats...), sortedMiddleSeats...)
	return seatsSortedForAllocation
}

func sortByRowAndColumn(seats []*seat.Seat) []*seat.Seat {
	sort.Slice(seats, func(i, j int) bool {
		if seats[i].Row() == seats[j].Row() {
			return seats[i].Column() < seats[j].Column()
		}
		return seats[i].Row() < seats[j].Row()
	})
	return seats
}

func filterSeatsByType(seats []*seat.Seat, typeToFilterBy seattype.SeatType) []*seat.Seat {
	filteredSeats := make([]*seat.Seat, 0, 0)
	for i := range seats {
		if typeToFilterBy == seats[i].SeatType() {
			filteredSeats = append(filteredSeats, seats[i])
		}
	}
	return filteredSeats
}
