package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLIFOQueue_Pop(t *testing.T) {
	queue := LIFOQueueRune{}
	queue.Push('1')
	queue.Push('2')
	result, _ := queue.Pop()
	assert.Equal(t, '2', result)
}
