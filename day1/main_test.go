package day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FuelCounter struct {
	count int
}

func (fc FuelCounter) Add(i int) int {
	return fc.count + 2
}

func TestAdd(t *testing.T) {
	expected := 2
	fuelCounter := FuelCounter{}

	assert.Equal(t, expected, fuelCounter.Add(2))
}

func getFuelConsumption(mass int) int {
	return mass/3 - 2
}

func TestCalculateFuelConsumption(t *testing.T) {
	tests := []struct {
		expected int
		mass     int
	}{
		{expected: 2, mass: 12},
		{expected: 2, mass: 14},
		{expected: 654, mass: 1969},
		{expected: 33583, mass: 100756},
	}
	t.Parallel()
	for _, test := range tests {
		t.Run(fmt.Sprintf("Mass %v", test.mass), func(t *testing.T) {
			assert.Equal(t, test.expected, getFuelConsumption(test.mass))
		})
	}
}
