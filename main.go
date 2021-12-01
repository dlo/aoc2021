package main

import (
	"flag"
	"fmt"
	"github.com/dlo/aoc2021/day1"
	"github.com/dlo/aoc2021/utils"
)

var day, part int

func init() {
	flag.IntVar(&day, "day", 1, "The day")
	flag.IntVar(&part, "part", 1, "The part")
}

/// Usage: go run main.go --day 1 --part 1
func main() {
	flag.Parse()

	switch day {
	case 1:
		fmt.Print("Day One, ")
		switch part {
		case 1:
			fmt.Print("Part One: ")
			items := utils.LinesFromFile("day1/day1_input.txt")
			fmt.Println(day1.Part1CountIncreases(items))

		case 2:
			fmt.Print("Part Two: ")
			items := utils.LinesFromFile("day1/day1_input.txt")
			fmt.Println(day1.Part2CountIncreasesInSlidingWindows(items))
		}

	case 2:
		switch part {
		case 1:
		case 2:
			break
		}
	}
}
