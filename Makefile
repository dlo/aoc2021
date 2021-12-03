PACKAGE := github.com/dlo/aoc2021/"$(FOLDER)"

.PHONY: init
init:
	mkdir "$(FOLDER)"
	cd "$(FOLDER)" && go mod init "$(PACKAGE)" && go get github.com/stretchr/testify/assert
	go get "$(PACKAGE)"
	go mod edit -require="$(PACKAGE)"@v0.0.0
	go mod edit -replace="$(PACKAGE)"=./"$(FOLDER)"

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