Advent of Code 2021
=

Howto
-

1. For each problem set, there is an associated day. Use `make init` recipe to generate all the required files.
   
    ```sh
    make init FOLDER=day3
    ```
   
2. Add new case statements in `main.go` to handle CLI parameters.

3. Start writing code in the `${DAY}` folder.

4. When you're ready to test, add a test file in the `${DAY}` folder. Use `day1/part2_test.go` as a template, and add the test to the `test` section in the `Makefile`. Run tests with:

    ```shell
   make test
   ```
