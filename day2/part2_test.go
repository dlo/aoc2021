package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
type Part2TestCase struct {
	filename string
	want     int
}

func TestPart2CalculateCoordinateFromCommands(t *testing.T) {
	var tests = []Part2TestCase{
		{"testdata/example_input.txt", 900},
		{"testdata/input.txt", 2044620088},
	}

	for _, tt := range tests {
		commands := Part1ParseCommands(tt.filename)
		position := Part2CalculateCoordinateFromCommands(commands)
		assert.Equal(t, tt.want, position.ProductOfCoordinates())
	}
}
