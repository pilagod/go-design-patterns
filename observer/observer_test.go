package observer

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID  int
	Msg string
}

func (to *TestObserver) Notify(msg string) {
	fmt.Printf("Observer %d: message '%s' received\n", to.ID, msg)
	to.Msg = msg
}

func TestPublisher(t *testing.T) {
	observer1 := &TestObserver{1, ""}
	observer2 := &TestObserver{1, ""}
	observer3 := &TestObserver{1, ""}

	t.Run("AddObserver", func(t *testing.T) {
		publisher := Publisher{}

		publisher.AddObserver(observer1)
		publisher.AddObserver(observer2)
		publisher.AddObserver(observer3)

		if len(publisher.ObserverList) != 3 {
			t.Fatal("Publisher should have 3 observers after adding")
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher := Publisher{}

		publisher.AddObserver(observer1)
		publisher.AddObserver(observer2)
		publisher.RemoveObserver(observer2)

		if len(publisher.ObserverList) != 1 {
			t.Fatal("Publisher should have only 1 observer after removing")
		}
		if publisher.ObserverList[0] != observer1 {
			t.Fatal("Publisher should keep only observer 1 in the list")
		}
	})

	t.Run("Notify", func(t *testing.T) {
		publisher := Publisher{}

		publisher.AddObserver(observer1)
		publisher.AddObserver(observer2)
		publisher.AddObserver(observer3)

		msg := "Hello World!"
		publisher.NotifyObservers(msg)

		for _, observer := range publisher.ObserverList {
			observer, _ := observer.(*TestObserver)

			if observer.Msg != msg {
				t.Errorf("observer %d should receive message %s, but got: %s", observer.ID, msg, observer.Msg)
			}
		}
	})
}
