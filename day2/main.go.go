package day2

func parseIntcode(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		if input[i] == 99 {
			return input
		}
		if input[i] == 1 {
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		} else {
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		}
	}

	return input

}

func swapValues(input []int, pos1, pos2 int) []int {
	result := input
	result[1] = pos1
	result[2] = pos2
	return result
}
