package day15

import (
	"fmt"
	"github.com/dlo/aoc2021/utils"
	"github.com/gdamore/tcell/v2"
	"math"
)

type Cavern struct {
	grid     utils.IntMatrix
	riskSums utils.IntMatrix
}

func (c Cavern) Println() {
	c.grid.Println()
}

func (c Cavern) Display(s tcell.Screen) {
	c.grid.Display(s, 0, 0)
	c.riskSums.Display(s, 0, 11)
}

func (c Cavern) PrintRisk() {
	for j := range c.riskSums {
		if j > 10 {
			break
		}

		for i := range c.riskSums[j] {
			if i > 10 {
				break
			}

			point := utils.Point{i, j}
			fmt.Printf("%4d", c.riskSums.UnsafeValueAt(point))
		}
		fmt.Println()
	}
}

func (c Cavern) Risk() int {
	return c.riskSums[0][0]
}

func (c Cavern) DijkstraPath(p utils.Point) utils.IntMatrix {
	return c.grid.DijkstraPath(p)
}

func (c Cavern) CalculateRiskSums(p utils.Point) int {
	if c.grid.IsBottomRightPoint(p) {
		return 0
	}

	potentials := []utils.Point{
		p.Move(1, 0),
		p.Move(0, 1),
	}

	min := math.MaxInt
	for _, potential := range potentials {
		potentialValue := -1
		if potentialValue, _ = c.riskSums.ValueAt(potential); potentialValue == 0 {
			potentialValue = c.grid.UnsafeValueAt(potential) + c.CalculateRiskSums(potential)
			c.riskSums.SetValue(potential, potentialValue)
		}

		if potentialValue != -1 && potentialValue < min {
			min = potentialValue
		}
	}

	c.riskSums.SetValue(p, min)
	return min
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

	riskSums := make(utils.IntMatrix, len(riskLevels))
	for i := range riskLevels {
		riskSums[i] = make([]int, len(riskLevels[0]))
	}
	return Cavern{riskLevels, riskSums}
}
