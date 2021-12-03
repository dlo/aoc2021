package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
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

func LinesFromFile(filename string) []string {
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
	lines := LinesFromFile(filename)
	var items []int
	for _, value := range lines {
		result, _ := strconv.Atoi(value)
		items = append(items, result)
	}
	return items
}

func BinaryLinesFromFile(filename string) ([]uint64, int) {
	lines := LinesFromFile(filename)
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
