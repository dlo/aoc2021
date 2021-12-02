package day1

import (
	"github.com/dlo/aoc2021/utils"
)

func Part2GenerateSlidingWindows(depths []int) []int {
	if len(depths) < 3 {
		return []int{}
	}

	sum := utils.SumSlice(depths[0:3])
	measurements := []int{sum}
	for i, depth := range depths[3:] {
		sum += depth - depths[i]
		measurements = append(measurements, sum)
	}

	return measurements
}

func Part2CountIncreasesInSlidingWindows(depths []int) int {
	return Part1CountIncreases(Part2GenerateSlidingWindows(depths))
}
