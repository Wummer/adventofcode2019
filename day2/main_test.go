package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	expected := []int{2, 0, 0, 0, 99}

	assert.Equal(t, expected, parseIntcode(input))
}

func Test2(t *testing.T) {
	input := []int{2, 3, 0, 3, 99}
	expected := []int{2, 3, 0, 6, 99}

	assert.Equal(t, expected, parseIntcode(input))
}

func Test3(t *testing.T) {
	input := []int{2, 4, 4, 5, 99, 0}
	expected := []int{2, 4, 4, 5, 99, 9801}

	assert.Equal(t, expected, parseIntcode(input))
}

func Test4(t *testing.T) {
	input := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	assert.Equal(t, expected, parseIntcode(input))
}

func TestSolution(t *testing.T) {
	expected := []int{3716293, 12, 2, 2, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 48, 2, 19, 6, 96, 2, 13, 23, 480, 1, 9, 27, 483, 2, 31, 9, 1449, 1, 6, 35, 1451, 2, 10, 39, 5804, 1, 5, 43, 5805, 1, 5, 47, 5806, 2, 51, 6, 11612, 2, 10, 55, 46448, 1, 59, 9, 46451, 2, 13, 63, 232255, 1, 10, 67, 232259, 1, 71, 5, 232260, 1, 75, 6, 232262, 1, 10, 79, 232266, 1, 5, 83, 232267, 1, 5, 87, 232268, 2, 91, 6, 464536, 2, 6, 95, 929072, 2, 10, 99, 3716288, 1, 103, 5, 3716289, 1, 2, 107, 3716291, 1, 6, 111, 0, 99, 2, 14, 0, 0}
	input := []int{
		1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1,
		5, 0, 3, 2, 10, 1, 19, 2, 19, 6, 23, 2, 13, 23,
		27, 1, 9, 27, 31, 2, 31, 9, 35, 1, 6, 35, 39, 2,
		10, 39, 43, 1, 5, 43, 47, 1, 5, 47, 51, 2, 51, 6,
		55, 2, 10, 55, 59, 1, 59, 9, 63, 2, 13, 63, 67, 1,
		10, 67, 71, 1, 71, 5, 75, 1, 75, 6, 79, 1, 10, 79,
		83, 1, 5, 83, 87, 1, 5, 87, 91, 2, 91, 6, 95, 2,
		6, 95, 99, 2, 10, 99, 103, 1, 103, 5, 107,
		1, 2, 107, 111, 1, 6, 111, 0, 99, 2, 14, 0, 0,
	}
	assert.Equal(t, expected, parseIntcode(input))
}

func parseIntcode(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		if input[i] == 99 {
			return input
		}
		if input[i] == 1 {
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		} else {
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		}

	}

	return input

}
