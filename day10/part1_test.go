package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	want     int
}

func TestGenerateSubsystemFromFile(t *testing.T) {
	tests := []TestCase{
		{"testdata/example_input.txt", 26397},
		{"testdata/input.txt", 294195},
	}

	for _, tt := range tests {
		subsystem := GenerateRawSubsystemFromFile(tt.filename)
		assert.Equal(t, tt.want, subsystem.SyntaxErrorScore())
	}
}
