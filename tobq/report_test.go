package report

import (
    "testing"
)

func TestGetLatestReport(t *testing.T) {
    _, err := getLatestReportAllBeaches()
    if err != nil {
	t.Error("Unable to get report:", err)
    }
}

func TestGetLatestReportSpecificBeach(t *testing.T) {
    for i:= 1; i < 12; i++ {
	_, err := getLatestReportSpecificBeach(i)
        if err != nil {
	    t.Error("Unable to get report:", err)
        }
    }
}
