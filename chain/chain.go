package chain

import (
	"fmt"
	"io"
	"strings"
)

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextLogger ChainLogger
}

func (fl *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if fl.NextLogger != nil {
		fl.NextLogger.Next(s)
	}
}

type SecondLogger struct {
	NextLogger ChainLogger
}

func (sl *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if sl.NextLogger != nil {
			sl.NextLogger.Next(s)
		}
		return
	}
	fmt.Print("Finishing in second logger\n\n")
}

type WriteLogger struct {
	NextLogger ChainLogger
	Writer     io.Writer
}

func (wl *WriteLogger) Next(s string) {
	if wl.Writer != nil {
		wl.Writer.Write([]byte("WriteLogger: " + s))
	}
	if wl.NextLogger != nil {
		wl.NextLogger.Next(s)
	}
}

type ClosureLogger struct {
	NextLogger ChainLogger
	Closure    func(string)
}

func (cl *ClosureLogger) Next(s string) {
	if cl.Closure != nil {
		cl.Closure(s)
	}
	if cl.NextLogger != nil {
		cl.NextLogger.Next(s)
	}
}
