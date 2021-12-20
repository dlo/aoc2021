package utils

import "errors"

type LIFOQueueRune []rune

func (s *LIFOQueueRune) Push(r rune) {
	*s = append(*s, r)
}

func (s *LIFOQueueRune) UnsafePop() {
	_, _ = s.Pop()
}

func (s *LIFOQueueRune) Pop() (rune, error) {
	if len(*s) == 0 {
		return '-', errors.New("queue is empty")
	}

	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result, nil
}
