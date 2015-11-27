package goyesql

import (
	"testing"
)

func TestScannerErrTags(t *testing.T) {
	tests := map[string]error{
		"missing":  ErrTagMissing,
		"doubloon": ErrTagOverwritten,
	}

	for key, expectedErr := range tests {
		_, err := ParseFile("tests/samples/tag_" + key + ".sql")
		if err != expectedErr {
			t.Errorf(
				"A %s tag should return a '%v' error, got '%v'",
				key, expectedErr, err,
			)
		}
	}
}

func TestScannerValid(t *testing.T) {
	file := "tests/samples/valid.sql"

	queries, err := ParseFile(file)
	if err != nil {
		t.Fatal(err)
	}

	expectedQueries := Queries{
		"simple":    "SELECT * FROM simple;",
		"multiline": "SELECT * FROM multiline WHERE line = 42;",
		"comments":  "SELECT * FROM comments;",
	}

	if len(queries) != len(expectedQueries) {
		t.Errorf(
			"%s should return %d requests, got %d",
			file, len(expectedQueries), len(queries),
		)
	}

	for key, expectedQuery := range expectedQueries {
		if queries[key] != expectedQuery {
			t.Errorf(
				"%s query should be '%s', got '%s'",
				key, expectedQuery, queries[key],
			)
		}
	}
}
