package main

import (
	"encoding/csv"
	"os"
)

func parseCsv(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic("Failed to read CSV file: " + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	reader.Read()
	records, err := reader.ReadAll()
	if err != nil {
		panic("Failed to read CSV file: " + err.Error())
	}
	return records
}
