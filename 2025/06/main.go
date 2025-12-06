package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput()
	fmt.Println(part1(input), part2(input))
}

type Problem struct {
	values    []int
	operator  string
	alignment string
}

func getResult(p Problem) int {
	result := p.values[0]
	for _, v := range p.values[1:] {
		if p.operator == "*" {
			result *= v
		} else {
			result += v
		}
	}
	return result
}

func parseInput() []Problem {
	bytes, _ := os.ReadFile("input.txt")
	lines := strings.FieldsFunc(string(bytes), func(c rune) bool { return c == '\n' || c == '\r' })

	var result []Problem
	next := Problem{}
	isNew := true
	strValues := make([]string, len(lines)-1)
	for col := 0; col < len(lines[0]); col++ {
		isBlank := true
		hasSpace := false
		for i, line := range lines {
			if line[col] != ' ' {
				if line[col] == '+' || line[col] == '*' {
					next.operator = string(line[col])
				} else {
					strValues[i] += string(line[col])
				}
				isBlank = false
			} else if isNew {
				hasSpace = true
			}
		}

		if isNew {
			isNew = false
			if hasSpace {
				next.alignment = "R"
			} else {
				next.alignment = "L"
			}
		}
		if isBlank || col == len(lines[0])-1 {
			next.values = make([]int, len(strValues))
			for i, str := range strValues {
				next.values[i], _ = strconv.Atoi(str)
			}

			result = append(result, next)
			next = Problem{}
			isNew = true
			strValues = make([]string, len(lines)-1)
		}
	}

	return result
}
