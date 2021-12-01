package day1

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2CalculateNumberOfIncreasesInSlidingWindows(t *testing.T) {
	var tests = []struct {
		filename string
		want     int
	}{
		{"day1_example_input.txt", 5},
		{"day1_input.txt", 1418},
	}

	for _, tt := range tests {
		items := utils.LinesFromFile(tt.filename)
		if got := Part2CalculateNumberOfIncreasesInSlidingWindows(items); got != tt.want {
			assert.Equal(t, tt.want, got)
		}
	}
}
