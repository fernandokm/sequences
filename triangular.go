package sequences

import "math/big"

// +gen sequenceGenerator:"Iterator[int64,uint64,*bigInt]"
type TriangularNumberGenerator struct {
	bufferSize int
}

func (fg *TriangularNumberGenerator) newInt64Generator() func() int64 {
	nextRow := int64(0)
	n := int64(0)
	return func() int64 {
		n += nextRow
		nextRow++
		return n
	}
}

func (fg *TriangularNumberGenerator) newUint64Generator() func() uint64 {
	nextRow := uint64(0)
	n := uint64(0)
	return func() uint64 {
		n += nextRow
		nextRow++
		return n
	}
}

func (fg *TriangularNumberGenerator) newBigIntGenerator() func() *big.Int {
	nextRow := big.NewInt(0)
	n := big.NewInt(0)
	one := big.NewInt(1)
	return func() *big.Int {
		n = new(big.Int).Add(n, nextRow)
		nextRow.Add(nextRow, one)
		return n
	}
}

func (fg *TriangularNumberGenerator) SetChannelBuffer(size int) {
	if size < 0 {
		panic("Buffer size cannot be less than zero.")
	}
	fg.bufferSize = size
}

func NewTriangularGenerator() *TriangularNumberGenerator {
	return &TriangularNumberGenerator{bufferSize: 5}
}
