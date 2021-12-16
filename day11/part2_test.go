package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstSynchronousStep(t *testing.T) {
	tests := []TestCase{
		{"testdata/example_input.txt", 195},
		{"testdata/input.txt", 298},
	}

	for _, tt := range tests {
		grid := ReadOctopusGrid(tt.filename)
		assert.Equal(t, tt.want, grid.FirstSynchronousStep())
	}
}
