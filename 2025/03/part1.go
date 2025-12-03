package main

func part1(input [][]int) int {
	result := 0
	for _, row := range input {
		max := getMax(row, 0, len(row)-1)
		max2 := getMax(row, max.index+1, len(row))
		result += 10*max.value + max2.value
	}

	return result
}
