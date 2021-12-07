package day4

import (
	"github.com/dlo/aoc2021/utils"
)

type BingoBoard struct {
	Grid [][]int
}

type Bingo struct {
	Numbers []int
	Boards  []BingoBoard
}

func (b *BingoBoard) IsWinning(numbers []int) (bool, int) {
	numbersSet := map[int]bool{}
	for _, number := range numbers {
		numbersSet[number] = true
	}

	colCounts := [5]int{0, 0, 0, 0, 0}
	rowCounts := [5]int{0, 0, 0, 0, 0}
	sum := 0
	isWinning := false
	for i, row := range b.Grid {
		for j, value := range row {
			if numbersSet[value] {
				rowCounts[i]++
				colCounts[j]++

				if rowCounts[i] == 5 ||
					colCounts[j] == 5 {
					isWinning = true
				}
			} else {
				sum += value
			}
		}
	}

	if isWinning {
		lastNumber := numbers[len(numbers)-1]
		return true, lastNumber * sum
	}

	return false, 0
}

func (b *Bingo) FindWinningScore() int {
	for i := range b.Numbers {
		for _, board := range b.Boards {
			if isWinning, score := board.IsWinning(b.Numbers[:i]); isWinning {
				return score
			}
		}
	}

	return -1
}

func ParseBingoCardDataFromFile(filename string) Bingo {
	lines := utils.ReadLinesFromFile(filename)
	numbers := utils.ParseNumbersFromString(lines[0])

	var boards []BingoBoard
	var grid [][]int
	for _, line := range lines[1:] {
		row := utils.ParseNumbersFromString(line)
		if len(row) > 0 {
			grid = append(grid, row)
		}

		if len(grid) == 5 {
			boards = append(boards, BingoBoard{grid})
			grid = nil
		}
	}

	return Bingo{Numbers: numbers, Boards: boards}
}
