package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreOfLastWinningBoard(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 1924},
		{"testdata/input.txt", 1284},
	}

	for _, tt := range tests {
		result := ParseBingoCardDataFromFile(tt.filename)
		assert.Equal(t, result.ScoreOfLastWinningBoard(), tt.want)
	}
}
