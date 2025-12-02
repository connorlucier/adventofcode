package main

import "strconv"

func part1(input [][2]int) int {
	result := 0
	for _, r := range input {
		for i := r[0]; i <= r[1]; i++ {
			str := strconv.Itoa(i)
			len := len(str)
			lh := str[:len/2]
			rh := str[len/2:]

			if lh == rh {
				result += i
			}
		}
	}

	return result
}
