package day05

import (
	"github.com/dlo/aoc2021/day02"
	"github.com/dlo/aoc2021/utils"
)

type LineSegment struct {
	p1 day02.Coordinate
	p2 day02.Coordinate
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
