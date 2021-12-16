package day03

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateOxygenGeneratorRating(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 23},
		{"testdata/input.txt", 3597},
	}

	for _, tt := range tests {
		reportValues, length := utils.BinaryLinesFromFile(tt.Filename)
		report := DiagnosticReport{reportValues, length}
		assert.EqualValues(t, tt.Want, int(report.CalculateOxygenGeneratorRating(1<<report.Length)))
	}
}

func TestCalculateCO2ScrubberRating(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 10},
		{"testdata/input.txt", 1389},
	}

	for _, tt := range tests {
		reportValues, length := utils.BinaryLinesFromFile(tt.Filename)
		report := DiagnosticReport{reportValues, length}
		assert.EqualValues(t, tt.Want, int(report.CalculateCO2ScrubberRating(1<<report.Length)))
	}
}

func TestCalculateLifeSupportRating(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 230},
		{"testdata/input.txt", 4996233},
	}

	for _, tt := range tests {
		reportValues, length := utils.BinaryLinesFromFile(tt.Filename)
		report := DiagnosticReport{reportValues, length}
		assert.EqualValues(t, tt.Want, int(report.CalculateLifeSupportRating()))
	}
}
