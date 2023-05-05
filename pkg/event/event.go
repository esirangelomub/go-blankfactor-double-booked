// pkg/event/event.go
package event

import (
	"sort"
	"time"
)

type Event struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
}

func SortEventsByStartTime(events []Event) []Event {
	sortedEvents := make([]Event, len(events))
	copy(sortedEvents, events)

	sort.Slice(sortedEvents, func(i, j int) bool {
		return sortedEvents[i].StartTime.Before(sortedEvents[j].StartTime)
	})

	return sortedEvents
}

func FindOverlappingEventPairs(events []Event) [][]Event {
	var overlappingPairs [][]Event

	for i := 0; i < len(events)-1; i++ {
		event1 := events[i]
		event2 := events[i+1]

		if event1.EndTime.After(event2.StartTime) {
			overlappingPairs = append(overlappingPairs, []Event{event1, event2})
		}
	}

	return overlappingPairs
}
