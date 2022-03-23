package seatmap

import (
	"airplane-seating/internal/seat"
	"airplane-seating/internal/seat/seattype"
	"airplane-seating/internal/util"
	"fmt"
)

const LEFT_MOST_COLUMN = 1

type Layout [][]int

type SeatMap struct {
	layout          Layout
	rightMostColumn int
	backMostRow     int
}

func NewSeatMap(layout Layout) *SeatMap {
	backMostRow, rightMostColumn, _ := computeExtremities(layout)
	return &SeatMap{
		layout:          layout,
		rightMostColumn: rightMostColumn,
		backMostRow:     backMostRow,
	}
}

var aisleSeatColumns []int
var windowSeatColumns []int
var middleSeatColumns []int

func (m *SeatMap) Initialise() ([]*seat.Seat, error) {
	aisleSeatColumns = m.computeAisleColumns()
	windowSeatColumns = m.computeWindowColumns()
	middleSeatColumns = m.computeMiddleColumns(windowSeatColumns, aisleSeatColumns)

	var seats []*seat.Seat
	var finalColumnOfLastGroup int
	for _, group := range m.layout {
		rows := group[0]
		columns := group[1]
		for r := 1; r <= rows; r++ {
			for c := 1; c <= columns; c++ {
				typeOfSeat := findSeatType(finalColumnOfLastGroup + c)
				seats = append(seats, seat.NewSeat(typeOfSeat, r, finalColumnOfLastGroup+c))
			}
		}
		finalColumnOfLastGroup += columns
	}

	//todo: remove later
	fmt.Printf("backMost -> %v\n rightMost -> %v\n", m.backMostRow, m.rightMostColumn)
	fmt.Printf("aisles -> %v\n", aisleSeatColumns)
	fmt.Printf("windows -> %v\n", windowSeatColumns)
	fmt.Printf("middles -> %v\n", middleSeatColumns)
	PrintSeats(seats)

	return seats, nil
}

func PrintSeats(seats []*seat.Seat) {
	for i := range seats {
		seats[i].PrintSeat()
	}
}

func findSeatType(column int) seattype.SeatType {
	if util.IsElementPresent(column, windowSeatColumns) {
		return seattype.WINDOW
	}
	if util.IsElementPresent(column, aisleSeatColumns) {
		return seattype.AISLE
	}
	return seattype.MIDDLE
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

func (m *SeatMap) computeAisleColumns() []int {
	var finalColumnOfLastGroup int
	var beginningColumnOfCurrentGroup int
	var finalColumnOfCurrentGroup int
	var aisleColumns []int
	for _, group := range m.layout {
		beginningColumnOfCurrentGroup = finalColumnOfLastGroup + 1
		finalColumnOfCurrentGroup = finalColumnOfLastGroup + group[1]
		if beginningColumnOfCurrentGroup != LEFT_MOST_COLUMN {
			aisleColumns = append(aisleColumns, beginningColumnOfCurrentGroup)
		}
		if finalColumnOfCurrentGroup != m.rightMostColumn {
			aisleColumns = append(aisleColumns, finalColumnOfCurrentGroup)
		}
		finalColumnOfLastGroup = finalColumnOfCurrentGroup
	}
	return aisleColumns
}

func (m *SeatMap) computeWindowColumns() []int {
	return []int{LEFT_MOST_COLUMN, m.rightMostColumn}
}

func (m *SeatMap) computeMiddleColumns(windowColumns, aisleColumns []int) []int {
	middleColumns := make([]int, 0, 0)
	for i := LEFT_MOST_COLUMN + 1; i < windowColumns[1]; i++ {
		if !util.IsElementPresent(i, aisleColumns) {
			middleColumns = append(middleColumns, i)
		}
	}
	return middleColumns
}
