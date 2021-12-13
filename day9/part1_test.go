package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	want     int
}

func TestLowPoints(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 15},
		{"testdata/input.txt", 486},
	}

	for _, tt := range tests {
		values := GenerateHeightMap(tt.filename)
		points := values.LowPoints()
		assert.Equal(t, tt.want, points.Risk)
	}
}
