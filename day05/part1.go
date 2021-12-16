package day05

import (
	"strconv"
	"strings"

	"github.com/dlo/aoc2021/day2"
	"github.com/dlo/aoc2021/utils"
)

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
