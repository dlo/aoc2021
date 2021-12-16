package day12

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountUniquePaths(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input_1.txt", 10},
		{"testdata/example_input_2.txt", 19},
		{"testdata/example_input_3.txt", 226},
		{"testdata/input.txt", 3708},
	}

	for _, tt := range tests {
		caveMap := ImportCaveMap(tt.Filename)
		assert.Equal(t, tt.Want, caveMap.CountUniquePaths(false))
	}
}

func TestCountUniquePathsWithRepeatVisitsToSmallCaves(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input_1.txt", 36},
		{"testdata/example_input_2.txt", 103},
		{"testdata/example_input_3.txt", 3509},
		{"testdata/input.txt", 93858},
	}

	for _, tt := range tests {
		caveMap := ImportCaveMap(tt.Filename)
		assert.Equal(t, tt.Want, caveMap.CountUniquePaths(true))
	}
}
