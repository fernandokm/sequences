package sequences

import "testing"

const takeCount = 1e3

func BenchmarkFibInt64(b *testing.B) {
	fib := NewFibonacciGenerator(true)
	for i := 0; i < b.N; i++ {
		for _ = range fib.TakeInt64(takeCount) {
		}
	}
}

func BenchmarkFibUint64(b *testing.B) {
	fib := NewFibonacciGenerator(true)
	for i := 0; i < b.N; i++ {
		for _ = range fib.TakeUint64(takeCount) {
		}
	}
}

func BenchmarkFibBigInt(b *testing.B) {
	fib := NewFibonacciGenerator(true)
	for i := 0; i < b.N; i++ {
		for _ = range fib.TakeBigInt(takeCount) {
		}
	}
}
