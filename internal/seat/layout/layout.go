package layout

import (
	"airplane-seating/internal/seat"
	"fmt"
)

type Layout [][]int

func Initialise(layout Layout) ([]*seat.Seat, error) {
	backMostSeatRow, rightMostSeatColumn, _ := computeExtremities(layout)
	fmt.Printf("backMost -> %v\n rightMost -> %v\n", backMostSeatRow, rightMostSeatColumn) //todo: remove later

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
