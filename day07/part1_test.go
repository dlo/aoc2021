package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	want     int
}

func TestCalculateCheapestPositionPart1(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 37},
		{"testdata/input.txt", 336040},
	}

	for _, tt := range tests {
		positions := ParseHorizontalPositions(tt.filename)
		assert.Equal(t, tt.want, CalculateCheapestPosition(positions, Part1CalculateFuelUsageForDelta))
	}
}
