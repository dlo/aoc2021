package day9

import (
	"fmt"
	"github.com/dlo/aoc2021/utils"
)

type Risk int
type HeightMap [][]int
type Point [3]int
type Points []Point

func (p Points) TotalRisk() int {
	risk := 0
	for _, point := range p {
		risk += point.Risk()
	}
	return risk
}

func (p Point) Coordinates() (x, y int) {
	return p[0], p[1]
}

func (p Point) Height() (height int) {
	return p[2]
}

func (p Point) Risk() int {
	return p.Height() + 1
}

func (hm HeightMap) IsValidPoint(point Point) bool {
	x, y := point.Coordinates()
	return y >= 0 && x >= 0 && y < len(hm) && x < len(hm[0])
}

func (hm HeightMap) IsLowPoint(point Point) bool {
	x, y := point.Coordinates()
	height := point.Height()
	if y > 0 && hm[y-1][x] <= height ||
		x > 0 && hm[y][x-1] <= height ||
		y < len(hm)-1 && hm[y+1][x] <= height ||
		x < len(hm[0])-1 && hm[y][x+1] <= height {
		return false
	}

	return true
}

func (hm HeightMap) Println() {
	for y := range hm {
		fmt.Println(hm[y])
	}
}

func (hm HeightMap) LowPoints() Points {
	var points []Point
	for y := range hm {
		for x := range hm[y] {
			point := Point([3]int{x, y, hm[y][x]})
			if hm.IsLowPoint(point) {
				points = append(points, point)
			}
		}
	}
	return points
}

func GenerateHeightMap(filename string) HeightMap {
	lines := utils.ReadLinesFromFile(filename)
	var heightmap HeightMap
	for y, line := range lines {
		heightmap = append(heightmap, []int{})
		for _, r := range []rune(line) {
			heightmap[y] = append(heightmap[y], int(r-'0'))
		}
	}
	return heightmap
}
