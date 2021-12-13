package day8

import (
	"fmt"
	"strconv"
)

type Number int
type Letters string
type Segment int
type SegmentMapping [7][7]int
type SegmentMappingNew struct {
	mapping    [7][7]int
	potentials map[int][]Number
}

/*
1. Take values.
2. Each value corresponds to a number or numbers. E.g., ab corresponds to 1.
3. For each letter in the value, find the _actual_ segment it could correspond to. For 1, that would be BC.
We would then increment a[B_segment] and a[C_segment] by 1, and b[B_segment] and b[C_segment] by 1.
4. When we display a pattern, e.g., cbdaf, go through each letter, and increment all of the segments, and display
the total count.
*/

const (
	F = iota
	A
	B
	G
	E
	C
	D
)

func InitialPotentialNumbers() map[int][]Number {
	return map[int][]Number{
		2: {1},
		3: {7},
		4: {4},
		5: {2, 3, 5},
		6: {6, 9},
		7: {8},
	}
}

func (s Letters) PotentialNumbers() []Number {
	potentials := map[int][]Number{
		2: {1},
		3: {7},
		4: {4},
		5: {2, 3, 5},
		6: {6, 9},
		7: {8},
	}

	return potentials[len(s)]
}

func (n Number) SegmentsForNumber() []Segment {
	segments := map[int][]Segment{
		1: {B, C},
		2: {A, B, G, E, D},
		3: {A, B, G, C, D},
		4: {F, G, C, D},
		5: {A, F, G, C, D},
		6: {A, F, G, E, C, D},
		7: {A, B, C},
		8: {A, B, C, D, E, F},
	}

	return segments[int(n)]
}

func RuneToSegment(r rune) Segment {
	runes := map[rune]Segment{
		'a': A,
		'b': B,
		'c': C,
		'd': D,
		'e': E,
		'f': F,
		'g': G,
	}

	return runes[r]
}

func (sm *SegmentMapping) ImportPattern(pattern string) {
	letters := Letters(pattern)
	for _, number := range letters.PotentialNumbers() {
		for _, segment := range number.SegmentsForNumber() {
			for _, r := range []rune(pattern) {
				value := RuneToSegment(r)
				sm[value][segment]++
			}
		}
	}
}

func (sm SegmentMapping) GenerateNumber(pattern string) [7]int {
	var output [7]int
	letters := Letters(pattern)
	for _, number := range letters.PotentialNumbers() {
		for _, segment := range number.SegmentsForNumber() {
			output[A] += sm[segment][A]
			output[B] += sm[segment][B]
			output[C] += sm[segment][C]
			output[D] += sm[segment][D]
			output[E] += sm[segment][E]
			output[F] += sm[segment][F]
			output[G] += sm[segment][G]
		}
	}
	return output
}

func (sm SegmentMapping) MostLikelyNumberForValue(value string) Number {
	letters := Letters(value)
	biggest := 0
	var totals [10]int
	var result Number
	for _, potentialNumber := range letters.PotentialNumbers() {
		total := 0
		for i := range sm {
			for _, segment := range potentialNumber.SegmentsForNumber() {
				total += sm[i][segment]
			}
		}

		totals[potentialNumber] += total

		if total > biggest {
			result = potentialNumber
		}
	}

	fmt.Println(value)
	for i := range totals {
		fmt.Print(totals[i], " ")
	}
	fmt.Println()
	fmt.Println()
	return result
}

func (sm SegmentMapping) DecodeNumber(values []string) string {
	var runes string
	for _, value := range values {
		runes = runes + strconv.FormatUint(uint64(sm.MostLikelyNumberForValue(value)), 10)
	}
	fmt.Println(runes)
	return runes
}

func ImportPatterns(patterns []string) SegmentMapping {
	var runes [7][7]int
	var sm SegmentMapping = runes

	for _, pattern := range patterns {
		sm.ImportPattern(pattern)
	}

	return sm
}
