package day17

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxY(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 45},
		{"testdata/input.txt", 9180},
	}

	for _, tt := range tests {
		area := ImportData(tt.Filename)
		solution := area.FindSolution()
		assert.Equal(t, tt.Want, solution.MaxY())
	}
}
