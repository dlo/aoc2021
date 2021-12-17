package day14

import "fmt"

type FastPolymerGenerationInstructions struct {
	rules PairInsertionRules
	pairs map[string]int
}

func (instructions PolymerGenerationInstructions) NewFastInstructions() FastPolymerGenerationInstructions {
	pairs := make(map[string]int)
	for _, pair := range instructions.TemplatePairs() {
		pairs[pair]++
	}

	return FastPolymerGenerationInstructions{instructions.rules, pairs}
}

func (instructions FastPolymerGenerationInstructions) Println() {
	for k, v := range instructions.pairs {
		for i := 0; i < v; i++ {
			fmt.Print(k)
		}
	}
	fmt.Println()
}

func (instructions FastPolymerGenerationInstructions) DifferenceOfMostCommonElementFromLeastCommonElement() int {
	counts := make(map[rune]int)
	for k, v := range instructions.pairs {
		runes := []rune(k)
		_, snd := runes[0], runes[1]
		counts[snd] += v
	}

	mostCommonCount := counts['N']
	leastCommonCount := counts['N']
	for _, v := range counts {
		if v > mostCommonCount {
			mostCommonCount = v
		} else if v < leastCommonCount {
			leastCommonCount = v
		}
	}

	return mostCommonCount - leastCommonCount
}

func (instructions *FastPolymerGenerationInstructions) StepCount(iterations int) {
	for i := 0; i < iterations; i++ {
		instructions.Step()
	}
}

func (instructions *FastPolymerGenerationInstructions) Step() {
	newPairs := make(map[string]int)
	for pair, count := range instructions.pairs {
		r := instructions.rules[pair]
		newPairs[string(pair[0])+string(r)] += count
		newPairs[string(r)+string(pair[1])] += count
	}
	instructions.pairs = newPairs
}
