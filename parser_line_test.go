package goyesql

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := map[string]ParsedLine{
		" ":              ParsedLine{LineBlank, ""},
		" SELECT * ":     ParsedLine{LineQuery, "SELECT *"},
		" -- name: tag ": ParsedLine{LineTag, "tag"},
		" -- comment ":   ParsedLine{LineComment, "comment"},
	}

	for line, expected := range tests {
		parsed := parseLine(line)

		if parsed != expected {
			t.Errorf("Invalid line parsing. Expected '%v', got '%v'", expected, parsed)
		}
	}
}
