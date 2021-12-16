package day01

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Part2TestCase struct {
	measurements []int
	want         []int
}

func TestPart2GenerateSlidingWindows(t *testing.T) {
	var tests = []Part2TestCase{
		{[]int{200, 210}, []int{}},
		{utils.IntegerLinesFromFile("part1_example_input.txt"), utils.IntegerLinesFromFile("part2_example_output.txt")},
		{utils.IntegerLinesFromFile("part1_input.txt"), utils.IntegerLinesFromFile("part2_output.txt")},
	}

	for _, tt := range tests {
		assert.ObjectsAreEqualValues(tt.want, Part2GenerateSlidingWindows(tt.measurements))
	}
}
