Code challenge completed for [Terem Technologies](https://terem.tech/)

CLI tool to parse yearly rainfall data from the BOM into JSON

Data sourced at:
http://www.bom.gov.au/jsp/ncc/cdio/weatherData/av?p_nccObsCode=136&p_display_type=dailyDataFile&p_startYear=&p_c=&p_stn_num=066062

Example JSON Format:

```
{
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
          {
            "Month": "January",
            "FirstRecordedDate": "2019-01-01",
            "LastRecordedDate": "2019-01-31",
            "TotalRainfall": 48.8,
            "AverageDailyRainfall": 1.5741935,
            "DaysWithNoRainfall": 21,
            "DaysWithRainfall": 10
          },
          {
            "Month": "February",
            "FirstRecordedDate": "2019-02-01",
            "LastRecordedDate": "2019-02-28",
            "TotalRainfall": 85.4,
            "AverageDailyRainfall": 3.05,
            "DaysWithNoRainfall": 16,
            "DaysWithRainfall": 12
          },
          {
            "Month": "March",
            "FirstRecordedDate": "2019-03-01",
            "LastRecordedDate": "2019-03-31",
            "TotalRainfall": 229.20001,
            "AverageDailyRainfall": 7.393549,
            "DaysWithNoRainfall": 13,
            "DaysWithRainfall": 18
          },
          {
            "Month": "April",
            "FirstRecordedDate": "2019-04-01",
            "LastRecordedDate": "2019-04-30",
            "TotalRainfall": 11.2,
            "AverageDailyRainfall": 0.37333333,
            "DaysWithNoRainfall": 24,
            "DaysWithRainfall": 6
          },
          {
            "Month": "May",
            "FirstRecordedDate": "2019-05-01",
            "LastRecordedDate": "2019-05-31",
            "TotalRainfall": 14.6,
            "AverageDailyRainfall": 0.47096774,
            "DaysWithNoRainfall": 29,
            "DaysWithRainfall": 2
          },
          {
            "Month": "June",
            "FirstRecordedDate": "2019-06-01",
            "LastRecordedDate": "2019-06-30",
            "TotalRainfall": 171.4,
            "AverageDailyRainfall": 5.713333,
            "DaysWithNoRainfall": 15,
            "DaysWithRainfall": 15
          },
          {
            "Month": "July",
            "FirstRecordedDate": "2019-07-01",
            "LastRecordedDate": "2019-07-31",
            "TotalRainfall": 43.4,
            "AverageDailyRainfall": 1.4000001,
            "DaysWithNoRainfall": 24,
            "DaysWithRainfall": 7
          },
          {
            "Month": "August",
            "FirstRecordedDate": "2019-08-01",
            "LastRecordedDate": "2019-08-31",
            "TotalRainfall": 73,
            "AverageDailyRainfall": 2.3548386,
            "DaysWithNoRainfall": 23,
            "DaysWithRainfall": 8
          },
          {
            "Month": "September",
            "FirstRecordedDate": "2019-09-01",
            "LastRecordedDate": "2019-09-30",
            "TotalRainfall": 112.799995,
            "AverageDailyRainfall": 3.7599998,
            "DaysWithNoRainfall": 22,
            "DaysWithRainfall": 8
          },
          {
            "Month": "October",
            "FirstRecordedDate": "2019-10-01",
            "LastRecordedDate": "2019-10-31",
            "TotalRainfall": 34.2,
            "AverageDailyRainfall": 1.1032258,
            "DaysWithNoRainfall": 23,
            "DaysWithRainfall": 8
          },
          {
            "Month": "November",
            "FirstRecordedDate": "2019-11-01",
            "LastRecordedDate": "2019-11-30",
            "TotalRainfall": 26.199999,
            "AverageDailyRainfall": 0.8733333,
            "DaysWithNoRainfall": 23,
            "DaysWithRainfall": 7
          },
          {
            "Month": "December",
            "FirstRecordedDate": "2019-12-01",
            "LastRecordedDate": "2019-12-31",
            "TotalRainfall": 1.6,
            "AverageDailyRainfall": 0.051612902,
            "DaysWithNoRainfall": 29,
            "DaysWithRainfall": 2
          }
        ]
      }
    }
  ]
}
```
