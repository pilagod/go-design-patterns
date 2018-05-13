package visitor

import (
	"fmt"
	"io"
	"os"
)

type Visitable interface {
	Accept(Visitor)
}

type MessageA struct {
	Msg    string
	Output io.Writer
}

func (ma *MessageA) Accept(v Visitor) {
	v.VisitA(ma)
}

func (ma *MessageA) Print() {
	if ma.Output == nil {
		ma.Output = os.Stdout
	}
	fmt.Fprintf(ma.Output, "A: %s", ma.Msg)
}

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (mb *MessageB) Accept(v Visitor) {
	v.VisitB(mb)
}

func (mb *MessageB) Print() {
	if mb.Output == nil {
		mb.Output = os.Stdout
	}
	fmt.Fprintf(mb.Output, "B: %s", mb.Msg)
}

type Visitor interface {
	VisitA(*MessageA)
	VisitB(*MessageB)
}

type MessageVisitor struct{}

func (mv *MessageVisitor) VisitA(ma *MessageA) {
	ma.Msg = fmt.Sprintf("%s %s", ma.Msg, "(Visited A)")
}

func (mv *MessageVisitor) VisitB(mb *MessageB) {
	mb.Msg = fmt.Sprintf("%s %s", mb.Msg, "(Visited B)")
}
