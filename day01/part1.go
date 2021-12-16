package day1

func Part1CountIncreases(depths []int) int {
	if len(depths) < 2 {
		return 0
	}

	count := 0
	previousDepth := depths[0]
	for _, depth := range depths[1:] {
		if depth > previousDepth {
			count++
		}
		previousDepth = depth
	}

	return count
}
