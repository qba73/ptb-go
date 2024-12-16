package isogram

import (
	"strings"
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
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		for j := i + 1; j < len(word); j++ {
			if word[i] == ' ' || word[j] == '-' {
				continue
			}
			if word[i] == word[j] {
				return false
			}
		}
	}
	return true
}
