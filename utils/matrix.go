package utils

import (
	"fmt"
)

type Point [2]int
type IntMatrix [][]int

func NewIntMatrix(height int, width int) IntMatrix {
	var m IntMatrix = make([][]int, height)
	for j := range m {
		m[j] = make([]int, width)
	}
	return m
}

func (m IntMatrix) Neighbors(p Point) []Point {
	return m.NeighborsForXY(p.Coordinates())
}

func (m IntMatrix) ValueAt(point Point) int {
	x, y := point.Coordinates()
	return m[y][x]
}

func (p Point) Coordinates() (x int, y int) {
	return p[0], p[1]
}

func (m IntMatrix) Println() {
	for y := range m {
		for x := range m[y] {
			point := Point{x, y}
			fmt.Print(m.ValueAt(point))
		}
		fmt.Println()
	}
}

func (m IntMatrix) IsValidXY(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(m) && x < len(m[y])
}

func (m IntMatrix) IsValidPoint(p Point) bool {
	return m.IsValidXY(p.Coordinates())
}

func (m IntMatrix) Area() int {
	return len(m) * len(m[0])
}

func (m IntMatrix) NeighborsForXY(x, y int) []Point {
	var neighbors []Point
	for _, point := range []Point{
		{x - 1, y - 1},
		{x - 1, y},
		{x - 1, y + 1},
		{x, y - 1},
		{x, y + 1},
		{x + 1, y - 1},
		{x + 1, y},
		{x + 1, y + 1},
	} {
		if !m.IsValidPoint(point) {
			continue
		}

		neighbors = append(neighbors, point)
	}
	return neighbors
}
