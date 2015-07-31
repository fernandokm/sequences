// Generated by: main
// TypeWriter: sequenceGenerator
// Directive: +gen on FibonacciGenerator

package sequences

import "math/big"

// Generated by gen with typewriter at github.com/fernandokm/sequences/writers
// GNU GPL 2.0
// Copyright (c) 2015 Fernando Matsumoto (ferkmatsumoto@hotmail.com)

func (generator *FibonacciGenerator) UpToInt64(maxValue int64) <-chan int64 {
	gen := generator.newInt64Generator()
	ch := make(chan int64, generator.bufferSize)
	go func() {
		for {
			val := gen()

			if val > maxValue {

				close(ch)
				return
			}
			ch <- val
		}
	}()
	return ch
}

func (generator *FibonacciGenerator) TakeInt64(maxIndex int) <-chan int64 {
	gen := generator.newInt64Generator()
	ch := make(chan int64, generator.bufferSize)
	go func() {
		for i := 0; i < maxIndex; i++ {
			ch <- gen()
		}
		close(ch)

	}()
	return ch
}

func (generator *FibonacciGenerator) WhileInt64(condition func(index int, value int64) bool) <-chan int64 {
	gen := generator.newInt64Generator()
	ch := make(chan int64, generator.bufferSize)
	go func() {
		i := 0
		for {
			val := gen()
			if !condition(i, val) {
				close(ch)
				return
			}
			ch <- val
			i++
		}
	}()
	return ch
}

func (generator *FibonacciGenerator) UpToUint64(maxValue uint64) <-chan uint64 {
	gen := generator.newUint64Generator()
	ch := make(chan uint64, generator.bufferSize)
	go func() {
		for {
			val := gen()

			if val > maxValue {

				close(ch)
				return
			}
			ch <- val
		}
	}()
	return ch
}

func (generator *FibonacciGenerator) TakeUint64(maxIndex int) <-chan uint64 {
	gen := generator.newUint64Generator()
	ch := make(chan uint64, generator.bufferSize)
	go func() {
		for i := 0; i < maxIndex; i++ {
			ch <- gen()
		}
		close(ch)

	}()
	return ch
}

func (generator *FibonacciGenerator) WhileUint64(condition func(index int, value uint64) bool) <-chan uint64 {
	gen := generator.newUint64Generator()
	ch := make(chan uint64, generator.bufferSize)
	go func() {
		i := 0
		for {
			val := gen()
			if !condition(i, val) {
				close(ch)
				return
			}
			ch <- val
			i++
		}
	}()
	return ch
}

func (generator *FibonacciGenerator) UpToBigInt(maxValue *big.Int) <-chan *big.Int {
	gen := generator.newBigIntGenerator()
	ch := make(chan *big.Int, generator.bufferSize)
	go func() {
		for {
			val := gen()

			if val.Cmp(maxValue) > 0 {

				close(ch)
				return
			}
			ch <- val
		}
	}()
	return ch
}

func (generator *FibonacciGenerator) TakeBigInt(maxIndex int) <-chan *big.Int {
	gen := generator.newBigIntGenerator()
	ch := make(chan *big.Int, generator.bufferSize)
	go func() {
		for i := 0; i < maxIndex; i++ {
			ch <- gen()
		}
		close(ch)

	}()
	return ch
}

func (generator *FibonacciGenerator) WhileBigInt(condition func(index int, value *big.Int) bool) <-chan *big.Int {
	gen := generator.newBigIntGenerator()
	ch := make(chan *big.Int, generator.bufferSize)
	go func() {
		i := 0
		for {
			val := gen()
			if !condition(i, val) {
				close(ch)
				return
			}
			ch <- val
			i++
		}
	}()
	return ch
}
