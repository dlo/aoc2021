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
	"github.com/dlo/aoc2021/day16"
	"github.com/dlo/aoc2021/utils"
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
)

var day, part int
var filename string

func init() {
	flag.IntVar(&day, "day", 1, "The day")
	flag.IntVar(&part, "part", 1, "The part")
	flag.StringVar(&filename, "filename", "input.txt", "The filename")
}

/// Usage: go run main.go --day 1 --part 1
func main() {
	flag.Parse()

	fmt.Printf("Day %d, ", day)
	formattedFile := fmt.Sprintf("day%02d/testdata/%s", day, filename)
	switch day {
	case 1:
		items := utils.IntegerLinesFromFile(formattedFile)

		switch part {
		case 1:
			fmt.Println("Part One: ", day01.Part1CountIncreases(items))

		case 2:
			fmt.Println("Part Two: ", day01.Part1CountIncreases(day01.Part2GenerateSlidingWindows(items)))
		}

	case 2:
		commands := day02.Part1ParseCommands(formattedFile)

		switch part {
		case 1:
			position := day02.Part1CalculateCoordinateFromCommands(commands)
			fmt.Println("Part One: ", position.ProductOfCoordinates())

		case 2:
			position := day02.Part2CalculateCoordinateFromCommands(commands)
			fmt.Println("Part Two: ", position.ProductOfCoordinates())
		}
	case 3:
		reportValues, length := utils.BinaryLinesFromFile(formattedFile)
		report := day03.DiagnosticReport{Numbers: reportValues, Length: length}

		switch part {
		case 1:
			fmt.Println("Part One: ", int(report.CalculatePowerConsumption()))
		case 2:
			fmt.Println("Part Two: ", int(report.CalculateLifeSupportRating()))
		}
	case 4:
		result := day04.ParseBingoCardDataFromFile(formattedFile)
		switch part {
		case 1:
			fmt.Println(result.FindWinningScore())

		case 2:
			fmt.Println(result.ScoreOfLastWinningBoard())
		}

	case 5:
		measurements := day05.ParseHydrothermalVentMeasurements(formattedFile)

		switch part {
		case 1:
			measurements.Process(false)
			fmt.Println("Part One:", measurements.CountIntersections())

		case 2:
			measurements.Process(true)
			fmt.Println("Part One:", measurements.CountIntersections())
		}

	case 6:
		ages := day06.ParseFishFile(formattedFile)
		switch part {
		case 1:
			size := day06.CalculateSchoolSizeAfterNDays(ages, 80)
			fmt.Println("Part One:", size)

		case 2:
			size := day06.CalculateSchoolSizeAfterNDays(ages, 256)
			fmt.Println("Part Two:", size)
		}

	case 7:
		positions := day07.ParseHorizontalPositions(formattedFile)
		switch part {
		case 1:
			position := day07.CalculateCheapestPosition(positions, day07.Part1CalculateFuelUsageForDelta)
			fmt.Println("Part One:", position)

		case 2:
			position := day07.CalculateCheapestPosition(positions, day07.Part2CalculateFuelUsageForDelta)
			fmt.Println("Part Two:", position)
		}

	case 8:
		values := day08.ParseDigitsAndOutputValues(formattedFile)
		switch part {
		case 1:
			break

		case 2:
			fmt.Println("Part Two:", values.SumDecodedInputs())
		}

	case 9:
		heightMap := day09.GenerateHeightMap(formattedFile)
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
		subsystem := day10.GenerateRawSubsystemFromFile(formattedFile)
		switch part {
		case 1:
			fmt.Println("Part One:", subsystem.SyntaxErrorScore())

		case 2:
			fmt.Println("Part Two:", subsystem.CompletionScore())
		}

	case 11:
		grid := day11.ReadOctopusGrid(formattedFile)
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
		caveMap := day12.ImportCaveMap(formattedFile)
		switch part {
		case 1:
			fmt.Println("Part 1:", caveMap.CountUniquePaths(false))

		case 2:
			fmt.Println("Part 2:", caveMap.CountUniquePaths(true))
		}

	case 13:
		grid := day13.ImportGrid(formattedFile)
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
		rules := day14.ImportRules(formattedFile)
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
		cavern := day15.ImportCavernData(formattedFile)
		switch part {
		case 1:
			m := cavern.DijkstraPath(utils.Point{0, 0})
			maxX, maxY := m.Size()
			point := utils.Point{maxX - 1, maxY - 1}
			fmt.Println(m.UnsafeValueAt(point))

		case 2:
			fmt.Println()
			nc := cavern.MultiplyCavern()
			m := nc.DijkstraPath(utils.Point{0, 0})
			maxX, maxY := m.Size()
			point := utils.Point{maxX - 1, maxY - 1}
			fmt.Println(m.UnsafeValueAt(point))
		}

		if false {
			s, err := tcell.NewScreen()
			if err != nil {
				log.Fatalf("%+v", err)
			}
			if err := s.Init(); err != nil {
				log.Fatalf("%+v", err)
			}

			// Set default text style
			background := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
			s.SetStyle(background)

			// Clear screen
			s.Clear()
			s.Show()

			quit := func() {
				s.Fini()
				fmt.Println()
				os.Exit(0)
			}

			for {
				ev := s.PollEvent()
				switch ev := ev.(type) {
				case *tcell.EventResize:
					s.Sync()

				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyEscape, tcell.KeyCtrlC:
						quit()

					case tcell.KeyRune:
						switch ev.Rune() {
						case 'q':
							quit()

						default:
							cavern.Display(s)
							s.Show()
						}
					}
				}
			}
		}

	case 16:
		transmissions := day16.ImportBITSTransmission(formattedFile)
		switch part {
		case 1:
			for _, transmission := range transmissions {
				packet, _ := transmission.Process()
				fmt.Println("Part 1:", packet.VersionNumberSum())
			}
		}
		//
		//switch part {
		//case 1:
		//	//fmt.Println("Part 1: ")
		//	cavern.CalculateRiskSums(utils.Point{0, 0})
		//	//fmt.Println(cavern.Risk())
		//	//cavern.PrintRisk()
		//
		//case 2:
		//}
	}
}
