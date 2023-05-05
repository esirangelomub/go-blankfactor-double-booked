# Double-Booked Events Finder

This program is designed to find overlapping events in multiple CSV files containing calendar event data. It parses the CSV files, sorts the events by start time, finds overlapping event pairs, and outputs the results.

## Prerequisites

- Go 1.19

## Installation

Clone this repository:

```bash
git clone https://github.com/esirangelomub/blankfactor-double-booked.git
cd blankfactor-double-booked
```

## Build the binary:

```bash
go build -o doublebooked cmd/doublebooked/main.go
```

## Usage

The program expects a directory containing CSV files with event data. Each CSV file should have the following format:

```csv
ID,Start Time,End Time
1,2023-05-01T10:00:00Z,2023-05-01T11:00:00Z
2,2023-05-01T12:00:00Z,2023-05-01T13:00:00Z
3,2023-05-01T13:00:00Z,2023-05-01T14:00:00Z
```

* ID: A unique identifier for the event (integer)
* Start Time: The start time of the event (RFC3339 format)
* End Time: The end time of the event (RFC3339 format)

To run the program, use the following command:

```bash
./doublebooked /path/to/csv/directory
```

The program will output a list of overlapping event pairs in the following format:

```bash
Processing file: events-waiting-[DATETIME].csv
Total events parsed: [NUM EVENTS]
Overlapping event pairs: [ [13, 13], [13, 10], ... ]
```

## Testing

To run the unit tests, use the following command:

```bash
go test ./...
or
go test ./... -v
```

To run the unit tests separately, use the following command:

```bash
go test ./pkg/csv
go test ./pkg/event
go test ./pkg/overlap
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
```

