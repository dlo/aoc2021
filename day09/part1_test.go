package day09

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowPoints(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 15},
		{"testdata/input.txt", 486},
	}

	for _, tt := range tests {
		values := GenerateHeightMap(tt.Filename)
		points := values.LowPoints()
		assert.Equal(t, tt.Want, points.TotalRisk())
	}
}
