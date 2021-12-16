package day13

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountDots(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 17},
		{"testdata/input.txt", 751},
	}

	for _, tt := range tests {
		puzzle := ImportGrid(tt.Filename)
		assert.Equal(t, tt.Want, puzzle.Fold().CountDots())
	}
}
