package atbash_test

import (
	"testing"

	"github.com/qba73/atbash"
)

func TestEncode(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got := atbash.Encode(tc.phrase)
			if tc.want != got {
				t.Errorf("Cipher('%s'): want '%s', got '%s'", tc.phrase, tc.want, got)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for b.Loop() {
		for _, test := range tt {
			atbash.Encode(test.phrase)
		}
	}
}

var tt = []struct {
	description string
	phrase      string
	want        string
}{
	{
		description: "encode yes",
		phrase:      "yes",
		want:        "bvh",
	}, {
		description: "encode no",
		phrase:      "no",
		want:        "ml",
	}, {
		description: "encode OMG",
		phrase:      "OMG",
		want:        "lnt",
	}, {
		description: "encode spaces",
		phrase:      "O M G",
		want:        "lnt",
	}, {
		description: "encode mindblowingly",
		phrase:      "mindblowingly",
		want:        "nrmwy oldrm tob",
	}, {
		description: "encode numbers",
		phrase:      "Testing,1 2 3, testing.",
		want:        "gvhgr mt123 gvhgr mt",
	}, {
		description: "encode deep thought",
		phrase:      "Truth is fiction.",
		want:        "gifgs rhurx grlm",
	}, {
		description: "encode all the letters",
		phrase:      "The quick brown fox jumps over the lazy dog.",
		want:        "gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt",
	},
}
