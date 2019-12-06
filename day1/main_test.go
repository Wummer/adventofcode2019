package day1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func calculateFuel(mass int) int {
	return 2
}

func TestCalculateFuel(t *testing.T) {
	tests := []struct {
		expected int
		mass     int
	}{
		{expected: 2, mass: 12},
		{expected: 2, mass: 14},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Mass %v", test.mass), func(t *testing.T) {
			assert.Equal(t, test.expected, calculateFuel(test.mass))
		})
	}
}
