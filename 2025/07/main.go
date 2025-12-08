package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, start := parseInput()
	fmt.Println(part1(input, start), part2(input, start))
}

func parseInput() ([][]string, int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var result [][]string
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ""))
	}

	start := 0
	for i, v := range result[0] {
		if v == "S" {
			start = i
		}
	}

	return result, start
}

func traverse(input [][]string, start int) (map[int]int, int) {
	paths := map[int]int{}
	paths[start] = 1
	splits := 0
	for _, row := range input {
		newVals := map[int]int{}
		for k, v := range paths {
			if row[k] == "^" {
				splits++
				if k > 0 {
					newVals[k-1] += v
				}
				if k < len(row)-1 {
					newVals[k+1] += v
				}
			} else {
				newVals[k] += v
			}
		}
		paths = newVals
	}

	return paths, splits
}
