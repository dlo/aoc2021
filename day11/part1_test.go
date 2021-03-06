package day11

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlashCount(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 1656},
		{"testdata/input.txt", 1721},
	}

	for _, tt := range tests {
		grid := ReadOctopusGrid(tt.Filename)
		grid.StepByCount(100)
		assert.Equal(t, tt.Want, grid.flashCount)
	}
}
