package day5

import (
	"github.com/dlo/aoc2021/day2"
	"github.com/dlo/aoc2021/utils"
)

type LineSegment struct {
	p1 day2.Coordinate
	p2 day2.Coordinate
}

func (s *LineSegment) MinX() int {
	return utils.Min(s.p1.X, s.p2.X)
}

func (s *LineSegment) MinY() int {
	return utils.Min(s.p1.Y, s.p2.Y)
}

func (s *LineSegment) MaxX() int {
	return utils.Max(s.p1.X, s.p2.X)
}

func (s *LineSegment) MaxY() int {
	return utils.Max(s.p1.Y, s.p2.Y)
}
