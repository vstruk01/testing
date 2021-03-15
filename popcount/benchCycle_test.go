package popcount

import (
	"testing"
)

func BenchmarkPopCountHard(b *testing.B) {
	for i := uint64(0); i < 1000000; i++ {
		PopCountHard(i)
	}
}

func BenchmarkPopCountCycle(b *testing.B) {
	for i := uint64(0); i < 1000000; i++ {
		PopCountCycle(i)
	}
}

func BenchmarkPopCountEachBit(b *testing.B) {
	for i := uint64(0); i < 1000000; i++ {
		PopCountEachBit(i)
	}
}
