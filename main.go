package main

import (
	"flag"
	"fmt"
	"github.com/dlo/aoc2021/day1"
	"github.com/dlo/aoc2021/day2"
	"github.com/dlo/aoc2021/day3"
	"github.com/dlo/aoc2021/day4"
	"github.com/dlo/aoc2021/day5"
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
		items := utils.IntegerLinesFromFile("day1/day1_input.txt")

		fmt.Print("Day One, ")
		switch part {
		case 1:
			fmt.Println("Part One: ", day1.Part1CountIncreases(items))

		case 2:
			fmt.Println("Part Two: ", day1.Part1CountIncreases(day1.Part2GenerateSlidingWindows(items)))
		}
	case 2:
		commands := day2.Part1ParseCommands("day2/testdata/input.txt")

		fmt.Print("Day Two, ")
		switch part {
		case 1:
			position := day2.Part1CalculateCoordinateFromCommands(commands)
			fmt.Println("Part One: ", position.ProductOfCoordinates())

		case 2:
			position := day2.Part2CalculateCoordinateFromCommands(commands)
			fmt.Println("Part Two: ", position.ProductOfCoordinates())
		}
	case 3:
		reportValues, length := utils.BinaryLinesFromFile("day3/testdata/input.txt")
		report := day3.DiagnosticReport{Numbers: reportValues, Length: length}

		fmt.Print("Day Three, ")
		switch part {
		case 1:
			fmt.Println("Part One: ", int(report.CalculatePowerConsumption()))
		case 2:
			fmt.Println("Part Two: ", int(report.CalculateLifeSupportRating()))
		}
	case 4:
		result := day4.ParseBingoCardDataFromFile("day4/testdata/input.txt")
		fmt.Print("Day Four, ")
		switch part {
		case 1:
			fmt.Println(result.FindWinningScore())

		case 2:
			fmt.Println(result.ScoreOfLastWinningBoard())
		}

	case 5:
		measurements := day5.ParseHydrothermalVentMeasurements("day5/testdata/example_input.txt")

		fmt.Print("Day Five, ")
		switch part {
		case 1:
			measurements.Process(false)
			fmt.Println("Part One ", measurements.CountIntersections())

		case 2:
			measurements.Process(true)
			fmt.Println("Part One ", measurements.CountIntersections())
		}
	}
}
