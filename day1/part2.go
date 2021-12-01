package day1

func Part2GenerateThreeMeasurementSlidingWindows(items []int) []int {
	sum := 0
	var values []int
	for i, value := range items {
		sum += value

		if i < 2 {
			continue
		}

		if i > 2 {
			sum -= items[i-3]
		}

		values = append(values, sum)
	}

	return values
}

func Part2CalculateNumberOfIncreasesInSlidingWindows(items []int) int {
	return Part1CalculateNumberOfIncreases(Part2GenerateThreeMeasurementSlidingWindows(items))
}
