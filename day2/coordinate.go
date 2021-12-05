package day2

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) ProductOfCoordinates() int {
	return c.X * c.Y
}

func (c *Coordinate) ProcessCommand(command Command) {
	switch command.direction {
	case Up:
		c.Y += command.distance

	case Down:
		c.Y -= command.distance

	case Forward:
		c.X += command.distance
	}
}
