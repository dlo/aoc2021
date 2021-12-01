package day1

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2CalculateNumberOfIncreasesInSlidingWindows(t *testing.T) {
	var tests = []struct {
		measurements []int
		want         int
	}{
		{[]int{200, 210}, 0},
		{utils.LinesFromFile("day1_example_input.txt"), 5},
		{utils.LinesFromFile("day1_input.txt"), 1418},
	}

	for _, tt := range tests {
		if got := Part2CountIncreasesInSlidingWindows(tt.measurements); got != tt.want {
			assert.Equal(t, tt.want, got)
		}
	}
}
