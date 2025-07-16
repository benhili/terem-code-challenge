package main

import (
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

func getWeatherDataForYear(yearData WeatherDataForYear, year string, lastDate string) WeatherDataForYear {
	yearData.LastRecordedDate = lastDate
	yearData.Year = year

	if yearData.TotalRainfall > 0 {
		yearData.AverageDailyRainfall = yearData.TotalRainfall / float32(yearData.DaysWithRainfall+yearData.DaysWithNoRainfall)
	} else {
		yearData.AverageDailyRainfall = 0
	}

	return yearData
}

func getRainfall(rainfallAmount string) float32 {
	if rainfallAmount == "" {
		return 0
	} else {
		parsedRainfall, err := strconv.ParseFloat(rainfallAmount, 32)
		if err != nil {
			panic(err)
		}
		return float32(parsedRainfall)
	}
}

func parseWeatherData(records [][]string) WeatherData {
	var data WeatherData

	var year WeatherDataForYear

	var monthlyAggregates MonthlyAggregates
	var month WeatherDataForMonth

	var prevYear string
	var prevMonth string

	var prevDate string

	for idx, record := range records {
		date := record[Year] + "-" + record[Month] + "-" + record[Day]

		if prevMonth != record[Month] {
			if idx != 0 {
				monthlyAggregates.WeatherDataForMonth = append(monthlyAggregates.WeatherDataForMonth, getWeatherDataForMonth(month, prevMonth, prevDate))
			}

			month = WeatherDataForMonth{}
			month.FirstRecordedDate = date
		}

		if prevYear != record[Year] {
			if idx != 0 {
				year.MonthlyAggregates = monthlyAggregates
				data.WeatherDataForYear = append(data.WeatherDataForYear, getWeatherDataForYear(year, prevYear, prevDate))
			}

			year = WeatherDataForYear{}
			monthlyAggregates = MonthlyAggregates{}
			year.FirstRecordedDate = date
		}

		rainfall := getRainfall(record[RainfallAmount])
		month.TotalRainfall += rainfall
		year.TotalRainfall += rainfall
		if rainfall > 0.0 {
			month.DaysWithRainfall += 1
			year.DaysWithRainfall += 1
		} else {
			month.DaysWithNoRainfall += 1
			year.DaysWithNoRainfall += 1
		}

		prevDate = date
		prevYear = record[Year]
		prevMonth = record[Month]
	}

	monthlyAggregates.WeatherDataForMonth = append(monthlyAggregates.WeatherDataForMonth, getWeatherDataForMonth(month, prevMonth, prevDate))

	year.MonthlyAggregates = monthlyAggregates
	data.WeatherDataForYear = append(data.WeatherDataForYear, getWeatherDataForYear(year, prevYear, prevDate))

	return data
}

func main() {
	records := parseCsv("bom_minimal.csv")
	monthlyAggregates := parseWeatherData(records)
	printJson(monthlyAggregates)
}
