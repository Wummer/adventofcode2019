package day5

import "fmt"

func parseIntcode(intcode []int, input int) []int {
	var index int
	for index < len(intcode) {
		switch intcode[index] {
		case 99:
			return intcode
		case 1:
			intcode[intcode[index+3]] = intcode[intcode[index+1]] + intcode[intcode[index+2]]
			index += 4
		case 2:
			intcode[intcode[index+3]] = intcode[intcode[index+1]] * intcode[intcode[index+2]]
			index += 4
		case 3:
			intcode[intcode[index+1]] = input
			index += 2
		case 4:
			fmt.Println(intcode[intcode[index+1]])
			index += 2

		default:
			panic(fmt.Sprintf("unknown opcode %d at index %d", intcode[index], index))
		}

	}

	return intcode

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
			actual := parseIntcode(newArray, 1)
			if actual[0] == 19690720 {
				fmt.Println(100*noun + verb)
				return actual[0]
			}
		}
	}
	return actual
}
