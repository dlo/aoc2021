package day14

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifferenceOfMostCommonElementFromLeastCommonElement40Steps(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 2188189693529},
		{"testdata/input.txt", 2587447599164},
	}

	for _, tt := range tests {
		rules := ImportRules(tt.Filename)
		fastInstructions := rules.NewFastInstructions()
		fastInstructions.StepCount(40)
		assert.Equal(t, tt.Want, fastInstructions.DifferenceOfMostCommonElementFromLeastCommonElement())
	}
}
