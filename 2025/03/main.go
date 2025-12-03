package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput()
	fmt.Println(part1(input), part2(input))
}

type IndexValueTuple struct {
	index int
	value int
}

func parseInput() [][]int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var result [][]int
	for scanner.Scan() {
		var line []int
		for _, v := range strings.Split(scanner.Text(), "") {
			num, _ := strconv.Atoi(v)
			line = append(line, num)
		}
		result = append(result, line)
	}

	return result
}

func getMax(values []int, start int, end int) IndexValueTuple {
	max := 0 // safe since input values > 0
	idx := 0
	for i := start; i < end; i++ {
		// '>' since want the leftmost max within the range
		if values[i] > max {
			max = values[i]
			idx = i
		}
	}

	return IndexValueTuple{idx, max}
}
