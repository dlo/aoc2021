package day02

type Status struct {
	coordinate Coordinate
	aim        int
}

func (s *Status) ProcessCommand(instruction Command) {
	switch instruction.direction {
	case Down:
		s.aim += instruction.distance

	case Up:
		s.aim -= instruction.distance

	case Forward:
		s.coordinate.X += instruction.distance
		s.coordinate.Y += s.aim * instruction.distance
	}
}
