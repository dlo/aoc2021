package utils

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func LinesFromFile(filename string) []int {
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
