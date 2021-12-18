package main

import (
	"flag"
	"fmt"
	"github.com/dlo/aoc2021/day01"
	"github.com/dlo/aoc2021/day02"
	"github.com/dlo/aoc2021/day03"
	"github.com/dlo/aoc2021/day04"
	"github.com/dlo/aoc2021/day05"
	"github.com/dlo/aoc2021/day06"
	"github.com/dlo/aoc2021/day07"
	"github.com/dlo/aoc2021/day08"
	"github.com/dlo/aoc2021/day09"
	"github.com/dlo/aoc2021/day10"
	"github.com/dlo/aoc2021/day11"
	"github.com/dlo/aoc2021/day12"
	"github.com/dlo/aoc2021/day13"
	"github.com/dlo/aoc2021/day14"
	"github.com/dlo/aoc2021/day15"
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

	fmt.Printf("Day %d, ", day)
	switch day {
	case 1:
		items := utils.IntegerLinesFromFile("day1/day1_input.txt")

		switch part {
		case 1:
			fmt.Println("Part One: ", day01.Part1CountIncreases(items))

		case 2:
			fmt.Println("Part Two: ", day01.Part1CountIncreases(day01.Part2GenerateSlidingWindows(items)))
		}

	case 2:
		commands := day02.Part1ParseCommands("day2/testdata/input.txt")

		switch part {
		case 1:
			position := day02.Part1CalculateCoordinateFromCommands(commands)
			fmt.Println("Part One: ", position.ProductOfCoordinates())

		case 2:
			position := day02.Part2CalculateCoordinateFromCommands(commands)
			fmt.Println("Part Two: ", position.ProductOfCoordinates())
		}
	case 3:
		reportValues, length := utils.BinaryLinesFromFile("day3/testdata/input.txt")
		report := day03.DiagnosticReport{Numbers: reportValues, Length: length}

		switch part {
		case 1:
			fmt.Println("Part One: ", int(report.CalculatePowerConsumption()))
		case 2:
			fmt.Println("Part Two: ", int(report.CalculateLifeSupportRating()))
		}
	case 4:
		result := day04.ParseBingoCardDataFromFile("day4/testdata/input.txt")
		switch part {
		case 1:
			fmt.Println(result.FindWinningScore())

		case 2:
			fmt.Println(result.ScoreOfLastWinningBoard())
		}

	case 5:
		measurements := day05.ParseHydrothermalVentMeasurements("day5/testdata/example_input.txt")

		switch part {
		case 1:
			measurements.Process(false)
			fmt.Println("Part One:", measurements.CountIntersections())

		case 2:
			measurements.Process(true)
			fmt.Println("Part One:", measurements.CountIntersections())
		}

	case 6:
		ages := day06.ParseFishFile("day6/testdata/input.txt")
		switch part {
		case 1:
			size := day06.CalculateSchoolSizeAfterNDays(ages, 80)
			fmt.Println("Part One:", size)

		case 2:
			size := day06.CalculateSchoolSizeAfterNDays(ages, 256)
			fmt.Println("Part Two:", size)
		}

	case 7:
		positions := day07.ParseHorizontalPositions("day7/testdata/input.txt")
		switch part {
		case 1:
			position := day07.CalculateCheapestPosition(positions, day07.Part1CalculateFuelUsageForDelta)
			fmt.Println("Part One:", position)

		case 2:
			position := day07.CalculateCheapestPosition(positions, day07.Part2CalculateFuelUsageForDelta)
			fmt.Println("Part Two:", position)
		}

	case 8:
		values := day08.ParseDigitsAndOutputValues("day8/testdata/input.txt")
		switch part {
		case 1:
			break

		case 2:
			fmt.Println("Part Two:", values.SumDecodedInputs())
		}

	case 9:
		heightMap := day09.GenerateHeightMap("day9/testdata/input.txt")
		switch part {
		case 1:
			points := heightMap.LowPoints()
			fmt.Println("Part One:", points.TotalRisk())

		case 2:
			basin := day09.BasinHolder{}
			basin.Initialize(heightMap)
			getBasinsResult := basin.GetBasins()
			basin.Basin.Println()
			fmt.Println("Part Two:", getBasinsResult.ProductOfSizesOfThreeLargestBasins())
		}

	case 10:
		subsystem := day10.GenerateRawSubsystemFromFile("day10/testdata/input.txt")
		switch part {
		case 1:
			fmt.Println("Part One:", subsystem.SyntaxErrorScore())

		case 2:
			fmt.Println("Part Two:", subsystem.CompletionScore())
		}

	case 11:
		grid := day11.ReadOctopusGrid("day11/testdata/input.txt")
		switch part {
		case 1:
			fmt.Print("Part 1")
			fmt.Println()
			grid.StepByCount(100)
			grid.Println()

		case 2:
			fmt.Print("Part 2")
			fmt.Println()
			//cavern.StepByCount(193)
			fmt.Println(grid.FirstSynchronousStep())
			grid.Println()
		}

	case 12:
		caveMap := day12.ImportCaveMap("day12/testdata/input.txt")
		switch part {
		case 1:
			fmt.Println("Part 1:", caveMap.CountUniquePaths(false))

		case 2:
			fmt.Println("Part 2:", caveMap.CountUniquePaths(true))
		}

	case 13:
		grid := day13.ImportGrid("day13/testdata/input.txt")
		switch part {
		case 1:
			fmt.Println("Part 1:", grid.Fold().CountDots())

		case 2:
			grid = grid.FoldAll()
			fmt.Println("Part 2:")
			fmt.Println()
			grid.Println()
		}

	case 14:
		rules := day14.ImportRules("day14/testdata/input.txt")
		switch part {
		case 1:
			rules.StepCount(10)
			fmt.Println("Part 1: ", rules.DifferenceOfMostCommonElementFromLeastCommonElement())

		case 2:
			fastRules := rules.NewFastInstructions()
			fastRules.StepCount(40)
			fmt.Println("Part 2: ", fastRules.DifferenceOfMostCommonElementFromLeastCommonElement())
		}

	case 15:
		cavern := day15.ImportCavernData("day15/testdata/input.txt")
		switch part {
		case 1:
			fmt.Println("Part 1: ")
			cavern.CalculateRiskSums(utils.Point{0, 0})
			fmt.Println(cavern.Risk())
			cavern.PrintRisk()

		case 2:
		}
	}
}
