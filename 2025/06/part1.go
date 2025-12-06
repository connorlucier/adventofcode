package main

func part1(problems []Problem) int {
	result := 0
	for _, p := range problems {
		result += getResult(p)
	}
	return result
}
