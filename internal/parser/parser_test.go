package parser

import (
	"airplane-seating/internal/seat/seatmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_ParseInput_ShouldReturnStringAs2DIntArray(t *testing.T) {
	expected := seatmap.Layout{{2, 2}, {4, 3}}

	actual, err := convertStringToLayout("[[2,2], [3, 4]]")

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func Test_separateIndividual1DArrayElements(t *testing.T) {
	assert.Equal(t, []string{"2,2"}, separateIndividual1DArrayElements("[[2,2]]"))
	assert.Equal(t, []string{"2,2", "3,3"}, separateIndividual1DArrayElements("[[2,2], [3,3]]"))
	assert.Equal(t, []string{"2,2", "3,3", "4,4"}, separateIndividual1DArrayElements("[[2,2], [3,3], [4,4]]"))
	assert.Equal(t, []string{"2,2", "3,3", "4,4"}, separateIndividual1DArrayElements("[[2,2],[3,3],[4,4]]"))
}
