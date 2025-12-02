package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges := parseInput()
	fmt.Println(part1(ranges), part2(ranges))
}

func parseInput() [][2]int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var result [][2]int
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			split := strings.Split(r, "-")
			min, _ := strconv.Atoi(split[0])
			max, _ := strconv.Atoi(split[1])
			result = append(result, [2]int{min, max})
		}
	}

	return result
}
