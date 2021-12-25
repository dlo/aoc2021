package day17

import (
	"errors"
	"github.com/dlo/aoc2021/utils"
	"strconv"
	"strings"
)

type Range [2]int
type Solution [3]int

type SolutionGroup struct {
	solutions []Solution
}

type TargetArea struct {
	x Range
	y Range
}

type Position struct {
	dx   int
	dy   int
	x    int
	y    int
	maxY int
}

func (s Solution) MaxY() int {
	return s[2]
}

func (s SolutionGroup) Count() int {
	return len(s.solutions)
}

func (s SolutionGroup) MaxY() int {
	maxY := 0
	for _, solution := range s.solutions {
		maxY = utils.Max(solution.MaxY(), maxY)
	}
	return maxY
}

func NewPosition(dx, dy int) Position {
	return Position{dx, dy, 0, 0, 0}
}

func (r Range) Max() int {
	return r[1]
}

func (r Range) Min() int {
	return r[0]
}

func (r Range) ContainsValue(i int) bool {
	return i >= r.Min() && i <= r.Max()
}

func (r Range) LessThanValue(i int) bool {
	return i > r.Max()
}

func (p *Position) Step() {
	p.x += p.dx
	p.y += p.dy
	p.maxY = utils.Max(p.y, p.maxY)

	if p.dx > 0 {
		p.dx--
	} else if p.dx < 0 {
		p.dx++
	}

	p.dy--
}

func (p *Position) StepUntilDone(a TargetArea) (bool, bool) {
	for !p.WithinTargetArea(a) && !p.PassedTargetArea(a) {
		p.Step()
	}

	return p.WithinTargetArea(a), p.PassedTargetArea(a)
}

func (p Position) WithinTargetArea(a TargetArea) bool {
	return a.x.ContainsValue(p.x) && a.y.ContainsValue(p.y)
}

func (p Position) PassedTargetArea(a TargetArea) bool {
	return p.x > a.x.Max() || p.y < a.y.Min()
}

func ParseInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func ParseRange(s string) (Range, error) {
	idx := strings.Index(s, "=")
	if idx == -1 {
		return [2]int{0, 0}, errors.New("invalid")
	}

	components := strings.Split(s[idx+1:], "..")
	return [2]int{ParseInt(components[0]), ParseInt(components[1])}, nil
}

func (a TargetArea) FindSolution() SolutionGroup {
	group := SolutionGroup{}
	for x := -100; x < 200; x++ {
		for y := -200; y < 200; y++ {
			position := NewPosition(x, y)
			success, _ := position.StepUntilDone(a)
			if success {
				solution := Solution{position.x, position.y, position.maxY}
				group.solutions = append(group.solutions, solution)
			}
		}
	}
	return group
}

func ImportData(filename string) TargetArea {
	line := utils.ReadLinesFromFile(filename)[0]
	ranges := strings.Split(line[13:], ", ")
	xRange, _ := ParseRange(ranges[0])
	yRange, _ := ParseRange(ranges[1])
	return TargetArea{xRange, yRange}
}
