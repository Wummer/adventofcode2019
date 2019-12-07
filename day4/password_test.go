package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution1(t *testing.T) {
	expected := 895
	actual := CountCombinations(284639, 748759)

	assert.Equal(t, expected, actual)
}

func TestHasDouble(t *testing.T) {
	input := 113456
	input2 := 123455

	assert.True(t, hasDouble(input))
	assert.True(t, hasDouble(input2))

}

func TestIsNotDecreasing(t *testing.T) {
	expectedHappy := 113456
	expectedUnhappy := 123436

	assert.True(t, isNotDecreasing(expectedHappy))
	assert.False(t, isNotDecreasing(expectedUnhappy))
}
