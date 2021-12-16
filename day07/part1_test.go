package day07

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateCheapestPositionPart1(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 37},
		{"testdata/input.txt", 336040},
	}

	for _, tt := range tests {
		positions := ParseHorizontalPositions(tt.Filename)
		assert.Equal(t, tt.Want, CalculateCheapestPosition(positions, Part1CalculateFuelUsageForDelta))
	}
}
