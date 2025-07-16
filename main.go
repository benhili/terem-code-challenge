/*
{
 "WeatherData": {
 "WeatherDataForYear": {
 "Year": "2019",
 "FirstRecordedDate": "2019-01-01",
 "LastRecordedDate": "2019-04-19",
 "TotalRainfall": "374.2",
 "AverageDailyRainfall": "3.433027523",
 "DaysWithNoRainfall": "65",
 "DaysWithRainfall": "44",
 "MonthlyAggregates": {
 "WeatherDataForMonth": {
 "Month": "January",
 "FirstRecordedDate": "2019-01-01",
 "LastRecordedDate": "2019-01-31",
 "TotalRainfall": "48.8",
 "AverageDailyRainfall": "1.574193548",
 "DaysWithNoRainfall": "21",
 "DaysWithRainfall": "10"
 }
 }
 }
 }
}
*/

package main

import (
	"encoding/csv"
	"os"
)

func getMonthlyAggregates(records [][]string) []map[string]string {
	var aggregates []map[string]string
	// for _, record := range records {

	// }
	return aggregates
}

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

func main() {
	records := parseCsv("IDCJAC0009_066062_1800_Data.csv")
	if len(records) == 0 {
		println("Error parsing CSV or no records found.")
		return
	}
	// for _, record := range records {
	// }
}
