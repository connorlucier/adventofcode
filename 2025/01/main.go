package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var input = parseInput()
	fmt.Println(part1(input), part2(input))
}

func parseInput() []int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var result []int
	for scanner.Scan() {
		dir := scanner.Text()[0]
		val, _ := strconv.Atoi(scanner.Text()[1:])

		if dir == 'L' {
			val *= -1
		}

		result = append(result, val)
	}

	return result
}

// wrote this to handle negative numbers more elegantly
// since the Go % operator doesn't work how I imagined
func mod(a int, b int) int {
	return (a%b + b) % b
}

// using this with 'b'=100 effectively chops off all digits
// except the tens and ones places regardless of whether 'a'
// is positive or negative
func rem(a int, b int) int {
	if int(math.Abs(float64(a))) <= b {
		return a
	}

	rem := mod(int(math.Abs(float64(a))), b)
	if a < 0 {
		rem *= -1
	}
	return rem
}
