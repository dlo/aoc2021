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
	go test github.com/dlo/aoc2021/day1 -cover
	go test github.com/dlo/aoc2021/day2 -cover
	go test github.com/dlo/aoc2021/day3 -cover
	go test github.com/dlo/aoc2021/day4 -cover
	go test github.com/dlo/aoc2021/day5 -cover
	go test github.com/dlo/aoc2021/day6 -cover
	go test github.com/dlo/aoc2021/day7 -cover
	go test github.com/dlo/aoc2021/day8 -cover
