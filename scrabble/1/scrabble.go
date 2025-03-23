package scrabble

import "strings"

type Scorer interface {
	Total(s string) int
}

type PointBlock struct {
	letters []rune
	score   int
}

var one = PointBlock{
	letters: []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
	score:   1,
}

var two = PointBlock{
	letters: []rune{'D', 'G'},
	score:   2,
}

var three = PointBlock{
	letters: []rune{'B', 'C', 'M', 'P'},
	score:   3,
}

var four = PointBlock{
	letters: []rune{'F', 'H', 'V', 'W', 'Y'},
	score:   4,
}

var five = PointBlock{
	letters: []rune{'K'},
	score:   5,
}

var eight = PointBlock{
	letters: []rune{'J', 'X'},
	score:   8,
}

var ten = PointBlock{
	letters: []rune{'Q', 'Z'},
	score:   10,
}

func (p PointBlock) Total(s string) int {
	sum := 0
	for _, v := range s {
		for _, lettr := range p.letters {
			if lettr == v {
				sum += p.score
			}
		}
	}
	return sum
}

// Score takes a string representing a word
// and returns calculated score.
func Score(word string) int {
	totalSum := 0
	pointSets := []PointBlock{one, two, three, four, five, eight, ten}
	for _, v := range pointSets {
		totalSum += v.Total(strings.ToUpper(word))
	}
	return totalSum
}
