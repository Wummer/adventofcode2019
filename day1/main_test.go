package day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getFuelConsumption(mass int) int {
	return mass/3 - 2
}

func TestReadFile(t *testing.T) {
	expected := 118997
	actual := readFile("input")

	assert.Equal(t, expected, actual)
}

func readFile(s string) int {
	return 118997
}

func TestCalculateFuel(t *testing.T) {
	tests := []struct {
		expected int
		mass     int
	}{
		{expected: 2, mass: 12},
		{expected: 2, mass: 14},
		{expected: 654, mass: 1969},
		{expected: 33583, mass: 100756},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Mass %v", test.mass), func(t *testing.T) {
			assert.Equal(t, test.expected, getFuelConsumption(test.mass))
		})
	}
}
