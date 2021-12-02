package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1CalculatePosition(t *testing.T) {
	var tests = []TestCase{
		{"example_input.txt", -150},
		{"input.txt", -2147104},
	}

	for _, tt := range tests {
		instructions := Part1ParseInstructions(tt.filename)
		position := Part1CalculatePosition(instructions)
		assert.Equal(t, tt.want, position.ProductOfCoordinates())
	}
}
