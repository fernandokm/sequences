package writers

import "github.com/clipperhouse/typewriter"

func init() {
	err := typewriter.Register(NewSequenceGeneratorWriter())
	if err != nil {
		panic(err)
	}
}
