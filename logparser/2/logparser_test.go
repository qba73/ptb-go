package logparser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/logparser"
)

var testsValidLine = []struct {
	desc string
	text string
	want bool
}{
	{
		desc: "Valid ERR message",
		text: "[ERR] A good error here",
		want: true,
	},
	{
		desc: "Valid INF message",
		text: "[INF] The latest information",
		want: true,
	},
	{
		desc: "Invalid ERR message",
		text: "Any old [ERR] text",
		want: false,
	},
	{
		desc: "Invalid INF message",
		text: "bad start to [INF] Message",
		want: false,
	},
	{
		desc: "Invalid tag",
		text: "[BOB] Any old text",
		want: false,
	},
	{
		desc: "Line with less characters than 5",
		text: "text",
		want: false,
	},
	{
		desc: "Empty line",
		text: "",
		want: false,
	},
}

func TestIsValidLine(t *testing.T) {
	for _, tc := range testsValidLine {
		t.Run(tc.desc, func(t *testing.T) {
			got := logparser.IsValidLine(tc.text)
			if tc.want != got {
				t.Errorf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}

var testsSplitLogLine = []struct {
	desc string
	text string
	want []string
}{
	{
		desc: "three sections",
		text: "section 1<*>section 2<~~~>section 3",
		want: []string{"section 1", "section 2", "section 3"},
	},
	{
		desc: "three sections with different symbols inside angular brackets",
		text: "section 1<=>section 2<-*~*->section 3",
		want: []string{"section 1", "section 2", "section 3"},
	},
	{
		desc: "two sections with no characters between angular brackets",
		text: "section 1<>section 2",
		want: []string{"section 1", "section 2"},
	},
	{
		desc: "single section with some angular brackets",
		text: "<start> <end>",
		want: []string{"<start> <end>"},
	},
	{
		desc: "empty text",
		text: "",
		want: []string{""},
	},
}

func TestSplitLogLine(t *testing.T) {
	for _, tc := range testsSplitLogLine {
		t.Run(tc.desc, func(t *testing.T) {
			got := logparser.SplitLogLine(tc.text)
			if !cmp.Equal(tc.want, got) {
				t.Error(cmp.Diff(tc.want, got))
			}
		})
	}
}

func TestCountQuotedPasswords(t *testing.T) {
	tt := []struct {
		desc  string
		lines []string
		want  int
	}{
		{
			desc: "text with two matches",
			lines: []string{
				``,
				`[INF] passWord`,
				`"passWord"`,
				`[INF] User saw error message "Unexpected Error" on page load.`,
				`[INF] The message "Please reset your password" was ignored by the user`,
			},
			want: 2,
		},
		{
			desc: "text with no matches",
			lines: []string{
				`passWord"passWor"`,
			},
			want: 0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := logparser.CountQuotedPasswords(tc.lines)
			if tc.want != got {
				t.Errorf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestRemoveEndOfLineText(t *testing.T) {
	tests := []struct {
		desc string
		text string
		want string
	}{
		{
			desc: "INF message",
			text: "[INF] end-of-line23033 Network Failure end-of-line27",
			want: "[INF]  Network Failure ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := logparser.RemoveEndOfLineText(tt.text)
			want := tt.want
			if want != got {
				t.Fatalf("want: %v, got: %v", want, got)
			}
		})
	}
}

func TestTagWithUserName(t *testing.T) {
	tests := []struct {
		desc  string
		lines []string
		want  []string
	}{
		{
			desc: "INF message",
			lines: []string{
				"[WRN] User James123 has exceeded storage space.",
				"[WRN] Host down. User   Michelle4 lost connection.",
				"[INF] Users can login again after 23:00.",
				"[DBG] We need to check that user names are at least 6 chars long.",
			},
			want: []string{
				"[USR] James123 [WRN] User James123 has exceeded storage space.",
				"[USR] Michelle4 [WRN] Host down. User   Michelle4 lost connection.",
				"[INF] Users can login again after 23:00.",
				"[DBG] We need to check that user names are at least 6 chars long.",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			got := logparser.TagWithUserName(tc.lines)
			if !cmp.Equal(tc.want, got) {
				t.Error(cmp.Diff(tc.want, got))
			}
		})
	}
}
