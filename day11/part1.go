package day11

import (
	"fmt"
	"github.com/dlo/aoc2021/utils"
	"github.com/fatih/color"
)

type EnergyLevel int
type OctopusGrid struct {
	flashed    map[utils.Point]bool
	grid       utils.IntMatrix
	flashCount int
}

func (g OctopusGrid) Println() {
	for y := range g.grid {
		for x := range g.grid[y] {
			point := utils.Point{x, y}
			if g.flashed[point] {
				color.New(color.FgRed).PrintFunc()(g.EnergyLevel(point))
			} else {
				fmt.Print(g.EnergyLevel(point))
			}
		}
		fmt.Println()
	}
}

func (g OctopusGrid) Neighbors(p utils.Point) []utils.Point {
	return g.grid.NeighborsForXY(p.Coordinates())
}

func (g OctopusGrid) EnergyLevel(point utils.Point) EnergyLevel {
	return EnergyLevel(g.grid.UnsafeValueAt(point))
}

func (g *OctopusGrid) IncreaseEnergy(p utils.Point) {
	if g.flashed[p] {
		return
	}

	x, y := p.Coordinates()
	g.grid[y][x]++
	if g.EnergyLevel(p) > 9 {
		g.Flash(p)
	}
}

func (g *OctopusGrid) Flash(p utils.Point) {
	g.flashed[p] = true
	for _, neighbor := range g.Neighbors(p) {
		g.IncreaseEnergy(neighbor)
	}
}

func (g *OctopusGrid) Step() {
	g.flashed = map[utils.Point]bool{}
	for j := range g.grid {
		for i := range g.grid[j] {
			g.IncreaseEnergy(utils.Point{i, j})
		}
	}

	for k := range g.flashed {
		x, y := k.Coordinates()
		g.grid[y][x] = 0
	}

	g.flashCount += len(g.flashed)
}

func (g *OctopusGrid) StepByCount(count int) {
	for i := 0; i < count; i++ {
		g.Step()
	}
}

func ReadOctopusGrid(filename string) OctopusGrid {
	lines := utils.ReadLinesFromFile(filename)
	grid := make([][]int, len(lines))
	for i, line := range lines {
		var row []int
		for _, r := range []rune(line) {
			row = append(row, int(r-'0'))
		}
		grid[i] = row
	}
	return OctopusGrid{map[utils.Point]bool{}, grid, 0}
}
