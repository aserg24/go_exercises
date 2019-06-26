package popcount2_test

import (
	"go_exercises/popcount2"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount2.PopCount(0x1234567890ABCDEF)
	}
}
