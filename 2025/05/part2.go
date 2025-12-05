package main

import (
	"slices"
)

func part2(fresh []Range) int {
	var grouped []Range
	for _, r := range getSorted(fresh) {
		index := findBy(grouped, r, func(a Range, b Range) bool { return a.min >= b.min && a.min <= b.max })
		if index >= 0 {
			if r.min < grouped[index].min {
				grouped[index].min = r.min
			}
			if r.max > grouped[index].max {
				grouped[index].max = r.max
			}
		} else {
			grouped = append(grouped, r)
		}
	}

	result := 0
	for _, g := range grouped {
		result += g.max - g.min + 1
	}
	return result
}

// sorting by Range.min eliminates cases where we don't know
// whether a range will overlap with another or not yet
func getSorted(values []Range) []Range {
	sorted := make([]Range, len(values))
	copy(sorted, values)
	slices.SortFunc(sorted, func(a Range, b Range) int { return a.min - b.min })
	return sorted
}

func findBy(values []Range, r Range, test func(a Range, b Range) bool) int {
	result := -1
	for i, v := range values {
		if test(r, v) {
			result = i
			break
		}
	}
	return result
}
