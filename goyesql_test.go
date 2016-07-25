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

func TestMustParseBytesPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParseBytes should panic if an error occurs, got '%s'", r)
		}
	}()
	MustParseBytes([]byte("I won't work"))
}

func TestMustParseBytesNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustParseBytes should not panic if an error occurs, got '%s'", r)
		}
	}()
	MustParseBytes([]byte("-- name: byte-me\nSELECT * FROM bytes;"))
}

func BenchmarkMustParseFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustParseFile("tests/samples/valid.sql")
	}
}
