package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	want int
}

func TestVersionNumberSum(t *testing.T) {
	tests := []TestCase{
		{16}, {12}, {23}, {31}, {974},
	}

	transmissions := ImportBITSTransmission("testdata/example_input.txt")
	for i, tt := range tests {
		transmission := transmissions[i]
		result, _ := transmission.Process()
		assert.Equal(t, tt.want, result.VersionNumberSum())
	}
}
