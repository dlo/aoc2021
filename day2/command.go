package day2

import (
	"github.com/dlo/aoc2021/utils"
	"strconv"
	"strings"
)

type Command struct {
	direction Direction
	distance  int
}

func Part1ParseCommands(filename string) []Command {
	lines := utils.LinesFromFile(filename)
	var instructions []Command
	for _, line := range lines {
		idx := strings.IndexRune(line, ' ')
		direction := DirectionFromRawDirection(line[:idx])
		distance, _ := strconv.Atoi(line[idx+1:])
		instructions = append(instructions, Command{direction, distance})
	}

	return instructions
}
