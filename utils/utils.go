package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func SumSlice(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

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
	var items []int
	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for _, value := range record {
			result, _ := strconv.Atoi(value)
			items = append(items, result)
		}
	}

	return items
}
