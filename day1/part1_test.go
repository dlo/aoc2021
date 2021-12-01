package day1

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1CalculateNumberOfIncreases(t *testing.T) {
	var tests = []struct {
		filename string
		want     int
	}{
		{"day1_input.txt", 1374},
	}

	for _, tt := range tests {
		items := utils.LinesFromFile(tt.filename)
		if got := Part1CalculateNumberOfIncreases(items); got != tt.want {
			assert.Equal(t, tt.want, got)
		}
	}
}
