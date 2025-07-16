package main

type WeatherData struct {
	WeatherDataForYear []WeatherDataForYear `json:"WeatherDataForYear"`
}

type WeatherDataForYear struct {
	Year                 string            `json:"Year"`
	FirstRecordedDate    string            `json:"FirstRecordedDate"`
	LastRecordedDate     string            `json:"LastRecordedDate"`
	TotalRainfall        float32           `json:"TotalRainfall"`
	AverageDailyRainfall float32           `json:"AverageDailyRainfall"`
	DaysWithNoRainfall   int16             `json:"DaysWithNoRainfall"`
	DaysWithRainfall     int16             `json:"DaysWithRainfall"`
	MonthlyAggregates    MonthlyAggregates `json:"MonthlyAggregates"`
}

type MonthlyAggregates struct {
	WeatherDataForMonth []WeatherDataForMonth `json:"WeatherDataForMonth"`
}

type WeatherDataForMonth struct {
	Month                string  `json:"Month"`
	FirstRecordedDate    string  `json:"FirstRecordedDate"`
	LastRecordedDate     string  `json:"LastRecordedDate"`
	TotalRainfall        float32 `json:"TotalRainfall"`
	AverageDailyRainfall float32 `json:"AverageDailyRainfall"`
	DaysWithNoRainfall   int16   `json:"DaysWithNoRainfall"`
	DaysWithRainfall     int16   `json:"DaysWithRainfall"`
}

const (
	ProductCode    = 0
	StationNumber  = 1
	Year           = 2
	Month          = 3
	Day            = 4
	RainfallAmount = 5
	RainfallPeriod = 6
	Quality        = 7
)

var Months = map[string]string{
	"01": "January",
	"02": "February",
	"03": "March",
	"04": "April",
	"05": "May",
	"06": "June",
	"07": "July",
	"08": "August",
	"09": "September",
	"10": "October",
	"11": "November",
	"12": "December",
}
