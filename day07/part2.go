package day07

import "github.com/dlo/aoc2021/utils"

func Part2CalculateFuelUsageForDelta(start int, end int) int {
	difference := utils.Abs(start - end)
	return difference * (difference + 1) / 2
}
