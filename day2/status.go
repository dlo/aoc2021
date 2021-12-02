package day2

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
		s.coordinate.x += instruction.distance
		s.coordinate.y += s.aim * instruction.distance
	}
}
