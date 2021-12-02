package day2

func (s *Status) Part2ProcessInstruction(instruction Instruction) {
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

func Part2CalculatePosition(instructions []Instruction) Coordinate {
	status := Status{Coordinate{0, 0}, 0}
	for _, instruction := range instructions {
		status.Part2ProcessInstruction(instruction)
	}

	return status.coordinate
}
