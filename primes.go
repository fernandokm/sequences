package sequences

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

	isPrime := func(n uint64) bool {
		return true
	}

	addPrime := func(prime uint64) {
		primes <- prime
		oldIsPrime := isPrime
		isPrime = func(n uint64) bool {
			return oldIsPrime(n) && (n%prime != 0)
		}
	}

	go func() {
		for _, p := range []uint64{2, 3, 5, 7, 11, 13} {
			addPrime(p)
		}
		for i := uint64(17); ; i += 2 {
			if isPrime(i) {
				addPrime(i)
			}
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
