package utils

import (
	"errors"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Point [2]int
type IntRow []int
type IntMatrix [][]int

func NewIntMatrix(height int, width int) IntMatrix {
	m := make(IntMatrix, height)
	for j := range m {
		m[j] = make(IntRow, width)
	}
	return m
}

func DisplayWord(s tcell.Screen, str string, x int, y int, style tcell.Style) {
	for i, r := range []rune(str) {
		s.SetContent(x+i, y, r, nil, style)
	}
}

func (m IntMatrix) Size() (int, int) {
	return len(m), len(m[0])
}

func (m IntMatrix) Display(s tcell.Screen, xOffset int, yOffset int) {
	styles := []tcell.Style{
		tcell.StyleDefault.Foreground(tcell.ColorOlive),
		tcell.StyleDefault.Foreground(tcell.ColorPurple),
		tcell.StyleDefault.Foreground(tcell.ColorYellow),
		tcell.StyleDefault.Foreground(tcell.ColorMaroon),
		tcell.StyleDefault.Foreground(tcell.ColorAquaMarine),
	}
	rand.Seed(time.Now().UnixNano())
	style := styles[rand.Intn(len(styles))]
	maxWidth := 0
	for j := range m {
		for i := range m[j] {
			point := Point{i, j}
			digits := int(math.Log10(float64(m.UnsafeValueAt(point))))
			if digits > maxWidth {
				maxWidth = digits
			}
		}
	}

	maxWidth += 2
	format := fmt.Sprintf("%%%dd", maxWidth)
	for j := range m {
		for i := range m[j] {
			point := Point{i, j}
			value := m.UnsafeValueAt(point)
			DisplayWord(s, fmt.Sprintf(format, value), i*maxWidth+xOffset, j+yOffset, style)
		}
	}
	s.Sync()
	s.Show()
}

func (m IntMatrix) DijkstraPath(p Point) IntMatrix {
	return m.DijkstraPathXY(p.Coordinates())
}

func (m IntMatrix) DijkstraPathXY(x, y int) IntMatrix {
	queue := map[Point]bool{}
	dist := map[Point]int{}

	for j := range m {
		for i := range m[j] {
			p := Point{i, j}

			// Distance * Max 9 per block
			dist[p] = (i + j) * 9
			queue[p] = true
		}
	}

	cur := Point{x, y}
	dist[cur] = 0
	var minCost int

	for queue[cur] {
		delete(queue, cur)

		neighbors := m.NonDiagonalNeighbors(cur)
		for _, neighbor := range neighbors {
			distance := m.UnsafeValueAt(neighbor)

			distanceToNeighbor := distance + dist[cur]
			if distanceToNeighbor < dist[neighbor] {
				dist[neighbor] = distanceToNeighbor
			}
		}

		minCost = math.MaxInt
		for k := range queue {
			v := dist[k]
			if v < minCost {
				cur = k
				minCost = v
			}
		}
	}

	distMatrix := make(IntMatrix, len(m))
	for j := range m {
		distMatrix[j] = make([]int, len(m[j]))
	}

	for k, v := range dist {
		distMatrix.SetValue(k, v)
	}

	return distMatrix
}

func (p Point) Move(deltaX, deltaY int) Point {
	x, y := p.Coordinates()
	return Point{x + deltaX, y + deltaY}
}

func (p Point) Coordinates() (x int, y int) {
	return p[0], p[1]
}

func (m IntMatrix) GenerateDistanceMatrix() IntMatrix {
	distances := make(IntMatrix, len(m))
	for j := range distances {
		distances[j] = make(IntRow, len(m[j]))
	}

	return distances
}

func (m IntMatrix) IsBottomRightPoint(p Point) bool {
	x, y := p.Coordinates()
	return y == len(m)-1 && x == len(m[0])-1
}

func (m IntMatrix) SetValue(p Point, value int) {
	x, y := p.Coordinates()
	m[y][x] = value
}

func (m IntMatrix) UnsafeValueAt(p Point) int {
	x, y := p.Coordinates()
	return m[y][x]
}

func (m IntMatrix) ValueAt(p Point) (int, error) {
	if !m.IsValidPoint(p) {
		return -1, errors.New("invalid")
	}

	return m.UnsafeValueAt(p), nil
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

func (m IntMatrix) NonDiagonalNeighbors(p Point) []Point {
	return m.NonDiagonalNeighborsForXY(p.Coordinates())
}

func (m IntMatrix) NonDiagonalNeighborsForXY(x, y int) []Point {
	var neighbors []Point
	for _, point := range []Point{
		{x - 1, y},
		{x, y - 1},
		{x, y + 1},
		{x + 1, y},
	} {
		if !m.IsValidPoint(point) {
			continue
		}

		neighbors = append(neighbors, point)
	}
	return neighbors
}

func (m IntMatrix) Neighbors(p Point) []Point {
	return m.NeighborsForXY(p.Coordinates())
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
