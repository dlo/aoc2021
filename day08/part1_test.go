package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	want     int
}

func TestCountOneFourSevenEights(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 26},
		{"testdata/input.txt", 310},
	}

	for _, tt := range tests {
		values := ParseDigitsAndOutputValues(tt.filename)
		assert.Equal(t, tt.want, values.CountOneFourSevenEights())
	}
}
