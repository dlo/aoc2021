package day10

import "errors"

type LIFOQueue []rune

func (s *LIFOQueue) Push(r rune) {
	*s = append(*s, r)
}

func (s *LIFOQueue) Pop() (rune, error) {
	if len(*s) == 0 {
		return '-', errors.New("queue is empty")
	}

	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result, nil
}
