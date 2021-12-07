PACKAGE := github.com/dlo/aoc2021/"$(FOLDER)"

.PHONY: init
init:
	mkdir -p "$(FOLDER)/testdata"

.PHONY: sweep
sweep:
	go mod tidy

build:
	go build -v ./...

#go test -v ./...
.PHONY: test
test:
	go test github.com/dlo/aoc2021/utils -cover
	go test github.com/dlo/aoc2021/day1 -cover
	go test github.com/dlo/aoc2021/day2 -cover
	go test github.com/dlo/aoc2021/day3 -cover
	go test github.com/dlo/aoc2021/day4 -cover
	go test github.com/dlo/aoc2021/day5 -cover
	go test github.com/dlo/aoc2021/day6 -cover
