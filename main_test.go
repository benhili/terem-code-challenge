package main

import (
	"encoding/json"
	"testing"
)

func TestParseWeatherData(t *testing.T) {
	// Load and parse CSV
	sampleRecords := parseCsv("bom_minimal.csv")
	result := parseWeatherData(sampleRecords)

	// Example expected JSON
	expectedJSON := `{
        "WeatherDataForYear": [
            {
                "Year": "2019",
                "FirstRecordedDate": "2019-01-01",
                "LastRecordedDate": "2019-12-31",
                "TotalRainfall": 851.8003,
                "AverageDailyRainfall": 2.3336995,
                "DaysWithNoRainfall": 262,
                "DaysWithRainfall": 103,
                "MonthlyAggregates": {
                    "WeatherDataForMonth": [
                        {"Month": "January", "FirstRecordedDate": "2019-01-01", "LastRecordedDate": "2019-01-31", "TotalRainfall": 48.8, "AverageDailyRainfall": 1.5741935, "DaysWithNoRainfall": 21, "DaysWithRainfall": 10},
                        {"Month": "February", "FirstRecordedDate": "2019-02-01", "LastRecordedDate": "2019-02-28", "TotalRainfall": 85.4, "AverageDailyRainfall": 3.05, "DaysWithNoRainfall": 16, "DaysWithRainfall": 12},
                        {"Month": "March", "FirstRecordedDate": "2019-03-01", "LastRecordedDate": "2019-03-31", "TotalRainfall": 229.20001, "AverageDailyRainfall": 7.393549, "DaysWithNoRainfall": 13, "DaysWithRainfall": 18},
                        {"Month": "April", "FirstRecordedDate": "2019-04-01", "LastRecordedDate": "2019-04-30", "TotalRainfall": 11.2, "AverageDailyRainfall": 0.37333333, "DaysWithNoRainfall": 24, "DaysWithRainfall": 6},
                        {"Month": "May", "FirstRecordedDate": "2019-05-01", "LastRecordedDate": "2019-05-31", "TotalRainfall": 14.6, "AverageDailyRainfall": 0.47096774, "DaysWithNoRainfall": 29, "DaysWithRainfall": 2},
                        {"Month": "June", "FirstRecordedDate": "2019-06-01", "LastRecordedDate": "2019-06-30", "TotalRainfall": 171.4, "AverageDailyRainfall": 5.713333, "DaysWithNoRainfall": 15, "DaysWithRainfall": 15},
                        {"Month": "July", "FirstRecordedDate": "2019-07-01", "LastRecordedDate": "2019-07-31", "TotalRainfall": 43.4, "AverageDailyRainfall": 1.4000001, "DaysWithNoRainfall": 24, "DaysWithRainfall": 7},
                        {"Month": "August", "FirstRecordedDate": "2019-08-01", "LastRecordedDate": "2019-08-31", "TotalRainfall": 73, "AverageDailyRainfall": 2.3548386, "DaysWithNoRainfall": 23, "DaysWithRainfall": 8},
                        {"Month": "September", "FirstRecordedDate": "2019-09-01", "LastRecordedDate": "2019-09-30", "TotalRainfall": 112.799995, "AverageDailyRainfall": 3.7599998, "DaysWithNoRainfall": 22, "DaysWithRainfall": 8},
                        {"Month": "October", "FirstRecordedDate": "2019-10-01", "LastRecordedDate": "2019-10-31", "TotalRainfall": 34.2, "AverageDailyRainfall": 1.1032258, "DaysWithNoRainfall": 23, "DaysWithRainfall": 8},
                        {"Month": "November", "FirstRecordedDate": "2019-11-01", "LastRecordedDate": "2019-11-30", "TotalRainfall": 26.199999, "AverageDailyRainfall": 0.8733333, "DaysWithNoRainfall": 23, "DaysWithRainfall": 7},
                        {"Month": "December", "FirstRecordedDate": "2019-12-01", "LastRecordedDate": "2019-12-31", "TotalRainfall": 1.6, "AverageDailyRainfall": 0.051612902, "DaysWithNoRainfall": 29, "DaysWithRainfall": 2}
                    ]
                }
            },
            {
                "Year": "2020",
                "FirstRecordedDate": "2020-01-01",
                "LastRecordedDate": "2020-08-31",
                "TotalRainfall": 1140.4004,
                "AverageDailyRainfall": 4.673772,
                "DaysWithNoRainfall": 127,
                "DaysWithRainfall": 117,
                "MonthlyAggregates": {
                    "WeatherDataForMonth": [
                        {"Month": "January", "FirstRecordedDate": "2020-01-01", "LastRecordedDate": "2020-01-31", "TotalRainfall": 71.399994, "AverageDailyRainfall": 2.3032255, "DaysWithNoRainfall": 16, "DaysWithRainfall": 15},
                        {"Month": "February", "FirstRecordedDate": "2020-02-01", "LastRecordedDate": "2020-02-29", "TotalRainfall": 441.60004, "AverageDailyRainfall": 15.227588, "DaysWithNoRainfall": 13, "DaysWithRainfall": 16},
                        {"Month": "March", "FirstRecordedDate": "2020-03-01", "LastRecordedDate": "2020-03-31", "TotalRainfall": 160.39998, "AverageDailyRainfall": 5.174193, "DaysWithNoRainfall": 10, "DaysWithRainfall": 21},
                        {"Month": "April", "FirstRecordedDate": "2020-04-01", "LastRecordedDate": "2020-04-30", "TotalRainfall": 27.600002, "AverageDailyRainfall": 0.9200001, "DaysWithNoRainfall": 20, "DaysWithRainfall": 10},
                        {"Month": "May", "FirstRecordedDate": "2020-05-01", "LastRecordedDate": "2020-05-31", "TotalRainfall": 118.8, "AverageDailyRainfall": 3.8322582, "DaysWithNoRainfall": 16, "DaysWithRainfall": 15},
                        {"Month": "June", "FirstRecordedDate": "2020-06-01", "LastRecordedDate": "2020-06-30", "TotalRainfall": 69.200005, "AverageDailyRainfall": 2.3066669, "DaysWithNoRainfall": 14, "DaysWithRainfall": 16},
                        {"Month": "July", "FirstRecordedDate": "2020-07-01", "LastRecordedDate": "2020-07-31", "TotalRainfall": 183, "AverageDailyRainfall": 5.903226, "DaysWithNoRainfall": 15, "DaysWithRainfall": 16},
                        {"Month": "August", "FirstRecordedDate": "2020-08-01", "LastRecordedDate": "2020-08-31", "TotalRainfall": 68.399994, "AverageDailyRainfall": 2.2064514, "DaysWithNoRainfall": 23, "DaysWithRainfall": 8}
                    ]
                }
            }
        ]
    }`

	// Marshal result to JSON for comparison
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		t.Fatalf("Failed to marshal result: %v", err)
	}

	// Unmarshal both to interface{} for deep equality check
	var expected, actual interface{}
	if err := json.Unmarshal([]byte(expectedJSON), &expected); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}
	if err := json.Unmarshal(resultJSON, &actual); err != nil {
		t.Fatalf("Failed to unmarshal actual JSON: %v", err)
	}

	if !deepEqual(expected, actual) {
		t.Errorf("Result does not match expected output.\nExpected: %s\nActual: %s", expectedJSON, string(resultJSON))
	}
}

// deepEqual compares two interfaces using json.Marshal for normalization
func deepEqual(a, b interface{}) bool {
	aj, _ := json.Marshal(a)
	bj, _ := json.Marshal(b)
	return string(aj) == string(bj)
}
