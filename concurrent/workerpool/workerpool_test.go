package workerpool

import (
	"fmt"
	"regexp"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	dispatcher := NewDispatcher(100)

	for i := 0; i < 3; i++ {
		dispatcher.LaunchWorker(&PrefixSuffixWorker{
			prefixStr: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixStr: " World",
		})
	}
	for i := 0; i < 10; i++ {
		dispatcher.MakeRequest(Request{
			Data: fmt.Sprintf("(Msg_id: %d) -> Hello", i),
			Handler: func(i interface{}) {
				str, ok := i.(string)

				if !ok {
					t.Fatal("Invalid casting to string")
				}
				ok, err := regexp.Match(`WorkerID\: \d* -\> \(MSG_ID: \d*\) -> [A-Z]*\sWorld`, []byte(str))

				if !ok || err != nil {
					t.Fatal("Unexpected result from worker")
				}
			},
		})
	}
}
