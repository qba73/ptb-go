package scrabble_test

import (
	"testing"

	"github.com/qba73/scrabble"
)

func TestScore(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := scrabble.Score(tc.input)
			if tc.want != got {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	for b.Loop() {
		for _, tc := range tt {
			_ = scrabble.Score(tc.input)
		}
	}
}

var tt = []struct {
	desc  string
	input string
	want  int
}{
	{
		desc:  "no letter",
		input: "",
		want:  0,
	},
	{
		desc:  "lowercase letter",
		input: "a",
		want:  1,
	},
	{
		desc:  "uppercase letter",
		input: "A",
		want:  1,
	},
	{
		desc:  "valuable letter",
		input: "f",
		want:  4,
	},
	{
		desc:  "short word",
		input: "at",
		want:  2,
	},
	{
		desc:  "short, valuable word",
		input: "zoo",
		want:  12,
	},
	{
		desc:  "medium word",
		input: "street",
		want:  6,
	},
	{
		desc:  "medium, valuable word",
		input: "quirky",
		want:  22,
	},
	{
		desc:  "long, mixed-case word",
		input: "OxyphenButazone",
		want:  41,
	},
	{
		desc:  "english-like word",
		input: "pinata",
		want:  8,
	},
	{
		desc:  "empty input",
		input: "",
		want:  0,
	},
	{
		desc:  "entire alphabet available",
		input: "abcdefghijklmnopqrstuvwxyz",
		want:  87,
	},
}
