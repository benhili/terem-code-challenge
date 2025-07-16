package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func getWeatherDataForMonth(monthData WeatherDataForMonth, month string, lastDate string) WeatherDataForMonth {
	monthData.LastRecordedDate = lastDate
	monthData.Month = Months[month]
	if monthData.TotalRainfall > 0 {
		monthData.AverageDailyRainfall = monthData.TotalRainfall / float32(monthData.DaysWithRainfall+monthData.DaysWithNoRainfall)
	} else {
		monthData.AverageDailyRainfall = 0
	}

	return monthData
}

func parseMonthlyAggregates(records [][]string) MonthlyAggregates {
	var monthlyAggregates MonthlyAggregates
	var prevMonth string
	var prevDate string
	var currMonthData WeatherDataForMonth

	for idx, record := range records {
		date := record[Year] + "-" + record[Month] + "-" + record[Day]

		if prevMonth != record[Month] {
			if idx != 0 {
				monthlyAggregates.WeatherDataForMonth = append(monthlyAggregates.WeatherDataForMonth, getWeatherDataForMonth(currMonthData, prevMonth, prevDate))
			}

			currMonthData = WeatherDataForMonth{}
			currMonthData.FirstRecordedDate = date
		}

		rainfall, err := strconv.ParseFloat(record[RainfallAmount], 32)
		if err != nil {
			panic(err)
		}

		currMonthData.TotalRainfall += float32(rainfall)
		if rainfall > 0.0 {
			currMonthData.DaysWithRainfall += 1
		} else {
			currMonthData.DaysWithNoRainfall += 1
		}

		prevDate = date
		prevMonth = record[Month]
	}

	monthlyAggregates.WeatherDataForMonth = append(monthlyAggregates.WeatherDataForMonth, getWeatherDataForMonth(currMonthData, prevMonth, prevDate))

	return monthlyAggregates
}

func main() {
	records := parseCsv("bom_minimal.csv")
	if len(records) == 0 {
		println("Error parsing CSV or no records found.")
		return
	}
	monthlyAggregates := parseMonthlyAggregates(records)
	jsonBytes, err := json.MarshalIndent(monthlyAggregates, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(jsonBytes))
}
