package day12

import (
	"fmt"
	"github.com/dlo/aoc2021/utils"
	"strings"
	"unicode"
)

type CaveMap utils.IntMatrix
type Cave string
type CaveIds map[Cave]int
type CaveStorage struct {
	cm         CaveMap
	ids        CaveIds
	idsToCaves map[int]Cave
}
type RawConnection string
type Visited map[int]bool

func (v Visited) Copy() Visited {
	result := make(Visited)
	for k, v := range v {
		result[k] = v
	}
	return result
}

func (c Cave) IsBig() bool {
	r := []rune(c)[0]
	return unicode.IsUpper(r) && c != "start"
}

func (cs CaveStorage) Println() {
	maxWidth := 0
	for k := range cs.ids {
		if len(k) > maxWidth {
			maxWidth = len(k)
		}
	}

	for y, cave := range cs.idsToCaves {
		format := fmt.Sprintf("%%%dv idx=%%d ", maxWidth)
		fmt.Printf(format, cave, cs.ids[cave])

		for x := range cs.cm[y] {
			point := utils.Point{x, y}
			fmt.Print(utils.IntMatrix(cs.cm).ValueAt(point))
		}

		fmt.Println()
	}
}

func (c RawConnection) ParseConnection() (fst Cave, snd Cave) {
	input := string(c)
	idx := strings.IndexRune(input, '-')
	return Cave(input[:idx]), Cave(input[idx+1:])
}

func (cs CaveStorage) CountUniquePaths(y int) int {
	startId := cs.ids["start"]
	return cs.InternalCountUniquePaths(startId, []Cave{"start"}, make(map[int]bool))
}

func (cs CaveStorage) InternalCountUniquePaths(y int, path []Cave, visited Visited) int {
	//y = start
	cave := cs.idsToCaves[y]
	if visited[y] && !cave.IsBig() {
		return 0
	}

	visited[y] = true
	path = append(path, cave)

	sum := 0
	for x, hasConnection := range cs.cm[y] {
		if hasConnection == 0 {
			continue
		}

		name := cs.idsToCaves[x]
		if name == "end" {
			path = append(path, name)
			//fmt.Println(path)
			sum += 1
		} else {
			sum += cs.InternalCountUniquePaths(x, path, visited.Copy())
		}
	}
	return sum
}

func ImportCaveMap(filename string) CaveStorage {
	lines := utils.ReadLinesFromFile(filename)
	caveIds := make(map[Cave]int)
	idsToCaves := make(map[int]Cave)
	i := 0
	for _, line := range lines {
		fst, snd := RawConnection(line).ParseConnection()
		if _, ok := caveIds[fst]; !ok {
			caveIds[fst] = i
			idsToCaves[i] = fst
			i++
		}

		if _, ok := caveIds[snd]; !ok {
			caveIds[snd] = i
			idsToCaves[i] = snd
			i++
		}
	}

	caveMap := CaveMap(utils.NewIntMatrix(len(caveIds), len(caveIds)))
	for _, line := range lines {
		fst, snd := RawConnection(line).ParseConnection()
		fstId := caveIds[fst]
		sndId := caveIds[snd]
		caveMap[fstId][sndId] = 1
		caveMap[sndId][fstId] = 1
	}

	visits := make([]int, i)
	visits[0] = 1
	return CaveStorage{caveMap, caveIds, idsToCaves}
}
