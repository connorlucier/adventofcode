package main

import (
	"strconv"
	"strings"
)

func part2(problems []Problem) int {
	result := 0
	for _, p := range transform(problems) {
		result += getResult(p)
	}
	return result
}

func transform(problems []Problem) []Problem {
	var result []Problem
	for _, p := range problems {
		// convert to strings to save headaches
		strs := make([]string, len(p.values))
		maxLen := 0
		for i, v := range p.values {
			strs[i] = strconv.Itoa(v)
			if len(strs[i]) > maxLen {
				maxLen = len(strs[i])
			}
		}

		// pad with spaces for left/right alignment
		for i := range strs {
			if len(strs[i]) < maxLen {
				pad := strings.Repeat(" ", maxLen-len(strs[i]))
				if p.alignment == "L" {
					strs[i] = strs[i] + pad
				} else {
					strs[i] = pad + strs[i]
				}
			}
		}

		// do the actual transformation
		newStrs := make([]string, maxLen)
		for i := maxLen - 1; i >= 0; i-- {
			for _, str := range strs {
				if str[i] != ' ' {
					newStrs[i] += string(str[i])
				}
			}
		}

		// convert back to ints
		newVals := make([]int, len(newStrs))
		for i, str := range newStrs {
			newVals[i], _ = strconv.Atoi(str)
		}

		result = append(result, Problem{newVals, p.operator, p.alignment})
	}

	return result
}
