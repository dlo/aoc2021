package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Part1TestCase struct {
	filename string
	want     int
}


func TestPart1CalculateCoordinateFromCommands(t *testing.T) {
	var tests = []Part1TestCase{
		{"testdata/example_input.txt", -150},
		{"testdata/input.txt", -2147104},
	}

	for _, tt := range tests {
		commands := Part1ParseCommands(tt.filename)
		position := Part1CalculateCoordinateFromCommands(commands)
		assert.Equal(t, tt.want, position.ProductOfCoordinates())
	}
}
