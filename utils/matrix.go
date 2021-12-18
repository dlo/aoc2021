package utils

import (
	"errors"
	"fmt"
	"strconv"
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

func (p Point) Move(deltaX, deltaY int) Point {
	x, y := p.Coordinates()
	return Point{x + deltaX, y + deltaY}
}

func (p Point) Coordinates() (x int, y int) {
	return p[0], p[1]
}

func (m IntMatrix) IsBottomRightPoint(p Point) bool {
	x, y := p.Coordinates()
	return y == len(m)-1 && x == len(m[0])-1
}

func (m IntMatrix) Neighbors(p Point) []Point {
	return m.NeighborsForXY(p.Coordinates())
}

func (m IntMatrix) SetValue(point Point, value int) {
	x, y := point.Coordinates()
	m[y][x] = value
}

func (m IntMatrix) UnsafeValueAt(point Point) int {
	x, y := point.Coordinates()
	return m[y][x]
}

func (m IntMatrix) ValueAt(point Point) (int, error) {
	if !m.IsValidPoint(point) {
		return -1, errors.New("invalid")
	}

	x, y := point.Coordinates()
	return m[y][x], nil
}

func (m IntMatrix) PrintlnWidth(width int64) {
	for y := range m {
		for x := range m[y] {
			point := Point{x, y}
			format := "%" + strconv.FormatInt(width, 10) + "d"
			fmt.Printf(format, m.UnsafeValueAt(point))
		}
		fmt.Println()
	}
}

func (m IntMatrix) Println() {
	for y := range m {
		for x := range m[y] {
			point := Point{x, y}
			fmt.Print(m.UnsafeValueAt(point))
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
