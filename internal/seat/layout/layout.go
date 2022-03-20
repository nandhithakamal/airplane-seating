package layout

import (
	"airplane-seating/internal/seat"
	"fmt"
)

const leftmostSeatColumn = 1

type Layout [][]int

func Initialise(layout Layout) ([]*seat.Seat, error) {
	backMostSeatRow, rightMostSeatColumn, _ := computeExtremities(layout)
	aisleSeatColumns := computeAisleColumns(layout, rightMostSeatColumn)
	windowSeatColumns := computeWindowColumns(rightMostSeatColumn)
	middleSeatColumns := computeMiddleColumns(windowSeatColumns, aisleSeatColumns)

	//todo: remove later
	fmt.Printf("backMost -> %v\n rightMost -> %v\n", backMostSeatRow, rightMostSeatColumn)
	fmt.Printf("aisles -> %v\n", aisleSeatColumns)
	fmt.Printf("windows -> %v\n", windowSeatColumns)
	fmt.Printf("middles -> %v\n", middleSeatColumns)

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

func computeWindowColumns(rightMost int) []int {
	return []int{leftmostSeatColumn, rightMost}
}

func computeMiddleColumns(windowColumns, aisleColumns []int) []int {
	var middleColumns []int
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
