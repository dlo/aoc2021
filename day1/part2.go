package day1

import (
	"github.com/psampaz/slice"
	"log"
)

func Part2GenerateSlidingWindows(depths []int) []int {
	if len(depths) < 3 {
		return []int{}
	}

	sum, err := slice.SumInt(depths[0:3])
	if err != nil {
		log.Fatal(err)
	}
	measurements := []int{sum}
	for i, depth := range depths[3:] {
		sum += depth - depths[i]
		measurements = append(measurements, sum)
	}

	return measurements
}

//func Part2CountIncreasesInSlidingWindows(depths []int) int {
//	return Part1CountIncreases(Part2GenerateSlidingWindows(depths))
//}
