package day06

import (
	"github.com/dlo/aoc2021/utils"
)

func CalculateSchoolSizeAfterNDays(timers []int, days int) int {
	counts := make([]int, 9)
	for _, timer := range timers {
		counts[timer]++
	}

	for i := 0; i < days; i++ {
		parentCount := counts[0]
		counts = counts[1:]
		counts[6] += parentCount
		counts = append(counts, parentCount)
	}

	return utils.SliceSum(counts)
}

func ParseFishFile(filename string) []int {
	lines := utils.ReadLinesFromFile(filename)
	numbers := utils.ParseNumbersFromString(lines[0])
	var ages []int
	for _, number := range numbers {
		ages = append(ages, number)
	}
	return ages
}
