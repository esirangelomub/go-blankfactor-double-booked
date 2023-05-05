package utils

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func MustParse(layout, value string) time.Time {
	parsed, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return parsed
}

func CreateSampleCSVFiles(t *testing.T) string {
	// Create a temporary directory with sample CSV files
	tempDir, err := os.MkdirTemp("", "events-waiting-test-")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}

	sampleCSVData := []string{
		"ID,StartTime,EndTime\n1,2023-05-01T10:00:00Z,2023-05-01T11:00:00Z\n2,2023-05-01T12:00:00Z,2023-05-01T13:00:00Z",
	}

	for i, data := range sampleCSVData {
		filename := fmt.Sprintf("events-waiting-%d.csv", i+1)
		filepath := path.Join(tempDir, filename)
		err := os.WriteFile(filepath, []byte(data), 0644)
		if err != nil {
			t.Fatalf("failed to write sample CSV file: %v", err)
		}
	}

	return tempDir
}
