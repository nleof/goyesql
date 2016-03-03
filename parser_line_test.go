package goyesql

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := map[string]parsedLine{
		" ":              {lineBlank, ""},
		" SELECT * ":     {lineQuery, "SELECT *"},
		" -- name: tag ": {lineTag, "tag"},
		" -- comment ":   {lineComment, "comment"},
	}

	for line, expected := range tests {
		parsed := parseLine(line)

		if parsed != expected {
			t.Errorf("Invalid line parsing. Expected '%v', got '%v'", expected, parsed)
		}
	}
}
