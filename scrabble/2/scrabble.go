package scrabble

import "strings"

// an array correlating to ABC's and the values of each letter in scrabble
// capital letter unicode value - 65 = place in array
var points = [26]int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

// Score takes a string representing a word
// and returns calculated score.
func Score(word string) int {
	sum := 0
	// make everything uppercase to make the array work
	workString := strings.ToUpper(word)

	// for each letter, subtract 65 from the unicode value and add the value in that index of the points array
	for _, set := range workString {
		sum += points[set-65]
	}
	return sum
}
