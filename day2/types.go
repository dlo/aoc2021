package day2

type Direction int

const (
	Up Direction = iota
	Down
	Forward
)

type Instruction struct {
	direction Direction
	distance  int
}

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) ProductOfCoordinates() int {
	return c.x * c.y
}

type TestCase struct {
	filename string
	want     int
}
