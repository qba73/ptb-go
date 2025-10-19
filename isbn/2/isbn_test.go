package isbn_test

import (
	"testing"

	"github.com/qba73/isbn"
)

var tt = []struct {
	desc string
	isbn string
	want bool
}{
	{
		desc: "invalid isbn check digit",
		isbn: "3-598-21508-9",
		want: false,
	},
	{
		desc: "check digit is a character other than X",
		isbn: "3-598-21507-A",
		want: false,
	},
	{
		desc: "invalid check digit in isbn is not treated as zero",
		isbn: "4-598-21507-B",
		want: false,
	},
	{
		desc: "invalid character in isbn is not treated as zero",
		isbn: "3-598-P1581-X",
		want: false,
	},
	{
		desc: "X is only valid as a check digit",
		isbn: "3-598-2X507-9",
		want: false,
	},
	{
		desc: "isbn without check digit and dashes",
		isbn: "359821507",
		want: false,
	},
	{
		desc: "too long isbn and no dashes",
		isbn: "3598215078X",
		want: false,
	},
	{
		desc: "too short isbn",
		isbn: "00",
		want: false,
	},
	{
		desc: "isbn without check digit",
		isbn: "3-598-21507",
		want: false,
	},
	{
		desc: "check digit of X should not be used for 0",
		isbn: "3-598-21515-X",
		want: false,
	},
	{
		desc: "empty isbn",
		isbn: "",
		want: false,
	},
	{
		desc: "input is 9 characters",
		isbn: "134456729",
		want: false,
	},
	{
		desc: "invalid characters are not ignored after checking length",
		isbn: "3132P34035",
		want: false,
	},
	{
		desc: "invalid characters are not ignored before checking length",
		isbn: "3598P215088",
		want: false,
	},
	{
		desc: "input is too long but contains a valid isbn",
		isbn: "98245726788",
		want: false,
	},
	{
		desc: "valid isbn",
		isbn: "3-598-21508-8",
		want: true,
	},
	{
		desc: "valid isbn with a check digit of 10",
		isbn: "3-598-21507-X",
		want: true,
	},
	{
		desc: "valid isbn without separating dashes",
		isbn: "3598215088",
		want: true,
	},
	{
		desc: "isbn without separating dashes and X as check digit",
		isbn: "359821507X",
		want: true,
	},
}

func TestIsValidISBN_FailsOnInvalidInput(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := isbn.IsValidISBN(tc.isbn)
			if tc.want != got {
				t.Errorf("IsValidISBN(%q)=%t, want: %t", tc.isbn, got, tc.want)
			}
		})
	}
}

func BenchmarkIsValidISBN(b *testing.B) {
	for b.Loop() {
		for _, n := range tt {
			_ = isbn.IsValidISBN(n.isbn)
		}
	}
}
