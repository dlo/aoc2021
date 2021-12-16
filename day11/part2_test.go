package day11

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstSynchronousStep(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 195},
		{"testdata/input.txt", 298},
	}

	for _, tt := range tests {
		grid := ReadOctopusGrid(tt.Filename)
		assert.Equal(t, tt.Want, grid.FirstSynchronousStep())
	}
}
