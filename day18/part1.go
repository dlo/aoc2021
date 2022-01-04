package day18

import (
	"fmt"
	"github.com/dlo/aoc2021/utils"
	"strconv"
)

type RawSnailfishNumber string

type SnailfishNumber struct {
	value *int
	left  *SnailfishNumber
	right *SnailfishNumber
}

func (n SnailfishNumber) Magnitude() int {
	if n.value != nil {
		return *n.value
	}

	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func (n SnailfishNumber) String() string {
	if n.value != nil {
		return fmt.Sprintf("%d", *n.value)
	} else {
		return fmt.Sprintf("[%s,%s]", n.left.String(), n.right.String())
	}
}

func (n SnailfishNumber) Println() {
	fmt.Println(n.String())
}

func (r RawSnailfishNumber) SeparatorIndex() int {
	count := 0
	for i, r := range []rune(r) {
		switch r {
		case '[':
			count++
		case ']':
			count--
		case ',':
			if count-1 == 0 {
				return i
			}
		}
	}
	return -1
}

func (r RawSnailfishNumber) Parse() SnailfishNumber {
	idx := r.SeparatorIndex()
	if idx != -1 {
		left := r[1:idx].Parse()
		right := r[idx+1 : len(r)-1].Parse()
		return SnailfishNumber{nil, &left, &right}
	}

	result, _ := strconv.Atoi(string(r))
	return SnailfishNumber{&result, nil, nil}
}

func ImportSnailfishNumbers(filename string) []SnailfishNumber {
	lines := utils.ReadLinesFromFile(filename)
	var numbers []SnailfishNumber
	for _, line := range lines {
		snailfishNumber := RawSnailfishNumber(line).Parse()
		numbers = append(numbers, snailfishNumber)
	}
	return numbers
}
