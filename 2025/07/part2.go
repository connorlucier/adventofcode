package main

func part2(input [][]string, start int) int {
	paths, _ := traverse(input, start)
	result := 0
	for _, v := range paths {
		result += v
	}

	return result
}
