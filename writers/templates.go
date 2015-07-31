package writers

import "github.com/clipperhouse/typewriter"

type generatorModel struct {
	Generator string
	Type      string
	TypeTitle string
}

var templates = typewriter.TemplateSlice{
	intGenerator,
}

var intGenerator = &typewriter.Template{
	Name: "Iterator",
	Text: `
	func (generator *{{.Generator}}) UpTo{{.TypeTitle}}(maxValue {{.Type}}) <-chan {{.Type}} {
		gen := generator.new{{.TypeTitle}}Generator()
		ch := make(chan {{.Type}}, generator.bufferSize)
		go func() {
			for {
				val := gen()
				{{ if or (eq .Type "*big.Int") (eq .Type "*big.Rat") }}
				if val.Cmp(maxValue) > 0 {
				{{ else }}
				if val > maxValue {
				{{ end }}
					close(ch)
					return
				}
				ch <- val
			}
		}()
		return ch
	}

	func (generator *{{.Generator}}) Take{{.TypeTitle}}(maxIndex int) <-chan {{.Type}} {
		gen := generator.new{{.TypeTitle}}Generator()
		ch := make(chan {{.Type}}, generator.bufferSize)
		go func() {
			for i := 0; i < maxIndex; i++{
				ch <- gen()
			}
			close(ch)

		}()
		return ch
	}

	func (generator *{{.Generator}}) While{{.TypeTitle}}(condition func(index int, value {{.Type}}) bool) <- chan {{.Type}} {
		gen := generator.new{{.TypeTitle}}Generator()
		ch := make(chan {{.Type}}, generator.bufferSize)
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
`}
