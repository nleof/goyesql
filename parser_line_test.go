package goyesql

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := map[string]parsedLine{
		" ":              parsedLine{lineBlank, ""},
		" SELECT * ":     parsedLine{lineQuery, "SELECT *"},
		" -- name: tag ": parsedLine{lineTag, "tag"},
		" -- comment ":   parsedLine{lineComment, "comment"},
	}

	for line, expected := range tests {
		parsed := parseLine(line)

		if parsed != expected {
			t.Errorf("Invalid line parsing. Expected '%v', got '%v'", expected, parsed)
		}
	}
}
