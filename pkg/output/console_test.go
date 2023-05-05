// pkg/output/console_test.go
package output

import (
	"bytes"
	"github.com/esirangelomub/blankfactor-double-booked/pkg/event"
	"testing"
	"time"
)

func TestPrintOverlappingPairs(t *testing.T) {
	overlappingPairs := [][]event.Event{
		{
			{ID: 1, StartTime: time.Now(), EndTime: time.Now()},
			{ID: 2, StartTime: time.Now(), EndTime: time.Now()},
		},
		{
			{ID: 2, StartTime: time.Now(), EndTime: time.Now()},
			{ID: 3, StartTime: time.Now(), EndTime: time.Now()},
		},
	}

	var buf bytes.Buffer
	PrintOverlappingPairs(&buf, overlappingPairs)

	expected := "Overlapping event pairs: [ [1, 2], [2, 3] ]\n"
	got := buf.String()

	if got != expected {
		t.Errorf("Expected printed output:\n%s\nGot:\n%s", expected, got)
	}
}
