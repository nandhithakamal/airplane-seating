package main

import (
	"airplane-seating/internal/allocator"
)

func main() {
	seatAllocator := allocator.NewAllocator([][]int{{2, 2}, {2, 2}}, 6)
	seatAllocator.AllocatePassengersToSeats()
}
