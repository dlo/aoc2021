package day15

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstraPathMultiplied(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 315},
		{"testdata/input.txt", 2879},
	}

	for _, tt := range tests {
		cavern := ImportCavernData(tt.Filename)
		nc := cavern.MultiplyCavern()
		path := nc.DijkstraPath(utils.Point{0, 0})
		maxX, maxY := path.Size()
		point := utils.Point{maxX - 1, maxY - 1}
		assert.Equal(t, tt.Want, path.UnsafeValueAt(point))
	}
}
