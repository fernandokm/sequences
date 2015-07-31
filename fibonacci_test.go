package sequences

import (
	"math/big"
	"testing"
)

const takeCount = 1e3

func BenchmarkInt64(b *testing.B) {
	fib := NewFibonacciGenerator()
	for i := 0; i < b.N; i++ {
		for _ = range fib.UpToInt64(takeCount) {
		}
	}
}

func BenchmarkUint64(b *testing.B) {
	fib := NewFibonacciGenerator()
	for i := 0; i < b.N; i++ {
		for _ = range fib.UpToUint64(takeCount) {
		}
	}
}

func BenchmarkBigInt(b *testing.B) {
	fib := NewFibonacciGenerator()
	count := big.NewInt(takeCount)
	for i := 0; i < b.N; i++ {
		for _ = range fib.UpToBigInt(count) {
		}
	}
}
