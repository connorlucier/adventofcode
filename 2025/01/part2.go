package main

import (
	"math"
)

func part2(values []int) int {
	current := 50
	result := 0
	for _, v := range values {
		clicks := int(math.Abs(float64(v))) / 100
		rem := rem(v, 100)

		if v < 0 && int(math.Abs(float64(rem))) >= current && current != 0 {
			clicks += 1
		} else if v > 0 && current+rem >= 100 {
			clicks += 1
		}

		current = mod(current+v, 100)
		result += clicks
	}

	return result
}
