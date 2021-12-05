package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	want     int
}

func TestFindWinningScore(t *testing.T) {
	var tests = []TestCase{
		{"testdata/example_input.txt", 4512},
		{"testdata/input.txt", 11536},
	}

	for _, tt := range tests {
		result := ParseBingoCardDataFromFile(tt.filename)
		assert.Equal(t, result.FindWinningScore(), tt.want)
	}
}
