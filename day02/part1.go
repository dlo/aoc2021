package day2

func Part1CalculateCoordinateFromCommands(commands []Command) Coordinate {
	position := Coordinate{0, 0}
	for _, command := range commands {
		position.ProcessCommand(command)
	}

	return position
}
