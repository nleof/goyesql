package goyesql

import (
	"testing"
)

func TestParseFileMissing(t *testing.T) {
	queries, err := ParseFile("tests/samples/missing.sql")

	if err == nil {
		t.Error("Parsing a missing file should return an error")
	}

	if len(queries) > 0 {
		t.Error("Parsing a missing file should not return queries")
	}
}

func TestParseFileExisting(t *testing.T) {
	_, err := ParseFile("tests/samples/valid.sql")
	if err != nil {
		t.Errorf("Parsing a valid file should not return an error, got '%s'", err)
	}
}

func TestMustParseFilePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseFile should panic if an error occurs, got '%s'", r)
		}
	}()
	MustParseFile("tests/samples/missing.sql")
}

func TestMustParseFileNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustParseFile should not panic if no error occurs, got '%s'", r)
		}
	}()
	MustParseFile("tests/samples/valid.sql")
}
