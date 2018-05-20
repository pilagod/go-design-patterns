package observer

import (
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	testingFunc func(string)
}

func (mw *mockWriter) Write(p []byte) (n int, err error) {
	mw.testingFunc(string(p))
	return len(p), nil
}

func TestWriterSubscriber(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	msg := "Hello"
	subscriber := NewWriterSubscriber(0, &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatalf("Incorrect string: %s", res)
			}
			wg.Done()
		},
	})
	err := subscriber.Notify(msg)

	if err != nil {
		wg.Done()
		t.Error(err)
	}
	wg.Wait()
	subscriber.Close()
}
