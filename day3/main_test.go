package day3

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"strconv"
	"testing"
)

func TestManhattan(t *testing.T) {
	expected := 4
	start := [2]int{0, 0}
	end := [2]int{0, 4}
	var actual int
	actual = ManhattanDistance(start, end)

	assert.Equal(t, expected, actual)
}

func TestParsePath(t *testing.T) {
	tests := []struct {
		input    string
		expected [2]int
	}{
		{"R75", [2]int{75, 0}},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual, err := ParsePath(test.input)
			require.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func ParsePath(input string) ([2]int, error) {
	start := [2]int{0, 0}
	if input[0] == 'R' {
		x, err := strconv.Atoi(input[1:])
		if err != nil {
			return [2]int{0, 0}, err
		}
		start[0] += x

	} else if input[0] == 'L' {
		x, err := strconv.Atoi(input[1:])
		if err != nil {
			return [2]int{0, 0}, err
		}
		start[0] -= x
	}
	return start, nil
}

func ManhattanDistance(x [2]int, y [2]int) int {
	result := 0.
	for i := 0; i < 2; i++ {
		result += math.Abs(float64(x[i]) - float64(y[i]))
	}
	return int(result)
}
