package day08

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountOneFourSevenEights(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 26},
		{"testdata/input.txt", 310},
	}

	for _, tt := range tests {
		values := ParseDigitsAndOutputValues(tt.Filename)
		assert.Equal(t, tt.Want, values.CountOneFourSevenEights())
	}
}
