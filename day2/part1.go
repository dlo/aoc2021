package day2

import (
	"github.com/dlo/aoc2021/utils"
	"strconv"
	"strings"
)

func Part1DirectionFromRawDirection(raw string) Direction {
	switch raw {
	case "up":
		return Up
	case "down":
		return Down
	default:
		return Forward
	}
}

func Part1ParseInstructions(filename string) []Instruction {
	lines := utils.LinesFromFile(filename)
	var instructions []Instruction
	for _, line := range lines {
		idx := strings.IndexRune(line, ' ')
		direction := Part1DirectionFromRawDirection(line[:idx])
		distance, _ := strconv.Atoi(line[idx+1:])
		instructions = append(instructions, Instruction{direction, distance})
	}

	return instructions
}

func Part1CalculatePosition(instructions []Instruction) Coordinate {
	position := Coordinate{0, 0}
	for _, instruction := range instructions {
		switch instruction.direction {
		case Up:
			position.y += instruction.distance

		case Down:
			position.y -= instruction.distance

		case Forward:
			position.x += instruction.distance
		}
	}

	return position
}
