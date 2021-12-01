PACKAGE := github.com/dlo/aoc2021/"$(FOLDER)"

.PHONY: init
init:
	mkdir "$(FOLDER)"
	cd "$(FOLDER)" && go mod init "$(PACKAGE)"
	go mod edit -require="$(PACKAGE)"@v0.0.0
	go mod edit -replace="$(PACKAGE)"=./"$(FOLDER)"

.PHONY: sweep
sweep:
	go mod tidy

build:
	go build -v ./...

.PHONY: test
test:
	go test -v ./...
	go test github.com/dlo/aoc2021/day1 -v -cover