package day5

import (
	"github.com/dlo/aoc2021/day2"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CoordinateFromStringTestCase struct {
	input string
	want  day2.Coordinate
}

type ParseLineSegmentFromLineTestCase struct {
	input string
	want  LineSegment
}

type CountIntersectionsTestCase struct {
	filename string
	want     int
}

func TestCoordinateFromString(t *testing.T) {
	var tests = []CoordinateFromStringTestCase{
		{"0,9", day2.Coordinate{X: 0, Y: 9}},
		{"5,9", day2.Coordinate{X: 5, Y: 9}},
	}

	for _, tt := range tests {
		coordinate := CoordinateFromString(tt.input)
		assert.Equal(t, tt.want.X, coordinate.X)
		assert.Equal(t, tt.want.Y, coordinate.Y)
	}
}

func TestParseLineSegmentFromLine(t *testing.T) {
	var tests = []ParseLineSegmentFromLineTestCase{
		{"5,5 -> 8,2", LineSegment{day2.Coordinate{X: 5, Y: 5}, day2.Coordinate{X: 8, Y: 2}}},
	}

	for _, tt := range tests {
		lineSegment := ParseLineSegmentFromLine(tt.input)
		assert.Equal(t, tt.want.p1.X, lineSegment.p1.X)
		assert.Equal(t, tt.want.p1.Y, lineSegment.p1.Y)
		assert.Equal(t, tt.want.p2.X, lineSegment.p2.X)
		assert.Equal(t, tt.want.p2.Y, lineSegment.p2.Y)
	}
}

func TestCountIntersections(t *testing.T) {
	var tests = []CountIntersectionsTestCase{
		{"testdata/example_input.txt", 5},
		{"testdata/input.txt", 5585},
	}

	for _, tt := range tests {
		measurements := ParseHydrothermalVentMeasurements(tt.filename)
		measurements.Process(false)
		assert.Equal(t, tt.want, measurements.CountIntersections())
	}
}

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
