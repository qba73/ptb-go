package drops_test

import (
	"testing"

	"github.com/qba73/drops"
)

var tt = []struct {
	desc  string
	input int
	want  string
}{
	{
		desc:  "the sound for 1 is 1",
		input: 1,
		want:  "1",
	},
	{
		desc:  "the sound for 3 is Pling",
		input: 3,
		want:  "Pling",
	},
	{
		desc:  "the sound for 5 is Plang",
		input: 5,
		want:  "Plang",
	},
	{
		desc:  "the sound for 7 is Plong",
		input: 7,
		want:  "Plong",
	},
	{
		desc:  "the sound for 6 is Pling as it has a factor 3",
		input: 6,
		want:  "Pling",
	},
	{
		desc:  "2 to the power 3 does not make a raindrop sound as 3 is the exponent not the base",
		input: 8,
		want:  "8",
	},
	{
		desc:  "the sound for 9 is Pling as it has a factor 3",
		input: 9,
		want:  "Pling",
	},
	{
		desc:  "the sound for 10 is Plang as it has a factor 5",
		input: 10,
		want:  "Plang",
	},
	{
		desc:  "the sound for 14 is Plong as it has a factor of 7",
		input: 14,
		want:  "Plong",
	},
	{
		desc:  "the sound for 15 is PlingPlang as it has factors 3 and 5",
		input: 15,
		want:  "PlingPlang",
	},
	{
		desc:  "the sound for 21 is PlingPlong as it has factors 3 and 7",
		input: 21,
		want:  "PlingPlong",
	},
	{
		desc:  "the sound for 25 is Plang as it has a factor 5",
		input: 25,
		want:  "Plang",
	},
	{
		desc:  "the sound for 27 is Pling as it has a factor 3",
		input: 27,
		want:  "Pling",
	},
	{
		desc:  "the sound for 35 is PlangPlong as it has factors 5 and 7",
		input: 35,
		want:  "PlangPlong",
	},
	{
		desc:  "the sound for 49 is Plong as it has a factor 7",
		input: 49,
		want:  "Plong",
	},
	{
		desc:  "the sound for 52 is 52",
		input: 52,
		want:  "52",
	},
	{
		desc:  "the sound for 105 is PlingPlangPlong as it has factors 3, 5 and 7",
		input: 105,
		want:  "PlingPlangPlong",
	},
	{
		desc:  "the sound for 3125 is Plang as it has a factor 5",
		input: 3125,
		want:  "Plang",
	},
}

func TestConvert(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := drops.Convert(tc.input)
			if got != tc.want {
				t.Errorf("Convert(%d) = %q, want: %q", tc.input, got, tc.want)
			}
		})
	}
}

func BenchmarkConvert(b *testing.B) {
	for b.Loop() {
		for _, tc := range tt {
			drops.Convert(tc.input)
		}
	}
}
