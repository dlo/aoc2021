package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	want     []int
}

func TestVersionNumberSum(t *testing.T) {
	tests := []TestCase{
		{"testdata/example_input.txt", []int{16, 12, 23, 31, 974}},
		{"testdata/input.txt", []int{974}},
	}

	for _, tt := range tests {
		transmissions := ImportBITSTransmission(tt.filename)
		for i, want := range tt.want {
			transmission := transmissions[i]
			result, _ := transmission.Process()
			assert.Equal(t, want, result.VersionNumberSum())
		}
	}
}
