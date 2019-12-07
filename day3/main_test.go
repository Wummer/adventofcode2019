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
	expected := [2]int{75, 0}
	start := [2]int{0, 0}
	input := "R75"
	actual, err := ParsePath(start, input)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)

}

func ParsePath(start [2]int, input string) ([2]int, error) {
	if input[0] == 'R' {
		num, err := strconv.Atoi(input[1:])
		if err != nil {
			return [2]int{0, 0}, err
		}
		start[0] = num

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
