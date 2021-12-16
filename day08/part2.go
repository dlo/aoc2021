package day08

import (
	"sort"
	"strings"
)

type Number int
type Letters string
type Pattern string
type Segment int
type SegmentMapping map[Pattern][]Number

type SortableRunes []rune

func (r SortableRunes) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r SortableRunes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r SortableRunes) Len() int {
	return len(r)
}

func (p Pattern) Contains(other Pattern) bool {
	for _, r := range []rune(other) {
		if !strings.ContainsRune(string(p), r) {
			return false
		}
	}
	return true
}

func (p *Pattern) Sort() {
	var ints []int
	for _, b := range []byte(*p) {
		ints = append(ints, int(b))
	}
	sort.Ints(ints)

	var sb strings.Builder
	for _, i := range ints {
		sb.WriteByte(byte(i))
	}

	*p = Pattern(sb.String())
}

func (sm SegmentMapping) DecodeInput(values []string) int {
	result := 0
	for _, value := range values {
		pattern := Pattern(value)
		pattern.Sort()

		result *= 10
		result += int(sm[pattern][0])
	}

	return result
}

func (pvs PatternsAndValues) SumDecodedInputs() int {
	sum := 0
	for _, pv := range pvs {
		mapping := GenerateSegmentMapping(pv.Patterns)
		sum += mapping.DecodeInput(pv.Value)
	}
	return sum
}

func GenerateSegmentMapping(rawPatterns []string) SegmentMapping {
	potentials := map[int][]Number{
		2: {1},
		3: {7},
		4: {4},
		5: {2, 3, 5},
		6: {0, 6, 9},
		7: {8},
	}

	patternsToNumbers := map[Pattern][]Number{}
	lengthsToPatterns := map[int][]Pattern{}
	for _, rawPattern := range rawPatterns {
		pattern := Pattern(rawPattern)
		pattern.Sort()
		patternsToNumbers[pattern] = potentials[len(pattern)]

		lengths := lengthsToPatterns[len(pattern)]
		lengthsToPatterns[len(pattern)] = append(lengths, pattern)
	}

	one := lengthsToPatterns[2][0]
	four := lengthsToPatterns[4][0]
	var six Pattern

	for _, pattern := range lengthsToPatterns[6] {
		if pattern.Contains(one) {
			if pattern.Contains(four) {
				patternsToNumbers[pattern] = []Number{9}
			} else {
				patternsToNumbers[pattern] = []Number{0}
			}
		} else {
			six = pattern
			patternsToNumbers[pattern] = []Number{6}
		}
	}

	for _, pattern := range lengthsToPatterns[5] {
		if pattern.Contains(one) {
			patternsToNumbers[pattern] = []Number{3}
		} else if six.Contains(pattern) {
			patternsToNumbers[pattern] = []Number{5}
		} else {
			patternsToNumbers[pattern] = []Number{2}
		}
	}

	return patternsToNumbers
}
