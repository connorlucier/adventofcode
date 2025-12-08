package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput()
	fmt.Println(part1(input), part2(input))
}

type Coord struct {
	id int
	x  int
	y  int
	z  int
}

type CoordPair struct {
	coords   [2]Coord
	distance float64
}

func parseInput() []Coord {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var result []Coord
	id := 1
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(strs[0])
		y, _ := strconv.Atoi(strs[1])
		z, _ := strconv.Atoi(strs[2])
		result = append(result, Coord{id, x, y, z})
		id++
	}

	return result
}

func find(p CoordPair, coords [][]Coord) (int, int) {
	c1 := p.coords[0]
	c2 := p.coords[1]

	idx1 := -1
	idx2 := -1
	for i, circ := range coords {
		for _, coord := range circ {
			if coord.id == c1.id {
				idx1 = i
			}
			if coord.id == c2.id {
				idx2 = i
			}
		}
	}

	return idx1, idx2
}

func distance(a Coord, b Coord) float64 {
	dx := math.Pow(float64(b.x-a.x), 2)
	dy := math.Pow(float64(b.y-a.y), 2)
	dz := math.Pow(float64(b.z-a.z), 2)
	return math.Sqrt(dx + dy + dz)
}
