package chain

import (
	"fmt"
	"strings"
	"testing"
)

type TestWriter struct {
	msg string
}

func (tw *TestWriter) Write(p []byte) (int, error) {
	tw.msg += string(p)
	return len(p), nil
}

func (tw *TestWriter) Next(s string) {
	tw.Write([]byte(s))
}

func TestCreateDefaultChain(t *testing.T) {
	t.Run("3 loggers, 2 of them writes to console, second only if it founds the word 'hello', third writes to some variable if second fonud 'hello'", func(t *testing.T) {
		testWriter := TestWriter{}

		writeLogger := WriteLogger{Writer: &testWriter}
		secondLogger := SecondLogger{NextLogger: &writeLogger}
		chain := FirstLogger{NextLogger: &secondLogger}

		chain.Next("message that breaks the chain\n")

		if testWriter.msg != "" {
			t.Fatal("Last logger should not receive any message")
		}
		chain.Next("Hello\n")

		if !strings.Contains(testWriter.msg, "Hello") {
			t.Fatal("Last logger didn't receive expected message")
		}
	})

	t.Run("2 loggers, second uses the closure implementation", func(t *testing.T) {
		testWriter := TestWriter{}
		closureLogger := ClosureLogger{
			Closure: func(s string) {
				fmt.Printf("My closure logger! Message: %s\n", s)
				testWriter.msg = s
			},
		}
		chain := FirstLogger{NextLogger: &closureLogger}

		chain.Next("Hello closure logger")

		if !strings.Contains(testWriter.msg, "Hello closure logger") {
			t.Fatal("Expected message wasn't received in testWriter")
		}
	})
}
