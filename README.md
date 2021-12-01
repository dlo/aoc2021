[Advent of Code 2021](https://adventofcode.com/)
=

![go test](https://github.com/dlo/aoc2021/actions/workflows/main.yml/badge.svg)

Howto
-

1. For each problem set, there is an associated day. Use `make init` to generate all the required files for each day. For example:
   
    ```sh
    make init FOLDER=day3
    ```
   
2. Add new case statements in `main.go` to handle CLI parameters.

3. Start writing code in the `${DAY}` folder.

4. When you're ready to test, add unit tests in the `${DAY}` folder. Use `day1/part2_test.go` as a template, and add a line to the `test` rule in `Makefile`. Afterwards, run tests with `make test`. E.g.:

    ```shell
   $ make test
   go test github.com/dlo/aoc2021/day1 -v -cover
   === RUN   TestPart1CalculateNumberOfIncreases
   --- PASS: TestPart1CalculateNumberOfIncreases (0.00s)
   === RUN   TestPart2CalculateNumberOfIncreasesInSlidingWindows
   --- PASS: TestPart2CalculateNumberOfIncreasesInSlidingWindows (0.00s)
   PASS
   coverage: 100.0% of statements
   ok      github.com/dlo/aoc2021/day1     (cached)        coverage: 100.0% of statements
   ```
