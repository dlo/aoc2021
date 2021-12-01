package day1

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1CalculateNumberOfIncreases(t *testing.T) {
	var tests = []struct {
		measurements []int
		want         int
	}{
		{[]int{200}, 0},
		{utils.LinesFromFile("day1_example_input.txt"), 7},
		{utils.LinesFromFile("day1_input.txt"), 1374},
	}

	for _, tt := range tests {
		if got := Part1CountIncreases(tt.measurements); got != tt.want {
			assert.Equal(t, tt.want, got)
		}
	}
}
