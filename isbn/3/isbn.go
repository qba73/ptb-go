package isbn

import (
	"errors"
	"unicode"
)

func IsValidISBN_Version1(isbn string) bool {
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

// IsValidISBN takes a string representing ISBN-10 number
// and verifies if the ISBN is valid.
func IsValidISBN_Version2(isbn string) bool {
	sum, multiplier := 0, 10

	for _, r := range isbn {
		if r >= '0' && r <= '9' {
			sum += int(r-'0') * multiplier
			multiplier--
			continue
		}
		if r == 'X' && multiplier == 1 {
			sum += int(r-'X'+10) * multiplier
			multiplier--
			continue
		}
		if r == '-' {
			continue
		}
		break
	}
	if multiplier != 0 {
		return false
	}
	return sum%11 == 0
}
