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

func TestSolution2(t *testing.T) {
	expected := 591
	actual := CountCombinationsStrict(284639, 748759)

	assert.Equal(t, expected, actual)
}

func TestHasDouble(t *testing.T) {
	input := 113456
	input2 := 123455

	assert.True(t, containsDouble(input))
	assert.True(t, containsDouble(input2))

}

func TestIsNotDecreasing(t *testing.T) {
	expectedHappy := 113456
	expectedUnhappy := 123436

	assert.True(t, isNotDecreasing(expectedHappy))
	assert.False(t, isNotDecreasing(expectedUnhappy))
}

func TestHasStrictlyDouble(t *testing.T) {
	expectedHappy := 112233
	expectedHappy2 := 111122
	expectedUnhappy := 123444

	assert.True(t, containsStrictDouble(expectedHappy))
	assert.True(t, containsStrictDouble(expectedHappy2))
	assert.False(t, containsStrictDouble(expectedUnhappy))
}
