package layout

import (
	"airplane-seating/internal/seat"
	"airplane-seating/internal/seat/seatType"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Initialise_ShouldReturnSeatsAccordingToLayout(t *testing.T) {
	tests := []struct {
		name     string
		layout   [][]int
		expected []*seat.Seat
	}{
		{
			"2By2Layout",
			[][]int{{2, 2}, {2, 2}},
			[]*seat.Seat{
				seat.NewSeat(seatType.WINDOW, 1, 1),
				seat.NewSeat(seatType.AISLE, 1, 2),
				seat.NewSeat(seatType.AISLE, 1, 3),
				seat.NewSeat(seatType.WINDOW, 1, 4),
				seat.NewSeat(seatType.WINDOW, 2, 1),
				seat.NewSeat(seatType.AISLE, 2, 2),
				seat.NewSeat(seatType.AISLE, 2, 3),
				seat.NewSeat(seatType.WINDOW, 2, 4),
			},
		},
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := Initialise(test.layout)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_computeExtremities_ShouldReturnBackMostRowAndRightMostColumnNumbers(t *testing.T) {
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
			name:              "3 groups layout",
			layout:            [][]int{{2, 3}, {4, 2}, {2, 4}},
			expectedBackMost:  4,
			expectedRightMost: 9,
		},

		{
			name:              "4 groups layout",
			layout:            [][]int{{2, 3}, {4, 2}, {3, 2}, {2, 3}},
			expectedBackMost:  4,
			expectedRightMost: 10,
		},

		// TODO: test cases
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
