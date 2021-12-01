package day1

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	items := utils.LinesFromFile("day1_input.txt")

	want := 1374
	if got := Part1CalculateNumberOfIncreases(items); got != want {
		assert.Equal(t, want, got)
	}
}
