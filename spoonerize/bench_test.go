package spoonerize

import (
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Spoonerize([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'})
	}
	b.ReportAllocs()
}
