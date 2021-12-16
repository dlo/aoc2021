package day08

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumDecodedInputs(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 61229},
		{"testdata/input.txt", 915941},
	}

	for _, tt := range tests {
		values := ParseDigitsAndOutputValues(tt.Filename)
		assert.Equal(t, tt.Want, values.SumDecodedInputs())
	}
}
