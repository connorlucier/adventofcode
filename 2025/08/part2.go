package main

import (
	"cmp"
	"slices"
)

func part2(input []Coord) int {
	circuits := make([][]Coord, len(input))
	for i, coord := range input {
		circuits[i] = []Coord{coord}
	}

	var pairs []CoordPair
	for i, c1 := range input {
		for _, c2 := range input[i+1:] {
			pairs = append(pairs, CoordPair{[2]Coord{c1, c2}, distance(c1, c2)})
		}
	}

	slices.SortFunc(pairs, func(a CoordPair, b CoordPair) int { return cmp.Compare(a.distance, b.distance) })
	i := 0
	for len(circuits) > 1 && i < len(pairs) {
		idx1, idx2 := find(pairs[i], circuits)
		i++
		if idx1 == idx2 {
			continue
		}
		circuits[idx1] = append(circuits[idx1], circuits[idx2]...)
		circuits = append(circuits[:idx2], circuits[idx2+1:]...)
	}

	return pairs[i-1].coords[0].x * pairs[i-1].coords[1].x
}
