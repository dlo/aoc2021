package day9

import (
	"errors"
	"fmt"
)

// y, x -> y, x -> risk
type RealBasinMap [][]PointWithRisk
type BasinMap map[Point]PointWithRisk
type BlahBLAHBLAH struct {
	basins [][]Point
	seen   map[Point]bool
}
type BasinV4 [][]Point

func (p Point) Coordinates() (int, int) {
	return p[0], p[1]
}

func (bm BasinMap) Println() {
	for y := range bm {
		fmt.Println(y, "=>", bm[y])
	}
}

func (bm BasinMap) InitializeWithHeightMap(hm HeightMap) {
	for y := range hm {
		for x := range hm[y] {
			point := [2]int{x, y}
			risk, _ := hm.RiskAtPoint(point)
			bm[point] = PointWithRisk{point, risk}
		}
	}
}

func (rbm RealBasinMap) InitializeWithHeightMap(hm HeightMap) {
	for y := range hm {
		for x := range hm[y] {
			point := [2]int{x, y}
			risk, _ := hm.RiskAtPoint(point)
			rbm[y][x] = PointWithRisk{point, risk}
		}
	}
}

func (hm HeightMap) GenerateBasins() {
	storage := BlahBLAHBLAH{}
	storage.seen = map[Point]bool{}

	// Go through all the points again and tally up the basin sizes, multiply the largest three
	for y := range hm {
		for _, x := range hm[y] {
			point := [2]int{x, y}
			risk, _ := hm.RiskAtPoint(point)
			if risk == 9 {
				continue
			}

			lowest := pointsToBasins.LowestPointForPoint(hm, point)
			pointsToBasins[y][x] = PointWithRisk{lowest, risk}
		}
	}

	fmt.Println(pointsToBasins)
}

func (hm HeightMap) RiskAtPoint(point Point) (int, error) {
	x, y := point.Coordinates()
	if hm.IsValidPoint(point) {
		return hm[y][x], nil
	}

	return -1, errors.New("invalid point")
}

func (rbm RealBasinMap) LowestPointForPoint(hm HeightMap, point Point) Point {
	if hm.IsLowPoint(point) {
		return point
	}

	value, err := hm.RiskAtPoint(point)
	if err != nil {
		panic("ahh")
	}

	x, y := point.Coordinates()
	lowest := value
	candidates := [4]Point{[2]int{x - 1, y}, [2]int{x + 1, y}, [2]int{x, y - 1}, [2]int{x, y + 1}}
	var result Point
	for _, candidate := range candidates {
		if candidateValue, err := hm.RiskAtPoint(candidate); candidateValue <= lowest && err == nil {
			lowest = candidateValue
			result = candidate
		}
	}

	return rbm.LowestPointForPoint(hm, result)
}
