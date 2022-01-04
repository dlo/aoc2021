package day18

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MagnitudeTestCase struct {
	input string
	want  int
}

func TestSnailfishNumber_Magnitude(t *testing.T) {
	tests := []MagnitudeTestCase{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, RawSnailfishNumber(test.input).Parse().Magnitude())
	}
}
