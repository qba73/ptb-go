package applog_test

import (
	"testing"

	"github.com/qba73/applog"
)

func TestApplication(t *testing.T) {
	tt := []struct {
		desc string
		log  string
		want string
	}{
		{
			desc: "single character recommendation",
			log:  "❗ recommended product",
			want: "recommendation",
		},
		{
			desc: "single character search",
			log:  "executed search 🔍",
			want: "search",
		},
		{
			desc: "single character weather",
			log:  "forecast: ☀ sunny",
			want: "weather",
		},
		{
			desc: "no characters default",
			log:  "error: could not proceed",
			want: "default",
		},
		{
			desc: "multiple characters recommendation(1/3)",
			log:  "❗ recommended search product 🔍",
			want: "recommendation",
		},
		{
			desc: "multiple characters recommendation(2/3)",
			log:  "🔍 search recommended product ❗",
			want: "search",
		},
		{
			desc: "multiple characters recommendation(3/3)",
			log:  "☀ weather is sunny ❗",
			want: "weather",
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := applog.Application(tc.log)
			if tc.want != got {
				t.Errorf("Application(\"%s\") = \"%s\", want \"%s\"", tc.log, got, tc.want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	tt := []struct {
		desc    string
		log     string
		oldChar rune
		newChar rune
		want    string
	}{
		{
			desc:    "single occurrence of replacement",
			log:     "❗ recommended product",
			oldChar: '❗',
			newChar: '?',
			want:    "? recommended product",
		},
		{
			desc:    "multiple occurrences of replacement",
			log:     "❗ recommended product ❗",
			oldChar: '❗',
			newChar: '?',
			want:    "? recommended product ?",
		},
		{
			desc:    "no occurrences of replacement",
			log:     "❗ recommended product ❗",
			oldChar: '?',
			newChar: '?',
			want:    "❗ recommended product ❗",
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := applog.Replace(tc.log, tc.oldChar, tc.newChar)
			if tc.want != got {
				t.Errorf("Replace(\"%s\", '%c', '%c') = \"%s\", want \"%s\"", tc.log, tc.oldChar, tc.newChar, got, tc.want)
			}
		})
	}
}

func TestWithinLimit(t *testing.T) {
	tt := []struct {
		desc  string
		log   string
		limit int
		want  bool
	}{
		{
			desc:  "exact limit",
			log:   "exercism❗",
			limit: 9,
			want:  true,
		},
		{
			desc:  "under limit",
			log:   "exercism❗",
			limit: 10,
			want:  true,
		},
		{
			desc:  "over limit",
			log:   "exercism❗",
			limit: 8,
			want:  false,
		},
		{
			desc:  "exact limit",
			log:   "exercism🔍",
			limit: 9,
			want:  true,
		},
		{
			desc:  "under limit",
			log:   "exercism🔍",
			limit: 10,
			want:  true,
		},
		{
			desc:  "over limit",
			log:   "exercism🔍",
			limit: 8,
			want:  false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := applog.WithinLimit(tc.log, tc.limit)
			if tc.want != got {
				t.Errorf("WithinLimit(\"%s\", %d) = %t, want %t", tc.log, tc.limit, got, tc.want)
			}
		})
	}
}
