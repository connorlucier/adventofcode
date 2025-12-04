package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := parseInput()
	fmt.Println(part1(input), part2(input))
}

func parseInput() [][]string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var result [][]string
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ""))
	}

	return result
}

func getAdjacent(grid [][]string, r int, c int) []string {
	var result []string

	// starting at up-left, moving clockwise
	if r > 0 {
		if c > 0 {
			result = append(result, grid[r-1][c-1])
		}
		result = append(result, grid[r-1][c])
		if c < len(grid[r])-1 {
			result = append(result, grid[r-1][c+1])
		}
	}
	if c < len(grid)-1 {
		result = append(result, grid[r][c+1])
	}
	if r < len(grid)-1 {
		if c < len(grid[r])-1 {
			result = append(result, grid[r+1][c+1])
		}
		result = append(result, grid[r+1][c])
		if c > 0 {
			result = append(result, grid[r+1][c-1])
		}
	}
	if c > 0 {
		result = append(result, grid[r][c-1])
	}

	return result
}
