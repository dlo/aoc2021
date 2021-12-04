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
	go test github.com/dlo/aoc2021/day1 -v -cover
	go test github.com/dlo/aoc2021/day2 -v -cover
	go test github.com/dlo/aoc2021/day3 -v -cover
	go test github.com/dlo/aoc2021/utils -v -cover