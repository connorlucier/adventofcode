package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fresh, ingredients := parseInput()
	fmt.Println(part1(fresh, ingredients), part2(fresh))
}

func parseInput() ([]Range, []int) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var fresh []Range
	var ingredients []int

	isFreshRange := true
	for scanner.Scan() {
		line := scanner.Text()

		// detect second half of input
		if len(line) == 0 {
			isFreshRange = false
			continue
		}

		if isFreshRange {
			strs := strings.Split(line, "-")
			min, _ := strconv.Atoi(strs[0])
			max, _ := strconv.Atoi(strs[1])
			fresh = append(fresh, Range{min, max})
		} else {
			num, _ := strconv.Atoi(line)
			ingredients = append(ingredients, num)
		}
	}

	return fresh, ingredients
}

type Range struct {
	min int
	max int
}
