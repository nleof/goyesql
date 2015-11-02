package goyesql

import (
	"regexp"
	"strings"
)

// A line has four possible type
const (
	LineBlank = iota
	LineQuery
	LineComment
	LineTag
)

// ParsedLine is a tuple (Type, Value)
type ParsedLine struct {
	Type  int
	Value string
}

var (
	reTag     *regexp.Regexp
	reComment *regexp.Regexp
)

func init() {
	reTag = regexp.MustCompile("^\\s*--\\s*name\\s*:\\s*(.+)")
	reComment = regexp.MustCompile("^\\s*--\\s*(.+)")
}

func parseLine(line string) ParsedLine {
	line = strings.Trim(line, " ")

	if line == "" {
		return ParsedLine{LineBlank, ""}
	} else if matches := reTag.FindStringSubmatch(line); len(matches) > 0 {
		return ParsedLine{LineTag, matches[1]}
	} else if matches := reComment.FindStringSubmatch(line); len(matches) > 0 {
		return ParsedLine{LineComment, matches[1]}
	}

	return ParsedLine{LineQuery, line}
}
