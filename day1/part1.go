package day1

func Part1CalculateNumberOfIncreases(items []int) int {
	previous := -1
	increments := -1
	for _, value := range items {
		if value > previous {
			increments++
		}
		previous = value
	}

	return increments
}
