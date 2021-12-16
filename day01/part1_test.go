package day01

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Part1TestCase struct {
	measurements []int
	want int
}

func TestPart1CountIncreases(t *testing.T) {
	var tests = []Part1TestCase{
		{[]int{200}, 0},
		{utils.IntegerLinesFromFile("testdata/part1_example_input.txt"), 7},
		{utils.IntegerLinesFromFile("testdata/part1_input.txt"), 1374},
		{utils.IntegerLinesFromFile("testdata/part2_example_output.txt"), 5},
		{utils.IntegerLinesFromFile("testdata/part2_output.txt"), 1418},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, Part1CountIncreases(tt.measurements))
	}
}
