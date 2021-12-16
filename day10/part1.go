package day10

import (
	"fmt"
	"github.com/dlo/aoc2021/utils"
)

type Delimiter rune
type RawChunks string
type RawSubsystem []RawChunks

//func (l Line) SyntaxErrorScore() int {
//
//}

func (s RawSubsystem) Println() {
	for _, line := range s {
		fmt.Println(line)
	}
}

func MatchingDelimiters() map[rune]rune {
	return map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
}

func DelimiterScores() map[rune]int {
	return map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
}

func (c RawChunks) SyntaxErrorScore() int {
	queue := LIFOQueue{}
	delimiters := MatchingDelimiters()
	scores := DelimiterScores()
	for _, r := range []rune(c) {
		switch r {
		case '(', '[', '{', '<':
			queue.Push(r)
			continue

		case ')', ']', '}', '>':
			opener, err := queue.Pop()
			if err != nil {
				panic("oh noes")
			}

			if delimiters[opener] != r {
				return scores[r]
			}

		default:
			panic("yikes")
		}
	}

	return 0
}

func (s RawSubsystem) SyntaxErrorScore() int {
	errorScore := 0
	for _, line := range s {
		errorScore += line.SyntaxErrorScore()
	}
	return errorScore
}

func GenerateRawSubsystemFromFile(filename string) RawSubsystem {
	lines := utils.ReadLinesFromFile(filename)
	var subsystem RawSubsystem

	for _, line := range lines {
		subsystem = append(subsystem, RawChunks(line))
	}
	return subsystem
}
