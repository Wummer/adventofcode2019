package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseIntcode(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Case1", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{"Case2", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"Case3", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"Case4", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	t.Parallel()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, parseIntcode(test.input, 1))
		})
	}
}

func TestSwapValues(t *testing.T) {
	expected := []int{1, 0, 0, 0, 99}
	actual := swapValues([]int{1, 2, 2, 0, 99}, 0, 0)

	assert.Equal(t, expected, actual)
}
