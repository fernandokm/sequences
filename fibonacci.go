// go:generate gen

package sequences

import "math/big"

type bigInt struct{}
type bigRat struct{}

// +gen sequenceGenerator:"Iterator[int64,uint64,*bigInt]"
type FibonacciGenerator struct {
	bufferSize int
}

func (fg *FibonacciGenerator) newInt64Generator() func() int64 {
	var prev, cur int64 = 0, 1
	return func() int64 {
		prev, cur = cur, prev+cur
		return prev
	}
}

func (fg *FibonacciGenerator) newUint64Generator() func() uint64 {
	var prev, cur uint64 = 0, 1
	return func() uint64 {
		prev, cur = cur, prev+cur
		return prev
	}
}

func (fg *FibonacciGenerator) newBigIntGenerator() func() *big.Int {
	prev, cur := big.NewInt(0), big.NewInt(1)
	return func() *big.Int {
		prev, cur = cur, new(big.Int).Add(prev, cur)
		return prev
	}
}

func (fg *FibonacciGenerator) SetChannelBuffer(size int) {
	if size < 0 {
		panic("Buffer size cannot be less than zero.")
	}
	fg.bufferSize = size
}

func NewFibonacciGenerator() *FibonacciGenerator {
	return &FibonacciGenerator{bufferSize: 5}
}
