package main

func part2(input [][]int) int {
	result := 0
	for _, row := range input {
		start := getMax(row, 0, len(row)-12)
		value := start.value
		skips := 0
		maxSkips := len(row) - 12 - start.index

		for i := 1; i < 12; i++ {
			nextMax := getMax(row, start.index+skips+i, start.index+maxSkips+i+1)
			skips = nextMax.index - start.index - i
			value = value*10 + nextMax.value
		}
		result += value
	}

	return result
}
