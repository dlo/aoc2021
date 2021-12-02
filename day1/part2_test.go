package day1

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Part2TestCase struct {
	measurements []int
	want []int
}

func TestPart2GenerateSlidingWindows(t *testing.T) {
	var tests = []Part2TestCase{
		{[]int{200, 210}, []int{}},
		{utils.LinesFromFile("part1_example_input.txt"), utils.LinesFromFile("part2_example_output.txt")},
		{utils.LinesFromFile("part1_input.txt"), utils.LinesFromFile("part2_output.txt")},
	}

	for _, tt := range tests {
		assert.ObjectsAreEqualValues(tt.want, Part2GenerateSlidingWindows(tt.measurements))
	}
}
