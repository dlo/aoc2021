package day15

import (
	"github.com/dlo/aoc2021/utils"
)

func (c Cavern) MultiplyCavern() Cavern {
	lenX, lenY := c.grid.Size()
	// %10 - 1
	grid := make(utils.IntMatrix, 5*lenY)
	for j := range grid {
		grid[j] = make([]int, 5*lenX)
	}

	for j := range c.grid {
		for i := range c.grid[j] {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					newValue := c.grid[j][i] + k + l
					if newValue > 9 {
						newValue = newValue%10 + 1
					}
					grid[j+lenY*l][i+lenX*k] = newValue
				}
			}
		}
	}

	return Cavern{grid, c.riskSums}
}
