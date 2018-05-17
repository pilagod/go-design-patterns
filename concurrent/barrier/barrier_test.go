package barrier

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func captureBarrierOutput(endpoints ...string) string {
	reader, writer, _ := os.Pipe()
	os.Stdout = writer
	out := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	barrier(endpoints...)
	writer.Close()
	return <-out
}

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "User-Agent") {
			t.Error("Barrier should aggregate all responses from endpoints")
		}
		t.Log(result)
	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Error") {
			t.Error("Barrier should return Error string given any one of the request fails")
		}
		t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
		timeoutMilliseconds = 1

		result := captureBarrierOutput(endpoints...)

		if !strings.Contains(result, "Timeout") {
			t.Error("Barrier should handle timeout situation")
		}
		t.Log(result)
	})
}
