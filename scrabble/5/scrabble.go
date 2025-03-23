package scrabble

import (
	"slices"
	"strings"
)

// Score takes a string represending a scrabble word
// and returns calculated score.
func Score(word string) int {
	score := 0
	one := []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}
	two := []rune{'D', 'G'}
	three := []rune{'B', 'C', 'M', 'P'}
	four := []rune{'F', 'H', 'V', 'W', 'Y'}
	eight := []rune{'J', 'X'}
	ten := []rune{'Q', 'Z'}

	word = strings.ToUpper(word)

	for _, letter := range word {
		if slices.Contains(one, letter) {
			score += 1
		} else if slices.Contains(two, letter) {
			score += 2
		} else if slices.Contains(three, letter) {
			score += 3
		} else if slices.Contains(four, letter) {
			score += 4
		} else if slices.Contains(eight, letter) {
			score += 8
		} else if slices.Contains(ten, letter) {
			score += 10
		} else {
			score += 5
		}
	}
	return score
}
