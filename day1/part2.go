package day1

func Part2GenerateThreeMeasurementSlidingWindows(items []int) []int {
	if len(items) < 3 {
		return []int{}
	}

	sum := items[0] + items[1] + items[2]
	values := []int{sum}

	for i, value := range items[3:] {
		sum += value - items[i]
		values = append(values, sum)
	}

	return values
}

func Part2CalculateNumberOfIncreasesInSlidingWindows(items []int) int {
	return Part1CalculateNumberOfIncreases(Part2GenerateThreeMeasurementSlidingWindows(items))
}
