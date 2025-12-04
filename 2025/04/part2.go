package main

func part2(grid [][]string) int {
	result := 0
	for {
		diff := 0
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
					diff++
					grid[r][c] = "."
				}
			}
		}
		if diff == 0 {
			break
		}
	}

	return result
}
