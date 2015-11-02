package goyesql

import (
	"os"
)

// Some helpers to read files

// ParseFile read a file and return Queries or an error
func ParseFile(path string) (Queries, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseBuffer(file)
}

// MustParseFile calls Load and panic if an error occurs
func MustParseFile(path string) Queries {
	queries, err := ParseFile(path)
	if err != nil {
		panic(err)
	}

	return queries
}
