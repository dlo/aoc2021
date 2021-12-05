package day5

import (
	"fmt"
	"github.com/dlo/aoc2021/day2"
	"github.com/dlo/aoc2021/utils"
	"strconv"
	"strings"
)

type LineSegment struct {
	p1 day2.Coordinate
	p2 day2.Coordinate
}

type HydrothermalVentMeasurements struct {
	Segments []LineSegment
	Grid     [][]int
	maxX     int
	maxY     int
}

func (m *HydrothermalVentMeasurements) PrintGrid() {
	for _, row := range m.Grid {
		fmt.Println(row)
	}
}

func (m *HydrothermalVentMeasurements) Process(countDiagonals bool) {
	for _, segment := range m.Segments {
		if segment.MinX() == segment.MaxX() {
			for y := segment.MinY(); y <= segment.MaxY(); y++ {
				m.Grid[y][segment.MinX()]++
			}
		} else if segment.MinY() == segment.MaxY() {
			for x := segment.MinX(); x <= segment.MaxX(); x++ {
				m.Grid[segment.MinY()][x]++
			}
		} else if countDiagonals {
			isXIncreasing := segment.p1.X < segment.p2.X
			isYIncreasing := segment.p1.Y < segment.p2.Y
			for offset := 0; offset <= segment.MaxX()-segment.MinX(); offset++ {
				if isXIncreasing && isYIncreasing {
					m.Grid[segment.MinY()+offset][segment.MinX()+offset]++
				} else if isXIncreasing && !isYIncreasing {
					m.Grid[segment.MaxY()-offset][segment.MinX()+offset]++
				} else if !isXIncreasing && isYIncreasing {
					m.Grid[segment.MinY()+offset][segment.MaxX()-offset]++
				} else {
					m.Grid[segment.MaxY()-offset][segment.MaxX()-offset]++
				}
			}
		}
	}
}

func (m *HydrothermalVentMeasurements) CountIntersections() int {
	count := 0
	for _, row := range m.Grid {
		for _, value := range row {
			if value > 1 {
				count++
			}
		}
	}
	return count
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (s *LineSegment) MinX() int {
	return Min(s.p1.X, s.p2.X)
}

func (s *LineSegment) MinY() int {
	return Min(s.p1.Y, s.p2.Y)
}

func (s *LineSegment) MaxX() int {
	return Max(s.p1.X, s.p2.X)
}

func (s *LineSegment) MaxY() int {
	return Max(s.p1.Y, s.p2.Y)
}

func CoordinateFromString(input string) day2.Coordinate {
	numbers := strings.Split(input, ",")
	X, _ := strconv.Atoi(numbers[0])
	Y, _ := strconv.Atoi(numbers[1])
	return day2.Coordinate{X: X, Y: Y}
}

func ParseLineSegmentFromLine(line string) LineSegment {
	rawCoordinates := strings.Split(line, " -> ")
	c1 := CoordinateFromString(rawCoordinates[0])
	c2 := CoordinateFromString(rawCoordinates[1])
	return LineSegment{c1, c2}
}

func ParseHydrothermalVentMeasurements(filename string) HydrothermalVentMeasurements {
	lines := utils.ReadLinesFromFile(filename)

	var segments []LineSegment
	maxX := 0
	maxY := 0
	for _, line := range lines {
		segment := ParseLineSegmentFromLine(line)
		if segment.MaxY() > maxY {
			maxY = segment.MaxY()
		}

		if segment.MaxX() > maxX {
			maxX = segment.MaxX()
		}

		segments = append(segments, segment)
	}

	grid := make([][]int, maxY+1)
	for i := range grid {
		grid[i] = make([]int, maxX+1)
	}

	return HydrothermalVentMeasurements{segments, grid, maxX, maxY}
}
