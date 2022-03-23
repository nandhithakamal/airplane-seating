package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsElementPresent_ShouldReturnFalseIfElementIsNotPresent(t *testing.T) {
	zero := 0
	first5NaturalNumbers := []int{1, 2, 3, 4, 5}

	assert.False(t, IsElementPresent(zero, first5NaturalNumbers))
}

func Test_IsElementPresent_ShouldReturnTrueIfElementIsPresent(t *testing.T) {
	four := 4
	first5NaturalNumbers := []int{1, 2, 3, 4, 5}

	assert.True(t, IsElementPresent(four, first5NaturalNumbers))
}
