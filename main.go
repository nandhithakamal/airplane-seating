package main

import (
	"airplane-seating/internal/allocator"
	"airplane-seating/internal/drawer"
	"fmt"
)

func main() {
	seatAllocator := allocator.NewAllocator([][]int{{2, 3}, {3, 3}, {2, 3}}, 15)
	allocatedSeats, err := seatAllocator.AllocatePassengersToSeats()
	if err != nil {
		fmt.Errorf("error allocating seats to passengers - %v", err)
	}

	drawer.DrawPassengerSeatMap(allocatedSeats)
}
