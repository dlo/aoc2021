package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValue(t *testing.T) {
	tests := []TestCase{
		{"testdata/example_input_2.txt", []int{3, 54, 7, 9, 1, 0, 0, 1}},
		{"testdata/input.txt", []int{180616437720}},
	}

	for _, tt := range tests {
		transmissions := ImportBITSTransmission(tt.filename)
		for i, want := range tt.want {
			transmission := transmissions[i]
			result, _ := transmission.Process()
			assert.Equal(t, want, result.Value())
		}
	}
}
