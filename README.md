[Advent of Code 2021](https://adventofcode.com/)
=

![go test](https://github.com/dlo/aoc2021/actions/workflows/main.yml/badge.svg)

> Advent of Code is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like.

This repo contains my solutions for this year's Advent of Code problem sets, written in Go. No one-liners here—we're doing this the hard way! I've organized each day's problem sets into Go modules and wrote a simple runner in `main.go` to quickly iterate on the solutions.

I haven't used Go in a while, so I'm relearning the module system. I have a feeling that this project can be better organized. Suggestions are welcome :).

I've also included unit tests which run on every `git push` via GitHub Actions—this makes the badge above all nice and green.

Howto
-

1. Each day has an associated problem set. Use `make init` to generate all the required files for each day. For example:

    ```sh
    make init FOLDER=day3
    ```

2. Add new case statements in `main.go`.

3. Start writing code in the `${DAY}` folder.

4. When you're ready to test, add unit tests in the `${DAY}` folder. Use `day1/part2_test.go` as a template, and add a line to the `test` rule in `Makefile`. Afterwards, run tests with `make test`. E.g.:

    ```shell
   $ make test
   go test github.com/dlo/aoc2021/day1 -cover
   ok      github.com/dlo/aoc2021/day1     0.234s  coverage: 100.0% of statements
   go test github.com/dlo/aoc2021/utils -cover
   ok      github.com/dlo/aoc2021/utils    (cached)        coverage: 12.1% of statements
   ```
