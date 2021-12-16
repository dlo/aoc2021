package day14

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifferenceOfMostCommonElementFromLeastCommonElement(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 1588},
		{"testdata/input.txt", 2435},
	}

	for _, tt := range tests {
		rules := ImportRules(tt.Filename)
		rules.StepCount(10)
		assert.Equal(t, tt.Want, rules.DifferenceOfMostCommonElementFromLeastCommonElement())
	}
}

func TestDifferenceOfMostCommonElementFromLeastCommonElement40Steps(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 1588},
		{"testdata/input.txt", 2435},
	}

	for _, tt := range tests {
		rules := ImportRules(tt.Filename)
		rules.StepCount(40)
		assert.Equal(t, tt.Want, rules.DifferenceOfMostCommonElementFromLeastCommonElement())
	}
}
