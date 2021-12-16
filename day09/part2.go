package day9

import (
	"fmt"
	"github.com/fatih/color"
	"sort"
)

type GetBasinsResult map[Point][]Point

type Basin [][]Point

type BasinHolder struct {
	hm    HeightMap
	seen  map[Point]bool
	Basin Basin
}

func AttributeForHeight(height int) color.Attribute {
	m := map[int]color.Attribute{
		0: color.FgBlack,
		1: color.FgBlue,
		2: color.FgMagenta,
		3: color.FgHiMagenta,
		4: color.FgHiBlue,
		5: color.FgYellow,
		6: color.FgHiYellow,
		7: color.FgRed,
		8: color.FgWhite,
		9: color.FgHiWhite,
	}
	return m[height%10]
}

func (b Basin) Println() {
	for j := range b {
		for i := range b[j] {
			point := b[j][i]
			attribute := AttributeForHeight(point.Height())
			color.New(attribute).PrintFunc()(point.Height())
		}
		fmt.Println()
	}
}

func (p Point) NeighborCoordinates() [4][2]int {
	x, y := p.Coordinates()
	return [4][2]int{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
}

func (g *BasinHolder) Initialize(hm HeightMap) {
	g.hm = hm
	g.seen = map[Point]bool{}
	for range hm {
		g.Basin = append(g.Basin, make([]Point, len(hm[0])))
	}

	for y := range hm {
		for x := range hm[y] {
			g.Basin[y][x], _ = hm.PointAt(x, y)
		}
	}

	for _, point := range hm.LowPoints() {
		g.ExpandBasinFromPoint(point)
	}
}

func (g *BasinHolder) ExpandBasinFromPoint(p Point) {
	if g.seen[p] {
		return
	}

	g.seen[p] = true
	neighborCoordinates := p.NeighborCoordinates()
	for i := range neighborCoordinates {
		x, y := neighborCoordinates[i][0], neighborCoordinates[i][1]
		neighbor, err := g.hm.PointAt(x, y)

		if err == nil && neighbor.Height() < 9 {
			g.Basin[y][x] = g.Basin[p.Y()][p.X()]
			g.ExpandBasinFromPoint(neighbor)
		}
	}
}

func (r GetBasinsResult) ProductOfSizesOfThreeLargestBasins() int {
	var sizes []int
	for _, v := range r {
		sizes = append(sizes, len(v))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func (g *BasinHolder) GetBasins() GetBasinsResult {
	results := GetBasinsResult{}
	for y := range g.Basin {
		for x := range g.Basin[y] {
			points := results[g.Basin[y][x]]
			point, _ := g.hm.PointAt(x, y)
			if point.Height() == 9 {
				continue
			}

			if points == nil {
				points = []Point{point}
			} else {
				points = append(points, point)
			}
			results[g.Basin[y][x]] = points
		}
	}
	return results
}
