package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	want     int
}

func TestCalculateCheapestPositionPart2(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 61229},
		{"testdata/input.txt", 915941},
	}

	for _, tt := range tests {
		values := ParseDigitsAndOutputValues(tt.filename)
		assert.Equal(t, tt.want, values.SumDecodedInputs())
	}
}
