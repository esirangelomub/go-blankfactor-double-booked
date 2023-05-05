// pkg/csv/csv_test.go
package csv

import (
	"github.com/esirangelomub/blankfactor-double-booked/pkg/utils"
	"testing"
	"time"

	"github.com/esirangelomub/blankfactor-double-booked/pkg/event"
)

func TestParseEventsFromCSV(t *testing.T) {
	// Create a temporary directory with sample CSV files
	tempDir := utils.CreateSampleCSVFiles(t)

	events, err := ParseEventsFromCSV(tempDir)
	if err != nil {
		t.Fatalf("ParseEventsFromCSV returned an error: %v", err)
	}

	expectedEvents := []event.Event{
		{
			ID:        1,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T10:00:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T11:00:00Z"),
		},
		{
			ID:        2,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T12:00:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T13:00:00Z"),
		},
	}

	if len(events) != len(expectedEvents) {
		t.Errorf("Expected %d events, got: %d", len(expectedEvents), len(events))
	}

	for i, event := range events {
		if event != expectedEvents[i] {
			t.Errorf("Expected event %v, got: %v", expectedEvents[i], event)
		}
	}
}
