package day10

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompletionScore(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 288957},
		{"testdata/input.txt", 3490802734},
	}

	for _, tt := range tests {
		subsystem := GenerateRawSubsystemFromFile(tt.Filename)
		assert.Equal(t, tt.Want, subsystem.CompletionScore())
	}
}
