package day11

import (
	"fmt"
	"github.com/dlo/aoc2021/utils"
	"github.com/fatih/color"
)

type EnergyLevel int
type Point [2]int
type OctopusGrid struct {
	flashed    map[Point]bool
	grid       [][]EnergyLevel
	flashCount int
}

func (p Point) Coordinates() (x int, y int) {
	return p[0], p[1]
}

func (g OctopusGrid) Println() {
	for y := range g.grid {
		for x := range g.grid[y] {
			point := Point{x, y}
			if g.flashed[point] {
				color.New(color.FgRed).PrintFunc()(g.EnergyLevel(point))
			} else {
				fmt.Print(g.EnergyLevel(point))
			}
		}
		fmt.Println()
	}
}

func (g OctopusGrid) IsValidXY(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(g.grid) && x < len(g.grid[y])
}

func (g OctopusGrid) IsValidPoint(p Point) bool {
	x, y := p.Coordinates()
	return g.IsValidXY(x, y)
}

func (g OctopusGrid) NeighborsForXY(x, y int) []Point {
	points := []Point{
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
		{x, y - 1},
		{x, y + 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
	}

	var validPoints []Point
	for _, point := range points {
		if g.IsValidPoint(point) {
			validPoints = append(validPoints, point)
		}
	}
	return validPoints
}

func (g OctopusGrid) Neighbors(p Point) []Point {
	return g.NeighborsForXY(p.Coordinates())
}

func (g OctopusGrid) EnergyLevel(point Point) EnergyLevel {
	x, y := point.Coordinates()
	return g.grid[y][x]
}

func (g *OctopusGrid) IncreaseEnergy(p Point) {
	if g.flashed[p] {
		return
	}

	x, y := p.Coordinates()
	g.grid[y][x]++
	if g.EnergyLevel(p) > 9 {
		g.flashed[p] = true
		g.Flash(p)
	}
}

func (g *OctopusGrid) Flash(p Point) {
	for _, neighbor := range g.Neighbors(p) {
		g.IncreaseEnergy(neighbor)
	}
}

func (g *OctopusGrid) Step() {
	g.flashed = map[Point]bool{}
	for j := range g.grid {
		for i := range g.grid[j] {
			g.IncreaseEnergy(Point{i, j})
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
	grid := make([][]EnergyLevel, len(lines))
	for i, line := range lines {
		var row []EnergyLevel
		for _, r := range []rune(line) {
			row = append(row, EnergyLevel(r-'0'))
		}
		grid[i] = row
	}
	return OctopusGrid{map[Point]bool{}, grid, 0}
}
