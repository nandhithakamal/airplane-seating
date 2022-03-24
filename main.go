package main

import (
	"airplane-seating/internal/allocator"
	"airplane-seating/internal/drawer"
	"airplane-seating/internal/parser"
	"bufio"
	"log"
	"os"
)

func main() {
	parser := parser.NewInputParser(bufio.NewReader(os.Stdin))
	inputLayout, numberOfPassengers, err := parser.ParseInput()
	if err != nil {
		log.Default().Fatal(err)
	}

	seatAllocator := allocator.NewAllocator(inputLayout, numberOfPassengers)
	allocatedSeatMap, err := seatAllocator.AllocatePassengersToSeats()
	if err != nil {
		log.Default().Fatalf("error allocating seats to passengers - %v\n", err)
	}

	drawer.DrawPassengerSeatMap(allocatedSeatMap)
}
