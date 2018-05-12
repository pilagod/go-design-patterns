package strategy

import (
	"io"
)

type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type PrintOutput struct {
	LogWriter io.Writer
	Writer    io.Writer
}

func (po *PrintOutput) SetLog(lw io.Writer) {
	po.LogWriter = lw
}

func (po *PrintOutput) SetWriter(w io.Writer) {
	po.Writer = w
}
