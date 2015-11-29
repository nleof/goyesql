// Package goyesql is a Go port of Yesql
//
// It allows you to write SQL queries in separate files.
//
// See rationale at https://github.com/krisajenkins/yesql#rationale
package goyesql

import (
	"os"
)

// Some helpers to read files

// ParseFile reads a file and return Queries or an error
func ParseFile(path string) (Queries, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseBuffer(file)
}

// MustParseFile calls ParseFile but panic if an error occurs
func MustParseFile(path string) Queries {
	queries, err := ParseFile(path)
	if err != nil {
		panic(err)
	}

	return queries
}
