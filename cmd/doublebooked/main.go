package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/esirangelomub/blankfactor-double-booked/pkg/csv"
	"github.com/esirangelomub/blankfactor-double-booked/pkg/event"
	"github.com/esirangelomub/blankfactor-double-booked/pkg/output"
)

func main() {
	run(os.Args, os.Stdout)
}

func run(args []string, w io.Writer) {
	flag.CommandLine.Parse(args[1:])
	if len(flag.Args()) < 1 {
		log.Fatal("Please provide the path to the directory containing the CSV files")
	}

	dirPath := flag.Arg(0)

	events, err := csv.ParseEventsFromCSV(dirPath)
	if err != nil {
		log.Fatalf("Error parsing events from CSV: %v", err)
	}

	sortedEvents := event.SortEventsByStartTime(events)
	overlappingPairs := event.FindOverlappingEventPairs(sortedEvents)

	output.PrintOverlappingPairs(w, overlappingPairs)
}
