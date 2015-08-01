// Generated by: main
// TypeWriter: sequenceGenerator
// Directive: +gen on PrimeGenerator

package sequences

// Generated by gen with typewriter at github.com/fernandokm/sequences/writers
// GNU GPL 2.0
// Copyright (c) 2015 Fernando Matsumoto (ferkmatsumoto@hotmail.com)

func (generator *PrimeGenerator) UpToInt64(maxValue int64) <-chan int64 {
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

func (generator *PrimeGenerator) TakeInt64(maxIndex int) <-chan int64 {
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

func (generator *PrimeGenerator) WhileInt64(condition func(index int, value int64) bool) <-chan int64 {
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

func (generator *PrimeGenerator) UpToUint64(maxValue uint64) <-chan uint64 {
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

func (generator *PrimeGenerator) TakeUint64(maxIndex int) <-chan uint64 {
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

func (generator *PrimeGenerator) WhileUint64(condition func(index int, value uint64) bool) <-chan uint64 {
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