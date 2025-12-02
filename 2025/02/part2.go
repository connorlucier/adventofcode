package main

import (
	"strconv"
)

func part2(input [][2]int) int {
	result := 0
	for _, r := range input {
		for i := r[0]; i <= r[1]; i++ {
			str := strconv.Itoa(i)
			for j := len(str) / 2; j >= 1; j-- {
				// moot to test if the string doesn't divide
				// evenly into substrings of length j
				if len(str)%j != 0 {
					continue
				}

				substrings := getSubstrings(str, j)
				test := substrings[0]
				invalid := true

				for _, substr := range substrings[1:] {
					if substr != test {
						invalid = false
						break
					}
				}

				if invalid {
					result += i
					break
				}
			}
		}
	}

	return result
}

// assumes len(str) % length == 0
func getSubstrings(str string, length int) []string {
	var result []string
	for i := 0; i < len(str); i += length {
		result = append(result, str[i:i+length])
	}

	return result
}
