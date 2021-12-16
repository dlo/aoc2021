package day04

import (
	"github.com/dlo/aoc2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreOfLastWinningBoard(t *testing.T) {
	var tests = []utils.SimpleTestCase{
		{"testdata/example_input.txt", 1924},
		{"testdata/input.txt", 1284},
	}

	for _, tt := range tests {
		result := ParseBingoCardDataFromFile(tt.Filename)
		assert.Equal(t, result.ScoreOfLastWinningBoard(), tt.Want)
	}
}
