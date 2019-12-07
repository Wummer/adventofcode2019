package day5

import "fmt"

func parseIntcode(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case 99:
			return input
		case 1:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		case 2:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		default:
			panic(fmt.Sprintf("unknown opcode %d at index %d", input[i], i))
		}

	}

	return input

}

func swapValues(input []int, pos1, pos2 int) []int {
	result := make([]int, len(input))
	copy(result, input)
	result[1] = pos1
	result[2] = pos2
	return result
}

func solve(input []int) int {
	var actual int
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			newArray := swapValues(input, noun, verb)
			actual := parseIntcode(newArray)
			if actual[0] == 19690720 {
				fmt.Println(100*noun + verb)
				return actual[0]
			}
		}
	}
	return actual
}
