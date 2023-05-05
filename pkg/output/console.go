// Package output pkg/output/console.go
package output

import (
	"fmt"
	"github.com/esirangelomub/blankfactor-double-booked/pkg/event"
	"io"
	"strings"
)

func PrintOverlappingPairs(w io.Writer, overlappingPairs [][]event.Event) {
	var pairs []string
	for _, pair := range overlappingPairs {
		pairs = append(pairs, fmt.Sprintf("[%d, %d]", pair[0].ID, pair[1].ID))
	}
	fmt.Fprintf(w, "Overlapping event pairs: [ %s ]\n", strings.Join(pairs, ", "))
}
