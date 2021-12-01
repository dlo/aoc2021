package day1

func Part1CalculateNumberOfIncreases(items []int) int {
	increments := 0
	previous := items[0]
	for _, value := range items[1:] {
		if value > previous {
			increments++
		}
		previous = value
	}

	return increments
}
