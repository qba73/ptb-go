package logparser

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`)
	if err != nil {
		panic(err)
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[~*=-]*>`)
	if err != nil {
		panic(err)
	}
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`(?i)".*password.*"`)
	if err != nil {
		panic(err)
	}
	cnt := 0
	for _, v := range lines {
		if re.MatchString(v) {
			cnt++
		}
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line[\d]+`)
	if err != nil {
		panic(err)
	}
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User[\s]+(\S*)\b`)
	if err != nil {
		panic(err)
	}
	res := make([]string, 0)
	for _, v := range lines {
		mat := re.FindStringSubmatch(v)
		if len(mat) > 1 {
			res = append(res, fmt.Sprintf("[USR] %v %v", mat[1], v))
		} else {
			res = append(res, v)
		}
	}
	return res
}
