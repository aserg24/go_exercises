package popcount3_test

import (
	"go_exercises/popcount3"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount3.PopCount(0x1234567890ABCDEF)
	}
}
