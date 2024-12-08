package isogram

import "unicode"

// IsIsogram determines if letters repeat in a word.
func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, l := range word {
		letter := unicode.ToLower(l)
		if letter == '-' || unicode.IsSpace(letter) {
			continue
		}
		if letters[letter] {
			return false
		}
		letters[letter] = true
	}
	return true
}
