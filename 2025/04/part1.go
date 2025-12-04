package main

func part1(grid [][]string) int {
	result := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != "@" {
				continue
			}

			adjacent := getAdjacent(grid, r, c)
			count := 0
			for _, adj := range adjacent {
				if adj == "@" {
					count++
				}
			}
			if count < 4 {
				result++
			}
		}
	}

	return result
}
