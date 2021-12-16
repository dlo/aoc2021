package day09

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductOfSizesOfThreeLargestBasins(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 1134},
		{"testdata/input.txt", 1059300},
	}

	for _, tt := range tests {
		values := GenerateHeightMap(tt.filename)

		basin := BasinHolder{}
		basin.Initialize(values)
		getBasinsResult := basin.GetBasins()
		assert.Equal(t, tt.want, getBasinsResult.ProductOfSizesOfThreeLargestBasins())
	}
}
