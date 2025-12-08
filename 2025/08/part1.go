package main

import (
	"cmp"
	"slices"
)

func part1(input []Coord) int {
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
	for _, p := range pairs[:1000] {
		idx1, idx2 := find(p, circuits)
		if idx1 == idx2 {
			continue
		}
		circuits[idx1] = append(circuits[idx1], circuits[idx2]...)
		circuits = append(circuits[:idx2], circuits[idx2+1:]...)
	}

	slices.SortFunc(circuits, func(a []Coord, b []Coord) int { return len(b) - len(a) })
	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}
