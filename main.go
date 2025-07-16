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

// func getMonthlyAggregates(records [][]string) []map[string]string {
// 	aggregates := make(map[string]string)
// 	for idx, record := range records {
// 		if idx == 0 {
// 			aggregates["firstRecordedDate"] = record[1]
// 		}

// 			aggregates["FirstRecordedDate"] = record[1]
// 			aggregates["LastRecordedDate"] = record[2]
// 			aggregates["TotalRainfall"] = record[3]
// 			aggregates["AverageDailyRainfall"] = record[4]
// 			aggregates["DaysWithNoRainfall"] = record[5]
// 			aggregates["DaysWithRainfall"] = record[6]
// 			break
// 		}
// 	}
// 	return aggregates
// }

func main() {
	file, err := os.Open("IDCJAC0009_066062_1800_Data.csv")
	if err != nil {
		println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
	}

	println("Weather data processing application started.")
}
