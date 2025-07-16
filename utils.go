package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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

	if len(records) == 0 {
		panic("Error parsing CSV or no records found.")
	}
	return records
}

func printJson(data WeatherData) {
	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(jsonBytes))
}
