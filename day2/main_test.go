package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	expected := []int{2, 0, 0, 0, 99}

	assert.Equal(t, expected, intcode(input))
}

func intcode(input []int) []int {
	return []int{}

}
