package seatmap

import (
	"airplane-seating/internal/seat"
	"airplane-seating/internal/seat/seattype"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeatMap_Initialise_ShouldReturnSeatsAccordingToLayout(t *testing.T) {
	tests := []struct {
		name     string
		layout   [][]int
		expected []*seat.Seat
	}{
		{
			name:   "2By2Layout",
			layout: [][]int{{2, 2}, {2, 2}},
			expected: []*seat.Seat{
				seat.NewSeat(seattype.WINDOW, 1, 1),
				seat.NewSeat(seattype.AISLE, 1, 2),
				seat.NewSeat(seattype.WINDOW, 2, 1),
				seat.NewSeat(seattype.AISLE, 2, 2),
				seat.NewSeat(seattype.AISLE, 1, 3),
				seat.NewSeat(seattype.WINDOW, 1, 4),
				seat.NewSeat(seattype.AISLE, 2, 3),
				seat.NewSeat(seattype.WINDOW, 2, 4),
			},
		},
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := NewSeatMap(test.layout).Initialise()
			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, actual)
		})
	}
}

func TestSeatMap_computeExtremities_ShouldReturnBackMostRowAndRightMostColumnNumbers(t *testing.T) {
	tests := []struct {
		name              string
		layout            [][]int
		expectedBackMost  int
		expectedRightMost int
		expectedErr       error
	}{
		{
			name:              "2By2Layout",
			layout:            [][]int{{2, 2}, {2, 2}},
			expectedBackMost:  2,
			expectedRightMost: 4,
		},

		{
			name:              "3 groups seatmap",
			layout:            [][]int{{2, 3}, {4, 2}, {2, 4}},
			expectedBackMost:  4,
			expectedRightMost: 9,
		},

		{
			name:              "4 groups seatmap",
			layout:            [][]int{{2, 3}, {4, 2}, {3, 2}, {2, 3}},
			expectedBackMost:  4,
			expectedRightMost: 10,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualBackMost, actualRightMost, actualErr := computeExtremities(test.layout)
			assert.NoError(t, actualErr)
			assert.Equal(t, test.expectedBackMost, actualBackMost)
			assert.Equal(t, test.expectedRightMost, actualRightMost)
		})
	}
}

func TestSeatMap_computeAisleColumns_ShouldReturnAisleColumnNumbers(t *testing.T) {
	tests := []struct {
		name                 string
		layout               [][]int
		rightMost            int
		expectedAisleColumns []int
	}{
		{
			name:                 "2By2Layout",
			layout:               [][]int{{2, 2}, {2, 2}},
			rightMost:            4,
			expectedAisleColumns: []int{2, 3},
		},

		{
			name:                 "3 groups seatmap",
			layout:               [][]int{{2, 3}, {4, 2}, {2, 4}},
			rightMost:            9,
			expectedAisleColumns: []int{3, 4, 5, 6},
		},

		{
			name:                 "4 groups seatmap",
			layout:               [][]int{{2, 3}, {4, 2}, {3, 2}, {2, 3}},
			rightMost:            10,
			expectedAisleColumns: []int{3, 4, 5, 6, 7, 8},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualAisleColumns := NewSeatMap(test.layout).computeAisleColumns()
			assert.Equal(t, test.expectedAisleColumns, actualAisleColumns)
		})
	}
}

func TestSeatMap_computeWindowColumns_ShouldReturnWindowColumnNumbers(t *testing.T) {
	tests := []struct {
		name                  string
		layout                Layout
		expectedWindowColumns []int
	}{
		{
			name:                  "2By2Layout",
			layout:                [][]int{{2, 2}, {2, 2}},
			expectedWindowColumns: []int{1, 4},
		},
		{
			name:                  "2By2Layout",
			layout:                [][]int{{2, 3}, {3, 3}, {2, 3}},
			expectedWindowColumns: []int{1, 9},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualWindowColumns := NewSeatMap(test.layout).computeWindowColumns()
			assert.Equal(t, test.expectedWindowColumns, actualWindowColumns)
		})
	}
}

func TestSeatMap_computeMiddleColumns_ShouldReturnMiddleColumnNumbers(t *testing.T) {
	tests := []struct {
		name                  string
		layout                Layout
		windowColumns         []int
		aisleColumns          []int
		expectedMiddleColumns []int
	}{
		{
			name:                  "no middle seats",
			layout:                [][]int{{2, 2}, {2, 2}},
			windowColumns:         []int{1, 4},
			aisleColumns:          []int{2, 3},
			expectedMiddleColumns: []int{},
		},
		{
			name:                  "3 column groups",
			layout:                [][]int{{2, 3}, {2, 3}},
			windowColumns:         []int{1, 6},
			aisleColumns:          []int{3, 4},
			expectedMiddleColumns: []int{2, 5},
		},
		{
			name:                  "multiple middle columns",
			layout:                [][]int{{2, 3}, {2, 2}, {2, 4}},
			windowColumns:         []int{1, 9},
			aisleColumns:          []int{3, 4, 5, 6},
			expectedMiddleColumns: []int{2, 7, 8},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualMiddleColumns := NewSeatMap(test.layout).computeMiddleColumns(test.windowColumns, test.aisleColumns)
			assert.Equal(t, test.expectedMiddleColumns, actualMiddleColumns)
		})
	}
}
