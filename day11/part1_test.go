package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	want     int
}

func TestFlashCount(t *testing.T) {
	tests := []TestCase{
		{"testdata/example_input.txt", 1656},
		{"testdata/input.txt", 1721},
	}

	for _, tt := range tests {
		grid := ReadOctopusGrid(tt.filename)
		grid.StepByCount(100)
		assert.Equal(t, tt.want, grid.flashCount)
	}
}
