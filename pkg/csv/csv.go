// pkg/csv/csv.go
package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/esirangelomub/blankfactor-double-booked/pkg/event"
)

func ParseEventsFromCSV(dirPath string) ([]event.Event, error) {
	var events []event.Event

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".csv" && strings.HasPrefix(file.Name(), "events-waiting-") {
			filePath := filepath.Join(dirPath, file.Name())
			fmt.Printf("Processing file: %s\n", file.Name())
			fileEvents, err := parseEventsFromCSVFile(filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to parse events from file %s: %w", file.Name(), err)
			}
			events = append(events, fileEvents...)
		}
	}

	fmt.Printf("Total events parsed: %d\n", len(events))

	return events, nil
}

func parseEventsFromCSVFile(filePath string) ([]event.Event, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	var events []event.Event

	// Skip the header row by starting the loop from index 1
	for i, record := range records {
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse ID: %w", err)
		}
		startTime, err := time.Parse(time.RFC3339, record[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse StartTime: %w", err)
		}

		endTime, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse EndTime: %w", err)
		}

		event := event.Event{
			ID:        id,
			StartTime: startTime,
			EndTime:   endTime,
		}

		events = append(events, event)
	}

	return events, nil
}
