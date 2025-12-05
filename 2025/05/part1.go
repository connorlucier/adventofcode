package main

func part1(fresh []Range, ingredients []int) int {
	result := 0
	for _, ingr := range ingredients {
		for _, r := range fresh {
			if r.min <= ingr && r.max >= ingr {
				result++
				break
			}
		}
	}
	return result
}
