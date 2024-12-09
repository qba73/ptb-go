package isbn

// IsValidISBN takes a string representing ISBN-10 number
// and verifies if the ISBN is valid.
func IsValidISBN(isbn string) bool {
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
