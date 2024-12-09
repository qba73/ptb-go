package isbn

import (
	"errors"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	sum := 0
	i := 10
	for _, r := range isbn {
		if r == '-' {
			continue
		}
		v, err := runeToInt(r, i == 1)
		if err != nil {
			return false
		}
		sum += v * i
		i--
	}
	// Size Check
	if i != 0 {
		return false
	}
	return sum%11 == 0
}

func runeToInt(r rune, isCheckDigit bool) (int, error) {
	if isCheckDigit && r == 'X' {
		return 10, nil
	}
	if unicode.IsNumber(r) {
		return int(r - '0'), nil
	}
	return -1, errors.New("invalid rune")
}
