package bti_test

import (
	"testing"

	"github.com/qba73/bti"
)

func BenchmarkBoolToInt(b *testing.B) {
	in := []bool{true, false, true, false}

	for b.Loop() {
		for _, i := range in {
			bti.BoolToInt(i)
		}
	}
}

func BenchmarkBoolToInt2(b *testing.B) {
	in := []bool{true, false, true, false}

	for b.Loop() {
		for _, i := range in {
			bti.BoolToInt2(i)
		}
	}
}

func BenchmarkBoolToInt3(b *testing.B) {
	in := []bool{true, false, true, false}

	for b.Loop() {
		for _, i := range in {
			bti.BoolToInt3(i)
		}
	}
}
