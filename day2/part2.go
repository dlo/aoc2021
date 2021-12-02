package day2

func Part2CalculatePosition(instructions []Instruction) Coordinate {
	aim := 0
	position := Coordinate{0, 0}
	for _, instruction := range instructions {
		switch instruction.direction {
		case Down:
			aim += instruction.distance

		case Up:
			aim -= instruction.distance

		case Forward:
			position.x += instruction.distance
			position.y += aim * instruction.distance
		}
	}

	return position
}
