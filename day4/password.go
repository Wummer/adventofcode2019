package day4

import (
	"fmt"
	"strings"
)

func CountCombinations(lower int, upper int) int {
	count := 0
	for i := lower; i < upper; i++ {
		if containsDouble(i) && isNotDecreasing(i) {
			count++
		}
	}
	return count
}
func CountCombinationsStrict(lower int, upper int) int {
	count := 0
	for i := lower; i < upper; i++ {
		if containsStrictDouble(i) && isNotDecreasing(i) {
			count++
		}
	}
	return count
}

func isNotDecreasing(input int) bool {
	stringInput := strings.Split(fmt.Sprintf("%d", input), "")

	for i := 1; i < len(stringInput); i++ {
		if stringInput[i] < stringInput[i-1] {
			return false
		}
	}
	return true
}

func containsDouble(input int) bool {
	stringInput := fmt.Sprintf("%d", input)
	previous := stringInput[0]
	for i := 1; i < len(stringInput); i++ {
		if stringInput[i] == previous {
			return true
		}
		previous = stringInput[i]
	}
	return false
}

func containsStrictDouble(input int) bool {
	stringInput := fmt.Sprintf("%d", input)
	previous := stringInput[0]
	groupCount := 1
	hasDouble := false
	for i := 1; i < len(stringInput); i++ {
		if stringInput[i] == previous {
			groupCount++

		} else {
			if groupCount == 2 {
				hasDouble = true
			}
			groupCount = 1
		}
		previous = stringInput[i]
	}
	return hasDouble || groupCount == 2
}
