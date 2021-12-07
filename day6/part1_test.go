package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	days     int
	want     int
}

func TestCoordinateFromString(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 18, 26},
		{"testdata/example_input.txt", 80, 5934},
		{"testdata/input.txt", 80, 372300},
		{"testdata/example_input.txt", 256, 26984457539},
		{"testdata/input.txt", 256, 1675781200288},
	}

	for _, tt := range tests {
		ages := ParseFishFile(tt.filename)
		assert.Equal(t, tt.want, CalculateSchoolSizeAfterNDays(ages, tt.days))
	}
}
