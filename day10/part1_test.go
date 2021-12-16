package day10

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyntaxErrorScore(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 26397},
		{"testdata/input.txt", 294195},
	}

	for _, tt := range tests {
		subsystem := GenerateRawSubsystemFromFile(tt.Filename)
		assert.Equal(t, tt.Want, subsystem.SyntaxErrorScore())
	}
}
