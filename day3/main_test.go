package day3

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestManhattan(t *testing.T) {
	expected := 4.
	start := [2]float64{0, 0}
	end := [2]float64{0, 4}
	var actual float64
	actual = ManhattanDistance(start, end)
	assert.Equal(t, expected, actual)
}

func ManhattanDistance(x [2]float64, y [2]float64) float64 {
	result := 0.
	for i := 0; i < 2; i++ {
		result += math.Abs(x[i] - y[i])
	}
	return result
}
