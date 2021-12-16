package day04

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindWinningScore(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 4512},
		{"testdata/input.txt", 11536},
	}

	for _, tt := range tests {
		result := ParseBingoCardDataFromFile(tt.Filename)
		assert.Equal(t, result.FindWinningScore(), tt.Want)
	}
}
