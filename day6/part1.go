package day6

import (
	"github.com/dlo/aoc2021/day4"
	"github.com/dlo/aoc2021/utils"
)

func CalculateSchoolSizeAfterNDays(school []int, days int) int {
	fishCounts := make([]int, 9)
	for _, age := range school {
		fishCounts[age]++
	}

	for i := 0; i < days; i++ {
		parentCount := fishCounts[0]
		fishCounts = fishCounts[1:]
		fishCounts[6] += parentCount
		fishCounts = append(fishCounts, parentCount)
	}

	return utils.SumSlice(fishCounts)
}

func ParseFishFile(filename string) []int {
	lines := utils.ReadLinesFromFile(filename)
	numbers := day4.ParseNumbersFromString(lines[0])
	var ages []int
	for _, number := range numbers {
		ages = append(ages, number)
	}
	return ages
}
