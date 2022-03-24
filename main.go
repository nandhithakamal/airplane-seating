package main

import (
	"airplane-seating/internal/allocator"
	"airplane-seating/internal/drawer"
	"fmt"
)

func main() {
	//seatAllocator := allocator.NewAllocator([][]int{{2, 2}, {3, 3}, {4, 2}, {2, 2}}, 22)
	seatAllocator := allocator.NewAllocator([][]int{{2, 3}, {3, 4}, {2, 2}}, 6)
	//seatAllocator := allocator.NewAllocator([][]int{{2, 3}, {2, 2}}, 4)
	allocatedSeatMap, err := seatAllocator.AllocatePassengersToSeats()
	if err != nil {
		fmt.Errorf("error allocating seats to passengers - %v", err)
	}

	drawer.DrawPassengerSeatMap(allocatedSeatMap)
}
