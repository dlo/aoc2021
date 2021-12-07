package utils

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func WriteLinesToFile(values []int, filename string) {
	writer, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func(writer *os.File) {
		_ = writer.Close()
	}(writer)

	csvWriter := csv.NewWriter(writer)
	var stringValues [][]string
	for _, value := range values {
		stringValues = append(stringValues, []string{strconv.Itoa(value)})
	}

	_ = csvWriter.WriteAll(stringValues)
}

func ReadLinesFromFile(filename string) []string {
	var reader io.Reader
	if filename == "stdin" {
		reader = os.Stdin
	} else {
		file, _ := os.Open(filename)
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		reader = file
	}

	var lines []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func LinesFromCSV(filename string) []string {
	var reader io.Reader
	if filename == "stdin" {
		reader = os.Stdin
	} else {
		file, _ := os.Open(filename)
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		reader = file
	}

	csvReader := csv.NewReader(reader)
	var items []string
	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for _, value := range record {
			items = append(items, value)
		}
	}

	return items
}

func IntegerLinesFromFile(filename string) []int {
	lines := LinesFromCSV(filename)
	var items []int
	for _, value := range lines {
		result, _ := strconv.Atoi(value)
		items = append(items, result)
	}
	return items
}

func BinaryLinesFromFile(filename string) ([]uint64, int) {
	lines := LinesFromCSV(filename)
	var items []uint64
	maxLength := 0
	for _, value := range lines {
		if len(value) > maxLength {
			maxLength = len(value)
		}
		result, _ := strconv.ParseInt(value, 2, 0)
		items = append(items, uint64(result))
	}
	return items, maxLength
}

func ParseNumbersFromString(input string) []int {
	regex := regexp.MustCompile(`[^0-9]+`)
	var numbers []int
	if len(input) == 0 {
		return numbers
	}

	input = strings.Trim(input, " ")

	for _, value := range regex.Split(input, -1) {
		number, _ := strconv.Atoi(value)
		numbers = append(numbers, number)
	}
	return numbers
}
