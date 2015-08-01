package sequences

import "math"

// +gen sequenceGenerator:"Iterator[int64,uint64]"
type PrimeGenerator struct {
	bufferSize int
}

func (fg *PrimeGenerator) newInt64Generator() func() int64 {
	gen := fg.newUint64Generator()
	return func() int64 { return int64(gen()) }
}

func (fg *PrimeGenerator) newUint64Generator() func() uint64 {
	primes := make(chan uint64, 1)

	primeList := []uint64{}

	isPrime := func(n uint64) bool {
		max := uint64(math.Sqrt(float64(n)))
		for _, p := range primeList {
			if n%p == 0 {
				return false
			}
			if p > max {
				break
			}
		}
		return true
	}

	addPrime := func(prime uint64) {
		primes <- prime
		primeList = append(primeList, prime)
	}

	tryNumber := func(n uint64) {
		if isPrime(n) {
			addPrime(n)
		}
	}

	go func() {
		for _, p := range []uint64{2, 3, 5, 7, 11, 13, 17} {
			addPrime(p)
		}
		for i := uint64(19); ; i += 2 {
			tryNumber(i)
		}
	}()

	return func() uint64 { return <-primes }
}

func (fg *PrimeGenerator) SetChannelBuffer(size int) {
	if size < 0 {
		panic("Buffer size cannot be less than zero.")
	}
	fg.bufferSize = size
}

func NewPrimeGenerator() *PrimeGenerator {
	return &PrimeGenerator{bufferSize: 30}
}
