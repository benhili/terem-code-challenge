package main

import (
	"testing"
)

func TestGetMonthlyAggregates(t *testing.T) {
	// Stub test: just checks function runs without panic
	sampleRecords := [][]string{
		{"IDCJAC0009", "066062", "1858", "01", "01", "", "", ""},
		{"IDCJAC0009", "066062", "1858", "01", "02", "", "", ""},
		{"IDCJAC0009", "066062", "1858", "01", "03", "", "", ""},
		{"IDCJAC0009", "066062", "1858", "01", "04", "", "", ""},
	}
	result := getMonthlyAggregates(sampleRecords)
	if result == nil {
		t.Error("Expected non-nil result from getMonthlyAggregates")
	}
}
