package day14

import (
	"errors"
	"fmt"
	"github.com/dlo/aoc2021/utils"
	"strings"
)

type PairInsertionRules map[string]rune
type PolymerTemplate string
type PolymerGenerationInstructions struct {
	polymer PolymerTemplate
	rules   PairInsertionRules
}
type RawPairInsertionRule string

func (instructions PolymerGenerationInstructions) Println() {
	fmt.Println(instructions.polymer)
}

func (r RawPairInsertionRule) Parse() (string, rune, error) {
	str := string(r)
	if strings.Index(str, "->") == -1 {
		return "", ' ', errors.New("invalid")
	}

	components := strings.Split(str, " -> ")
	return components[0], []rune(components[1])[0], nil
}

func (instructions PolymerGenerationInstructions) TemplatePairs() []string {
	result := make([]string, 0)
	for i := 0; i < len(instructions.polymer)-1; i++ {
		result = append(result, string(instructions.polymer[i:i+2]))
	}
	return result
}

func (instructions *PolymerGenerationInstructions) StepCount(iterations int) {
	for i := 0; i < iterations; i++ {
		instructions.Step()
	}
}

func (instructions *PolymerGenerationInstructions) Step() {
	var template strings.Builder
	var lastPair uint8
	for _, pair := range instructions.TemplatePairs() {
		template.WriteByte(pair[0])
		template.WriteRune(instructions.rules[pair])
		lastPair = pair[1]
	}
	template.WriteByte(lastPair)
	instructions.polymer = PolymerTemplate(template.String())
}

func (instructions PolymerGenerationInstructions) DifferenceOfMostCommonElementFromLeastCommonElement() int {
	elements := make(map[rune]int)
	for _, element := range []rune(instructions.polymer) {
		elements[element]++
	}

	mostCommonCount := elements['N']
	leastCommonCount := elements['N']
	for _, v := range elements {
		if v > mostCommonCount {
			mostCommonCount = v
		} else if v < leastCommonCount {
			leastCommonCount = v
		}
	}

	return mostCommonCount - leastCommonCount
}

func ImportRules(filename string) PolymerGenerationInstructions {
	lines := utils.ReadLinesFromFile(filename)
	template := PolymerTemplate(lines[0])
	rules := make(PairInsertionRules)
	for _, line := range lines[2:] {
		rawRule := RawPairInsertionRule(line)
		input, output, err := rawRule.Parse()
		if err != nil {
			continue
		}

		rules[input] = output
	}

	return PolymerGenerationInstructions{template, rules}
}
