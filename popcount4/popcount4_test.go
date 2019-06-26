package popcount4_test

import (
	"go_exercises/popcount4"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount4.PopCount(0x1234567890ABCDEF)
	}
}
