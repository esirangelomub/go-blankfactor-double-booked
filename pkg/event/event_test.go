// pkg/event/event_test.go
package event

import (
	"github.com/esirangelomub/blankfactor-double-booked/pkg/utils"
	"testing"
	"time"
)

func TestSortEventsByStartTime(t *testing.T) {
	unsortedEvents := []Event{
		{
			ID:        1,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T12:00:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T13:00:00Z"),
		},
		{
			ID:        2,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T10:00:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T11:00:00Z"),
		},
	}

	sortedEvents := SortEventsByStartTime(unsortedEvents)

	expectedSortedEvents := []Event{
		{
			ID:        2,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T10:00:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T11:00:00Z"),
		},
		{
			ID:        1,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T12:00:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T13:00:00Z"),
		},
	}

	if len(sortedEvents) != len(expectedSortedEvents) {
		t.Fatalf("Expected %d events, got: %d", len(expectedSortedEvents), len(sortedEvents))
	}

	for i, event := range sortedEvents {
		if event != expectedSortedEvents[i] {
			t.Errorf("At index %d, expected event %v, got: %v", i, expectedSortedEvents[i], event)
		}
	}
}
