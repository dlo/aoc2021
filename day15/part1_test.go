package day15

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifferenceOfMostCommonElementFromLeastCommonElement40Steps(t *testing.T) {
	tests := []utils.SimpleTestCase{
		{"testdata/example_input.txt", 40},
		{"testdata/input.txt", 540},
	}

	for _, tt := range tests {
		cavern := ImportCavernData(tt.Filename)
		m := cavern.DijkstraPath(utils.Point{0, 0})
		maxX, maxY := m.Size()
		point := utils.Point{maxX - 1, maxY - 1}
		assert.Equal(t, tt.Want, m.UnsafeValueAt(point))
	}
}
