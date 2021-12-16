package day07

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateCheapestPositionPart2(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 168},
		{"testdata/input.txt", 94813675},
	}

	for _, tt := range tests {
		positions := ParseHorizontalPositions(tt.Filename)
		assert.Equal(t, tt.Want, CalculateCheapestPosition(positions, Part2CalculateFuelUsageForDelta))
	}
}
