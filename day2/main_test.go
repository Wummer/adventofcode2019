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

func parseIntcode(input []int) []int {
	result := input

	result[0] = input[input[1]] + input[input[2]]

	return result

}
