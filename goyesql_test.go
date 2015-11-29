package goyesql

import (
	"testing"
)

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
