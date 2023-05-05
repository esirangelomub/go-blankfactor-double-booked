// pkg/event/overlap_test.go
package event

import (
	"github.com/esirangelomub/blankfactor-double-booked/pkg/utils"
	"testing"
	"time"
)

func TestFindOverlappingEventPairs(t *testing.T) {
	events := []Event{
		{
			ID:        1,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T10:00:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T11:00:00Z"),
		},
		{
			ID:        2,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T10:30:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T11:30:00Z"),
		},
		{
			ID:        3,
			StartTime: utils.MustParse(time.RFC3339, "2023-05-01T12:00:00Z"),
			EndTime:   utils.MustParse(time.RFC3339, "2023-05-01T13:00:00Z"),
		},
	}

	overlappingPairs := FindOverlappingEventPairs(events)

	expectedOverlappingPairs := [][]Event{
		{
			events[0],
			events[1],
		},
	}

	if len(overlappingPairs) != len(expectedOverlappingPairs) {
		t.Fatalf("Expected %d overlapping pairs, got: %d", len(expectedOverlappingPairs), len(overlappingPairs))
	}

	for i, pair := range overlappingPairs {
		expectedPair := expectedOverlappingPairs[i]

		if len(pair) != len(expectedPair) {
			t.Fatalf("At index %d, expected pair length %d, got: %d", i, len(expectedPair), len(pair))
		}

		for j, event := range pair {
			if event != expectedPair[j] {
				t.Errorf("At index %d, expected event %v, got: %v", i, expectedPair[j], event)
			}
		}
	}
}
