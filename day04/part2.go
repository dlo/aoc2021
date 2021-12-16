package day04

func (b *Bingo) ScoreOfLastWinningBoard() int {
	scoreOfLastWinningBoard := 0
	completedBoards := make([]bool, len(b.Boards))
	for i := range b.Numbers {
		for j, board := range b.Boards {
			if completedBoards[j] {
				continue
			}

			if isWinning, score := board.IsWinning(b.Numbers[:i]); isWinning {
				completedBoards[j] = true
				scoreOfLastWinningBoard = score
			}
		}
	}

	return scoreOfLastWinningBoard
}
