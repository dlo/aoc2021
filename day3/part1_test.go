package day3

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CalculateGammaRateTestCase struct {
	filename string
	want int
}

func TestCalculateGammaRate(t *testing.T) {
	var tests = []CalculateGammaRateTestCase{
		{"testdata/example_input.txt", 198},
		{"testdata/input.txt", 3912944},
	}

	for _, tt := range tests {
		reportValues, length := utils.BinaryLinesFromFile(tt.filename)
		report := DiagnosticReport{reportValues, length}
		assert.EqualValues(t, tt.want, int(report.CalculatePowerConsumption()))
	}
}
