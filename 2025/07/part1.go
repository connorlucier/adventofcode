package main

func part1(input [][]string, start int) int {
	_, splits := traverse(input, start)
	return splits
}
