package day2

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) ProductOfCoordinates() int {
	return c.x * c.y
}

func (c *Coordinate) ProcessCommand(command Command) {
	switch command.direction {
	case Up:
		c.y += command.distance

	case Down:
		c.y -= command.distance

	case Forward:
		c.x += command.distance
	}
}
