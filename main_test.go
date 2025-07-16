package main

import (
	"reflect"
	"testing"
)

var expected = WeatherData{
	WeatherDataForYear: []WeatherDataForYear{
		{
			Year:                 "2019",
			FirstRecordedDate:    "2019-01-01",
			LastRecordedDate:     "2019-12-31",
			TotalRainfall:        851.8003,
			AverageDailyRainfall: 2.3336995,
			DaysWithNoRainfall:   262,
			DaysWithRainfall:     103,
			MonthlyAggregates: MonthlyAggregates{
				WeatherDataForMonth: []WeatherDataForMonth{
					{"January", "2019-01-01", "2019-01-31", 48.8, 1.5741935, 21, 10},
					{"February", "2019-02-01", "2019-02-28", 85.4, 3.05, 16, 12},
					{"March", "2019-03-01", "2019-03-31", 229.20001, 7.393549, 13, 18},
					{"April", "2019-04-01", "2019-04-30", 11.2, 0.37333333, 24, 6},
					{"May", "2019-05-01", "2019-05-31", 14.6, 0.47096774, 29, 2},
					{"June", "2019-06-01", "2019-06-30", 171.4, 5.713333, 15, 15},
					{"July", "2019-07-01", "2019-07-31", 43.4, 1.4000001, 24, 7},
					{"August", "2019-08-01", "2019-08-31", 73, 2.3548386, 23, 8},
					{"September", "2019-09-01", "2019-09-30", 112.799995, 3.7599998, 22, 8},
					{"October", "2019-10-01", "2019-10-31", 34.2, 1.1032258, 23, 8},
					{"November", "2019-11-01", "2019-11-30", 26.199999, 0.8733333, 23, 7},
					{"December", "2019-12-01", "2019-12-31", 1.6, 0.051612902, 29, 2},
				},
			},
		},
		{
			Year:                 "2020",
			FirstRecordedDate:    "2020-01-01",
			LastRecordedDate:     "2020-08-31",
			TotalRainfall:        1140.4004,
			AverageDailyRainfall: 4.673772,
			DaysWithNoRainfall:   127,
			DaysWithRainfall:     117,
			MonthlyAggregates: MonthlyAggregates{
				WeatherDataForMonth: []WeatherDataForMonth{
					{"January", "2020-01-01", "2020-01-31", 71.399994, 2.3032255, 16, 15},
					{"February", "2020-02-01", "2020-02-29", 441.60004, 15.227588, 13, 16},
					{"March", "2020-03-01", "2020-03-31", 160.39998, 5.174193, 10, 21},
					{"April", "2020-04-01", "2020-04-30", 27.600002, 0.9200001, 20, 10},
					{"May", "2020-05-01", "2020-05-31", 118.8, 3.8322582, 16, 15},
					{"June", "2020-06-01", "2020-06-30", 69.200005, 2.3066669, 14, 16},
					{"July", "2020-07-01", "2020-07-31", 183, 5.903226, 15, 16},
					{"August", "2020-08-01", "2020-08-31", 68.399994, 2.2064514, 23, 8},
				},
			},
		},
	},
}

func TestParseWeatherData(t *testing.T) {
	sampleRecords := parseCsv("bom_minimal.csv")
	result := parseWeatherData(sampleRecords)

	if !reflect.DeepEqual(result, expected) {
		t.Error("Result did not equal expected value")
	}
}
