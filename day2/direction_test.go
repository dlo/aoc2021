package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DirectionTestCase struct {
	input string
	want  Direction
}

func TestDirectionFromRawDirection(t *testing.T) {
	var tests = []DirectionTestCase{
		{"up", Up},
		{"down", Down},
		{"forward", Forward},
		{"random", Invalid},
	}


	for _, tt := range tests {
		assert.Equal(t, tt.want, DirectionFromRawDirection(tt.input))
	}
}