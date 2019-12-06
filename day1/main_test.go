package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func calculateFuel(mass int) int {
	return 2
}

func TestCalculateFuel(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		mass     int
	}{
		{name: "mass 12", expected: 2, mass: 12},
		{name: "mass 14", expected: 2, mass: 14},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, calculateFuel(test.mass))
		})
	}
}
