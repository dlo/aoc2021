package day8

import (
	"github.com/dlo/aoc2021/utils"
	"strings"
)

type PatternAndValue struct {
	Patterns []string
	Value    []string
}

func SplitBySpaces(s string) []string {
	return strings.Split(s, " ")
}

func (pv *PatternAndValue) UniqueDigitCount() int {
	count := 0
	for _, value := range pv.Value {
		switch len(value) {
		case 2, 3, 4, 7:
			count++

		default:
			continue
		}
	}
	return count
}

func CountUniqueDigitsInValues(patternsAndValues []PatternAndValue) int {
	count := 0
	for _, patternAndValue := range patternsAndValues {
		count += patternAndValue.UniqueDigitCount()
	}
	return count
}

func ParseDigitsAndOutputValues(filename string) []PatternAndValue {
	lines := utils.ReadLinesFromFile(filename)
	var patternsAndValues []PatternAndValue

	var pattern []string
	for _, line := range lines {
		length := len(line)
		if length == 0 {
			break
		}

		if line[length-1] == '|' {
			pattern = SplitBySpaces(line[:length-1])
		} else {
			pattern = []string{}
		}

		var values []string
		var patternAndValue PatternAndValue
		if len(pattern) > 0 {
			values = SplitBySpaces(line)
			patternAndValue = PatternAndValue{pattern, values}
			pattern = []string{}
		} else {
			pos := strings.Index(line, " | ")
			pattern = SplitBySpaces(line[:pos])
			values = SplitBySpaces(line[pos+3:])
			patternAndValue = PatternAndValue{pattern, values}
		}

		patternsAndValues = append(patternsAndValues, patternAndValue)
	}

	return patternsAndValues
}
