package day10

import (
	"github.com/dlo/aoc2021/utils"
	"sort"
)

func CompletionScore() map[rune]int {
	return map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
}

func (c RawChunks) CompletionScore() int {
	queue := utils.LIFOQueueRune{}
	delimiters := MatchingDelimiters()
	scores := CompletionScore()
	for _, r := range []rune(c) {
		switch r {
		case '(', '[', '{', '<':
			queue.Push(r)

		case ')', ']', '}', '>':
			queue.UnsafePop()

		default:
			panic("yikes")
		}
	}

	score := 0
	for {
		delimiter, err := queue.Pop()
		if err != nil {
			break
		}
		score = score*5 + scores[delimiters[delimiter]]
	}
	return score
}

func (s RawSubsystem) CompletionScore() int {
	var scores []int
	for _, line := range s {
		if line.SyntaxErrorScore() > 0 {
			continue
		}
		scores = append(scores, line.CompletionScore())
	}

	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}
