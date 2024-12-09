package applog

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application takes a string representing a log line and returns name of identified application.
// If application is not identified it returns a default application name.
func Application(log string) string {
	for _, chr := range log {
		switch fmt.Sprintf("%U", chr) {
		case "U+2757":
			return "recommendation"
		case "U+1F50D":
			return "search"
		case "U+2600":
			return "weather"
		default:
		}
	}
	return "default"
}

// Replace takes a log string and replaces all occurrences of old rune with new.
// It returns a modified log line.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit checks if number of characters in provided log line is within the given limit.
func WithinLimit(log string, limit int) bool {
	return limit >= utf8.RuneCountInString(log)
}
