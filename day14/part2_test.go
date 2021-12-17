package day14

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifferenceOfMostCommonElementFromLeastCommonElement40Steps(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 1588},
		{"testdata/input.txt", 2435},
	}

	for _, tt := range tests {
		rules := ImportRules(tt.Filename)
		fastInstructions := rules.NewFastInstructions()
		fastInstructions.StepCount(40)
		assert.Equal(t, tt.Want, fastInstructions.DifferenceOfMostCommonElementFromLeastCommonElement())
	}
}
