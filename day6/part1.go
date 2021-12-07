package day6

import (
	"github.com/dlo/aoc2021/day4"
	"github.com/dlo/aoc2021/utils"
)

func CalculateSchoolSizeAfterNDays(school []int, days int) int {
	counts := make([]int, 9)
	for _, age := range school {
		counts[age]++
	}

	for i := 0; i < days; i++ {
		parentCount := counts[0]
		counts = counts[1:]
		counts[6] += parentCount
		counts = append(counts, parentCount)
	}

	return utils.SumSlice(counts)
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
