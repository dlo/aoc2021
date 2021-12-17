package day15

import "github.com/dlo/aoc2021/utils"

type Cavern struct {
	riskLevels utils.IntMatrix
}

func (c Cavern) Println() {
	c.riskLevels.Println()
}

func ImportCavernData(filename string) Cavern {
	lines := utils.ReadLinesFromFile(filename)
	riskLevels := make(utils.IntMatrix, 0)
	for _, line := range lines {
		row := make([]int, 0)
		for _, r := range []rune(line) {
			row = append(row, int(r-'0'))
		}
		riskLevels = append(riskLevels, row)
	}
	return Cavern{riskLevels}
}
