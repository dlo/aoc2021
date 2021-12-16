package day07

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/psampaz/slice"
)

func ParseHorizontalPositions(filename string) []int {
	lines := utils.ReadLinesFromFile(filename)
	numbers := utils.ParseNumbersFromString(lines[0])
	var positions []int
	for _, number := range numbers {
		positions = append(positions, number)
	}
	return positions
}

func Part1CalculateFuelUsageForDelta(start int, end int) int {
	return utils.Abs(start - end)
}

func CalculateFuelForPosition(positions []int, target int, measurement func(int, int) int) int {
	sum := 0
	for _, position := range positions {
		sum += measurement(position, target)
	}
	return sum
}

func CalculateCheapestPosition(positions []int, measurement func(int, int) int) int {
	min, _ := slice.MinInt(positions)
	max, _ := slice.MaxInt(positions)
	minFuel := CalculateFuelForPosition(positions, max, measurement)
	for i := min; i < max; i++ {
		candidate := CalculateFuelForPosition(positions, i, measurement)
		if candidate < minFuel {
			minFuel = candidate
		}
	}
	return minFuel
}
