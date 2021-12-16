package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateCheapestPositionPart2(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 168},
		{"testdata/input.txt", 94813675},
	}

	for _, tt := range tests {
		positions := ParseHorizontalPositions(tt.filename)
		assert.Equal(t, tt.want, CalculateCheapestPosition(positions, Part2CalculateFuelUsageForDelta))
	}
}
