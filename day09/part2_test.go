package day09

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductOfSizesOfThreeLargestBasins(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 1134},
		{"testdata/input.txt", 1059300},
	}

	for _, tt := range tests {
		values := GenerateHeightMap(tt.Filename)

		basin := BasinHolder{}
		basin.Initialize(values)
		getBasinsResult := basin.GetBasins()
		assert.Equal(t, tt.Want, getBasinsResult.ProductOfSizesOfThreeLargestBasins())
	}
}
