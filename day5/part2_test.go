package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountIntersectionsWithDiagonals(t *testing.T) {
	var tests = []CountIntersectionsTestCase{
		{"testdata/example_input.txt", 12},
		{"testdata/input.txt", 17193},
	}

	for _, tt := range tests {
		measurements := ParseHydrothermalVentMeasurements(tt.filename)
		measurements.Process(true)
		assert.Equal(t, tt.want, measurements.CountIntersections())
	}
}
