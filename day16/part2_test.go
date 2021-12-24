package day16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValue(t *testing.T) {
	tests := []TestCase{
		{3}, {54}, {7}, {9}, {1}, {0}, {0}, {1},
	}

	transmissions := ImportBITSTransmission("testdata/example_input_2.txt")
	for i, tt := range tests {
		transmission := transmissions[i]
		result, _ := transmission.Process()
		assert.Equal(t, tt.want, result.Value())
	}
}
