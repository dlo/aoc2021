package day13

import (
	"errors"
	"fmt"
	"github.com/dlo/aoc2021/utils"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type Grid [][]bool
type Instruction string

type Puzzle struct {
	grid  Grid
	dots  []utils.Point
	folds []utils.Point
}

func (g Grid) Println() {
	for y := range g {
		for x := range g[y] {
			if g[y][x] {
				color.New(color.FgHiRed, color.Bold).PrintFunc()("#")
				//fmt.Print("#")
			} else {
				color.New(color.FgBlack).PrintFunc()(".")
			}
		}
		fmt.Println()
	}
}

func (i Instruction) ToPoint() (utils.Point, error) {
	value := string(i)
	idx := strings.IndexRune(string(i), ',')
	invalid := utils.Point{-1, -1}
	if idx == -1 {
		return invalid, errors.New("invalid")
	}

	x, err := strconv.Atoi(value[:idx])
	if err != nil {
		return invalid, err
	}
	y, err := strconv.Atoi(value[idx+1:])
	if err != nil {
		return invalid, err
	}

	return utils.Point{x, y}, nil
}

func (p Puzzle) Println() {
	p.grid.Println()
}

func (p Puzzle) CountDots() int {
	count := 0
	for y := range p.grid {
		for x := range p.grid[y] {
			if p.grid[y][x] {
				count++
			}
		}
	}
	return count
}

func (p Puzzle) FoldAll() Puzzle {
	result := p
	for len(result.folds) > 0 {
		result = result.Fold()
	}
	return result
}

func (p Puzzle) Fold() Puzzle {
	fold := p.folds[0]
	folds := p.folds[1:]
	foldX, foldY := fold.Coordinates()
	var newDots []utils.Point
	switch {
	case foldX == 0:
		for _, dot := range p.dots {
			x, y := dot.Coordinates()
			if y > foldY {
				newDots = append(newDots, utils.Point{x, 2*foldY - y})
			} else {
				newDots = append(newDots, dot)
			}
		}

	case foldY == 0:
		for _, dot := range p.dots {
			x, y := dot.Coordinates()
			if x > foldX {
				newDots = append(newDots, utils.Point{2*foldX - x, y})
			} else {
				newDots = append(newDots, dot)
			}
		}
	}

	grid := GenerateGrid(newDots)
	return Puzzle{grid, newDots, folds}
}

func GenerateGrid(dots []utils.Point) Grid {
	maxX := 0
	maxY := 0
	for _, dot := range dots {
		x, y := dot.Coordinates()
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	grid := make(Grid, maxY+1)
	for j := range grid {
		grid[j] = make([]bool, maxX+1)
	}

	for _, dot := range dots {
		x, y := dot.Coordinates()
		grid[y][x] = true
	}

	return grid
}

func ImportGrid(filename string) Puzzle {
	lines := utils.ReadLinesFromFile(filename)
	var folds []utils.Point
	var dots []utils.Point
	for _, line := range lines {
		point, err := Instruction(line).ToPoint()
		if err != nil {
			if strings.HasPrefix(line, "fold along") {
				jdx := strings.IndexRune(line, '=')
				isX := line[jdx-1:jdx] == "x"
				value, _ := strconv.Atoi(line[jdx+1:])
				var line utils.Point
				if isX {
					line = utils.Point{value, 0}
				} else {
					line = utils.Point{0, value}
				}
				folds = append(folds, line)
			}

			continue
		}

		dots = append(dots, point)
	}

	grid := GenerateGrid(dots)
	return Puzzle{grid, dots, folds}
}
