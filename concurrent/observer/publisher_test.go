package observer

import (
	"sync"
	"testing"
	"time"
)

type mockSubscriber struct {
	notifyTestingFunc func(msg interface{})
	closeTestingFunc  func()
}

func (ms *mockSubscriber) Notify(msg interface{}) error {
	ms.notifyTestingFunc(msg)
	return nil
}

func (ms *mockSubscriber) Close() {
	ms.closeTestingFunc()
}

func TestPublisher(t *testing.T) {
	var wg sync.WaitGroup

	msg := "Hello"
	pub := NewPublisher()
	sub := &mockSubscriber{
		notifyTestingFunc: func(msg interface{}) {
			defer wg.Done()

			s, ok := msg.(string)

			if !ok {
				t.Fatal("Could not assert result")
			}
			if s != msg {
				t.Fail()
			}
		},
		closeTestingFunc: func() {
			wg.Done()
		},
	}
	pub.AddSubscriberCh() <- sub
	wg.Add(1)
	pub.PublishingCh() <- msg
	wg.Wait()

	p := pub.(*publisher)

	if len(p.subscribers) != 1 {
		t.Error("Unexpected number of subscribers")
	}
	pub.RemoveSubscriberCh() <- sub
	wg.Add(1)
	wg.Wait()

	if len(p.subscribers) != 0 {
		t.Error("Expected no subscribers")
	}
	pub.Stop()
	time.Sleep(1 * time.Second)
}
