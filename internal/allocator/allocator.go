package allocator

import (
	"airplane-seating/internal/seat"
	"airplane-seating/internal/seat/layout"
	"airplane-seating/internal/seat/seatType"
	"errors"
	"fmt"
	"sort"
)

type Allocator struct {
	seatLayout                   layout.Layout
	numberOfPassengersToBeSeated int
}

func NewAllocator(l layout.Layout, n int) *Allocator {
	return &Allocator{
		seatLayout:                   l,
		numberOfPassengersToBeSeated: n,
	}
}

func (a *Allocator) AllocatePassengersToSeats() error {
	seats, _ := layout.Initialise(a.seatLayout)
	if a.numberOfPassengersToBeSeated > len(seats) {
		return errors.New(fmt.Sprintf("Not enough seats for %v passengers", a.numberOfPassengersToBeSeated))
	}

	seatsSortedForAllocation := sortSeatsForAllocation(seats)
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>\n")
	layout.PrintSeats(seatsSortedForAllocation)

	a.blockSeatsForPassengers(seatsSortedForAllocation)
	fmt.Printf("--------------------\n")
	layout.PrintSeats(seats)
	return nil
}

func (a *Allocator) blockSeatsForPassengers(seats []*seat.Seat) {
	for i := 0; i < a.numberOfPassengersToBeSeated; i++ {
		seats[i].BlockSeat(i + 1)
	}
}

func sortSeatsForAllocation(seats []*seat.Seat) []*seat.Seat {
	aisleSeats := filterSeatsByType(seats, seatType.AISLE)
	sortedAisleSeats := sortByRowAndColumn(aisleSeats)
	//layout.PrintSeats(sortedAisleSeats)

	windowSeats := filterSeatsByType(seats, seatType.WINDOW)
	sortedWindowSeats := sortByRowAndColumn(windowSeats)
	//layout.PrintSeats(sortedWindowSeats)

	middleSeats := filterSeatsByType(seats, seatType.MIDDLE)
	sortedMiddleSeats := sortByRowAndColumn(middleSeats)
	//layout.PrintSeats(sortedMiddleSeats)

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

func filterSeatsByType(seats []*seat.Seat, typeToFilterBy seatType.SeatType) []*seat.Seat {
	filteredSeats := make([]*seat.Seat, 0, 0)
	for i, _ := range seats {
		if typeToFilterBy == seats[i].SeatType() {
			filteredSeats = append(filteredSeats, seats[i])
		}
	}
	return filteredSeats
}
