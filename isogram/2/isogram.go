package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram takes a string representing a word
// and returns true if the word is isogram or
// false if it is not.
//
// An isogram (also known as a "non-pattern word")
// is a word or phrase without a repeating letter,
// however spaces and hyphens are allowed to
// appear multiple times.
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	word = strings.ToLower(word)
	for _, l := range word {
		if !unicode.IsLetter(l) {
			continue
		}
		_, ok := seen[l]
		if ok {
			return false
		}
		seen[l] = true
	}
	return true
}
