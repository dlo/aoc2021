PACKAGE := github.com/dlo/aoc2021/"$(FOLDER)"

.PHONY: init
init:
	mkdir -p "$(FOLDER)/testdata"
	touch "$(FOLDER)/testdata/example_input.txt"
	touch "$(FOLDER)/testdata/input.txt"
	echo "package $(FOLDER)" > "$(FOLDER)"/part1.go
	echo "package $(FOLDER)" > "$(FOLDER)"/part1_test.go
	git add .
	git add -f "$(FOLDER)"/testdata/*txt
	git commit -m "start $(FOLDER) solution"

.PHONY: sweep
sweep:
	go mod tidy

build:
	go build -v ./...

#go test -v ./...
.PHONY: test
test:
	go test github.com/dlo/aoc2021/utils -cover
	go test github.com/dlo/aoc2021/day01 -cover
	go test github.com/dlo/aoc2021/day02 -cover
	go test github.com/dlo/aoc2021/day03 -cover
	go test github.com/dlo/aoc2021/day04 -cover
	go test github.com/dlo/aoc2021/day05 -cover
	go test github.com/dlo/aoc2021/day06 -cover
	go test github.com/dlo/aoc2021/day07 -cover
	go test github.com/dlo/aoc2021/day08 -cover
	go test github.com/dlo/aoc2021/day09 -cover
	go test github.com/dlo/aoc2021/day10 -cover
	go test github.com/dlo/aoc2021/day11 -cover
	go test github.com/dlo/aoc2021/day12 -cover
