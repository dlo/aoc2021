package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2CalculatePosition(t *testing.T) {
	var tests = []TestCase{
		{"example_input.txt", 900},
		{"input.txt", 2044620088},
	}

	for _, tt := range tests {
		instructions := Part1ParseInstructions(tt.filename)
		position := Part2CalculatePosition(instructions)
		assert.Equal(t, tt.want, position.ProductOfCoordinates())
	}
}
