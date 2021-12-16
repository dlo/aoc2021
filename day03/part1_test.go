package day03

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateGammaRate(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 198},
		{"testdata/input.txt", 3912944},
	}

	for _, tt := range tests {
		reportValues, length := utils.BinaryLinesFromFile(tt.Filename)
		report := DiagnosticReport{reportValues, length}
		assert.EqualValues(t, tt.Want, int(report.CalculatePowerConsumption()))
	}
}
