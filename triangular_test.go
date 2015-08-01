package sequences

import "testing"

func BenchmarkTriangularInt64(b *testing.B) {
	tri := NewTriangularGenerator()
	for i := 0; i < b.N; i++ {
		for _ = range tri.TakeInt64(takeCount) {
		}
	}
}

func BenchmarkTriangularUint64(b *testing.B) {
	tri := NewTriangularGenerator()
	for i := 0; i < b.N; i++ {
		for _ = range tri.TakeUint64(takeCount) {
		}
	}
}

func BenchmarkTriangularBigInt(b *testing.B) {
	tri := NewTriangularGenerator()
	for i := 0; i < b.N; i++ {
		for _ = range tri.TakeBigInt(takeCount) {
		}
	}
}

func TestTriangular(t *testing.T) {
	triangulars := []uint64{0, 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, 66, 78, 91,
		105, 120, 136, 153, 171, 190, 210, 231, 253, 276, 300, 325, 351, 378, 406,
	}
	gen := NewTriangularGenerator().UpToUint64(407)
	i := 0
	for val := range gen {
		if triangulars[i] != val {
			t.Fatalf("Expected %d, got %d", triangulars[i], val)
		}
		i++
	}
}
