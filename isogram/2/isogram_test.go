package isogram_test

import (
	"testing"

	"github.com/qba73/isogram"
)

var tt = []struct {
	desc  string
	input string
	want  bool
}{
	{
		desc:  "empty string",
		input: "",
		want:  true,
	},
	{
		desc:  "isogram with only lower case characters",
		input: "isogram",
		want:  true,
	},
	{
		desc:  "word with one duplicated character",
		input: "eleven",
		want:  false,
	},
	{
		desc:  "word with one duplicated character from the end of the alphabet",
		input: "zzyzx",
		want:  false,
	},
	{
		desc:  "longest reported english isogram",
		input: "subdermatoglyphic",
		want:  true,
	},
	{
		desc:  "word with duplicated character in mixed case",
		input: "Alphabet",
		want:  false,
	},
	{
		desc:  "word with duplicated character in mixed case, lowercase first",
		input: "alphAbet",
		want:  false,
	},
	{
		desc:  "hypothetical isogrammic word with hyphen",
		input: "thumbscrew-japingly",
		want:  true,
	},
	{
		desc:  "hypothetical word with duplicated character following hyphen",
		input: "thumbscrew-jappingly",
		want:  false,
	},
	{
		desc:  "isogram with duplicated hyphen",
		input: "six-year-old",
		want:  true,
	},
	{
		desc:  "made-up name that is an isogram",
		input: "Emily Jung Schwartzkopf",
		want:  true,
	},
	{
		desc:  "duplicated character in the middle",
		input: "accentor",
		want:  false,
	},
	{
		desc:  "same first and last characters",
		input: "angola",
		want:  false,
	},
	{
		desc:  "word with duplicated character and with two hyphens",
		input: "up-to-date",
		want:  false,
	},
}

func TestIsIsogram(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := isogram.IsIsogram(tc.input)
			if tc.want != got {
				t.Errorf("IsIsogram(%q) = %t, want: %t", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range tt {
			isogram.IsIsogram(tc.input)
		}
	}
}
