package day02

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1CalculateCoordinateFromCommands(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", -150},
		{"testdata/input.txt", -2147104},
	}

	for _, tt := range tests {
		commands := Part1ParseCommands(tt.Filename)
		position := Part1CalculateCoordinateFromCommands(commands)
		assert.Equal(t, tt.Want, position.ProductOfCoordinates())
	}
}
