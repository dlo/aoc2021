package day2

func Part2CalculateCoordinateFromCommands(commands []Command) Coordinate {
	status := Status{Coordinate{0, 0}, 0}
	for _, command := range commands {
		status.ProcessCommand(command)
	}

	return status.coordinate
}
