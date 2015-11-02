package goyesql

import (
	"bufio"
	"errors"
	"io"
)

var (
	// ErrTagMissing occurs when a query has no tag
	ErrTagMissing = errors.New("Query without tag")

	// ErrTagOverwritten occurs when a tag is overwritten by a new one
	ErrTagOverwritten = errors.New("Tag overwritten")
)

// Tag is a string
type Tag string

// Query is a string
type Query string

// Queries is a map associating a Tag to a Query
type Queries map[Tag]Query

func parseBuffer(reader io.Reader) (Queries, error) {
	var (
		lastTag  Tag
		lastLine ParsedLine
	)

	queries := make(Queries)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := parseLine(scanner.Text())

		switch line.Type {
		case LineBlank, LineComment:
			// we don't care about blank and comment lines
			continue

		case LineQuery:
			// got a query but no tag before
			if lastTag == "" {
				return nil, ErrTagMissing
			}

			query := line.Value
			// if query is multiline
			if queries[lastTag] != "" {
				query = " " + query
			}
			queries[lastTag] += Query(query)

		case LineTag:
			// got a tag after another tag
			if lastLine.Type == LineTag {
				return nil, ErrTagOverwritten
			}

			lastTag = Tag(line.Value)
		}

		lastLine = line
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return queries, nil
}
