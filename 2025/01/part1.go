package main

func part1(values []int) int {
	current := 50
	result := 0
	for _, v := range values {
		current = mod(current+v, 100)
		if current == 0 {
			result += 1
		}
	}

	return result
}
