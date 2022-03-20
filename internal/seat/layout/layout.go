package layout

import (
	"airplane-seating/internal/seat"
	"airplane-seating/internal/seat/seatType"
	"fmt"
)

const leftmostSeatColumn = 1

type Layout [][]int

var aisleSeatColumns []int
var windowSeatColumns []int
var middleSeatColumns []int

func Initialise(layout Layout) ([]*seat.Seat, error) {
	backMostSeatRow, rightMostSeatColumn, _ := computeExtremities(layout)
	aisleSeatColumns = computeAisleColumns(layout, rightMostSeatColumn)
	windowSeatColumns = computeWindowColumns(rightMostSeatColumn)
	middleSeatColumns = computeMiddleColumns(windowSeatColumns, aisleSeatColumns)

	var seats []*seat.Seat
	var finalColumnOfLastGroup int
	for _, group := range layout {
		rows := group[0]
		columns := group[1]
		for r := 1; r <= rows; r++ {
			for c := 1; c <= columns; c++ {
				seatType := findSeatType(finalColumnOfLastGroup + c)
				seats = append(seats, seat.NewSeat(seatType, r, finalColumnOfLastGroup+c))
			}
		}
		finalColumnOfLastGroup += columns
	}

	//todo: remove later
	fmt.Printf("backMost -> %v\n rightMost -> %v\n", backMostSeatRow, rightMostSeatColumn)
	fmt.Printf("aisles -> %v\n", aisleSeatColumns)
	fmt.Printf("windows -> %v\n", windowSeatColumns)
	fmt.Printf("middles -> %v\n", middleSeatColumns)
	PrintSeats(seats)

	return seats, nil
}

func PrintSeats(seats []*seat.Seat) {
	for i, _ := range seats {
		seats[i].PrintSeat()
	}
}

func findSeatType(column int) seatType.SeatType {
	if isElementPresent(column, windowSeatColumns) {
		return seatType.WINDOW
	}
	if isElementPresent(column, aisleSeatColumns) {
		return seatType.AISLE
	}
	return seatType.MIDDLE
}

func computeExtremities(layout Layout) (int, int, error) {
	backMost := 0
	rightMost := 0
	for _, group := range layout {
		if backMost < group[0] {
			backMost = group[0]
		}
		rightMost += group[1]
	}
	return backMost, rightMost, nil
}

func computeAisleColumns(layout Layout, rightMost int) []int {
	var finalColumnOfLastGroup int
	var beginningColumnOfCurrentGroup int
	var finalColumnOfCurrentGroup int
	var aisleColumns []int
	for _, group := range layout {
		beginningColumnOfCurrentGroup = finalColumnOfLastGroup + 1
		finalColumnOfCurrentGroup = finalColumnOfLastGroup + group[1]
		if beginningColumnOfCurrentGroup != leftmostSeatColumn {
			aisleColumns = append(aisleColumns, beginningColumnOfCurrentGroup)
		}
		if finalColumnOfCurrentGroup != rightMost {
			aisleColumns = append(aisleColumns, finalColumnOfCurrentGroup)
		}
		finalColumnOfLastGroup = finalColumnOfCurrentGroup
	}
	return aisleColumns
}

func computeWindowColumns(rightMost int) []int {
	return []int{leftmostSeatColumn, rightMost}
}

func computeMiddleColumns(windowColumns, aisleColumns []int) []int {
	middleColumns := make([]int, 0, 0)
	for i := leftmostSeatColumn + 1; i < windowColumns[1]; i++ {
		if !isElementPresent(i, aisleColumns) {
			middleColumns = append(middleColumns, i)
		}
	}
	return middleColumns
}

func isElementPresent(element int, arr []int) bool {
	for _, i := range arr {
		if i == element {
			return true
		}
	}
	return false
}
