package logparser

import (
	"fmt"
	"regexp"
)

var (
	regValid    = regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	regSplit    = regexp.MustCompile(`\<[(\*\~\=\-)]*\>`)
	regPasswds  = regexp.MustCompile(`(?i)(\".*password\")`)
	regEol      = regexp.MustCompile(`(end\-of\-line\d*)`)
	regUsername = regexp.MustCompile(`User\s+([a-zA-Z0-9]*)`)
)

func IsValidLine(text string) bool {
	return regValid.MatchString(text)
}

func SplitLogLine(text string) []string {
	return regSplit.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	sum := 0
	for _, x := range lines {
		if regPasswds.MatchString(x) {
			sum++
		}
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	return regEol.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	outputLines := lines[:]
	for i, x := range lines {
		match := regUsername.FindStringSubmatch(x)
		if match != nil {
			outputLines[i] = fmt.Sprintf("[USR] %s %s", match[1], x)
		} else {
			outputLines[i] = x
		}
	}
	return outputLines
}
