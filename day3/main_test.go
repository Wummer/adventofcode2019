package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnd2End(t *testing.T) {
	tests := []struct {
		wire1Path string
		name      string
		wire2Path string
		expected  int
	}{
		{"Example1", "R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"Example2", "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			coordinatesWire1 := FindCoordinates(test.wire1Path)
			t.Log(coordinatesWire1)
			coordinatesWire2 := FindCoordinates(test.wire2Path)
			t.Log(coordinatesWire2)

			actual := CalculateClosestIntersection(coordinatesWire1, coordinatesWire2)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestManhattan(t *testing.T) {
	expected := 4
	end := [2]int{0, 4}
	var actual int
	actual = ManhattanDistance(end)

	assert.Equal(t, expected, actual)
}

func TestDrawPath(t *testing.T) {
	expected := [][2]int{{0, 0}, {1, 0}, {2, 0}}
	input := "R2"

	actual := DrawPath([2]int{0, 0}, input)
	assert.Equal(t, expected, actual)
}
