package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	values []int
	want   int
}

func TestSumSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	var tests = []TestCase{
		{slice[:3], 6},
		{slice[3:], 15},
		{slice[1:2], 2},
	}

	for _, tt := range tests {
		got := SliceSum(tt.values)
		assert.Equal(t, got, tt.want)
	}
}
