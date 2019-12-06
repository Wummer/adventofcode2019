package day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FuelCounter struct {
	count int
}

func (fc *FuelCounter) Add(i int) int {
	fc.count += i
	return fc.count
}

func TestAdd(t *testing.T) {
	expected := 2
	fuelCounter := FuelCounter{}

	assert.Equal(t, expected, fuelCounter.Add(2))
}

func TestAddTable(t *testing.T) {
	fuelCounter := FuelCounter{}
	tests := []struct {
		input    int
		expected int
	}{
		{input: 2, expected: 2},
		{input: 2, expected: 4},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Adding %d", test.input), func(t *testing.T) {
			assert.Equal(t, test.expected, fuelCounter.Add(test.input))
		})
	}
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
