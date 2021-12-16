package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompletionScore(t *testing.T) {
	tests := []TestCase{
		{"testdata/example_input.txt", 288957},
		{"testdata/input.txt", 3490802734},
	}

	for _, tt := range tests {
		subsystem := GenerateRawSubsystemFromFile(tt.filename)
		assert.Equal(t, tt.want, subsystem.CompletionScore())
	}
}
