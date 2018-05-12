package shapes

import (
	"bytes"
	"io"

	"github.com/pilagod/go-design-patterns/strategy"
)

type TextSquare struct {
	strategy.PrintOutput
}

func (ts *TextSquare) Print() error {
	reader := bytes.NewReader([]byte("Square"))
	io.Copy(ts.Writer, reader)
	return nil
}
