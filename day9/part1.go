package day9

import (
	"fmt"
	"github.com/dlo/aoc2021/utils"
)

type HeightMap [][]int
type Point [2]int
type PointWithRisk struct {
	point Point
	risk  int
}
type PointsWithRisk struct {
	Points []PointWithRisk
	Risk   int
}

func (hm HeightMap) IsLowPoint(x, y int) bool {
	value := hm[y][x]
	if y > 0 && hm[y-1][x] <= value ||
		x > 0 && hm[y][x-1] <= value ||
		y < len(hm)-1 && hm[y+1][x] <= value ||
		x < len(hm[0])-1 && hm[y][x+1] <= value {
		return false
	}

	return true
}

func (hm HeightMap) Println() {
	for y := range hm {
		fmt.Println(hm[y])
	}
}

func (hm HeightMap) LowPoints() PointsWithRisk {
	result := PointsWithRisk{[]PointWithRisk{}, 0}
	for y := range hm {
		for x := range hm[y] {
			if hm.IsLowPoint(x, y) {
				risk := hm[y][x] + 1
				result.Risk += risk
				result.Points = append(result.Points, PointWithRisk{[2]int{x, y}, risk})
			}
		}
	}
	return result
}

func GenerateHeightMap(filename string) HeightMap {
	lines := utils.ReadLinesFromFile(filename)
	var heightmap [][]int
	for y, line := range lines {
		heightmap = append(heightmap, []int{})
		for _, r := range []rune(line) {
			heightmap[y] = append(heightmap[y], int(r-'0'))
		}
	}
	return heightmap
}
