package day05

import "fmt"

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
