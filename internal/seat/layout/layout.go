package layout

import (
	"airplane-seating/internal/seat"
	"fmt"
)

const leftmostSeatColumn = 1

type Layout [][]int

func Initialise(layout Layout) ([]*seat.Seat, error) {
	backMostSeatRow, rightMostSeatColumn, _ := computeExtremities(layout)
	aisles := computeAisleColumns(layout, rightMostSeatColumn)

	//todo: remove later
	fmt.Printf("backMost -> %v\n rightMost -> %v\n", backMostSeatRow, rightMostSeatColumn)
	fmt.Printf("aisles -> %v\n", aisles)

	return nil, nil
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
