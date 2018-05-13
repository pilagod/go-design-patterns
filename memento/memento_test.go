package memento

import "testing"

func TestOriginatorExtractAndStoreState(t *testing.T) {
	originator := originator{state: State{Description: "Idle"}}
	activeMemento := memento{state: State{Description: "Active"}}
	originator.ExtractAndStoreState(activeMemento)

	if originator.state.Description != "Active" {
		t.Errorf("Unexpected state found, expected: %s, but got: %s", "Active", originator.state.Description)
	}
}

func TestCareTakerAdd(t *testing.T) {
	originator := originator{}
	originator.state = State{Description: "Idle"}

	mem := originator.NewMemento()

	if mem.state.Description != "Idle" {
		t.Errorf("Expected state was not found")
	}
	careTaker := careTaker{}
	currentLen := len(careTaker.mementoList)
	careTaker.Add(mem)

	if len(careTaker.mementoList) != currentLen+1 {
		t.Errorf("No new elements were added on the list")
	}
}

func TestCareTakerMemento(t *testing.T) {
	originator := originator{}
	originator.state = State{Description: "Idle"}

	careTaker := careTaker{}
	careTaker.Add(originator.NewMemento())

	mem, err := careTaker.Memento(0)

	if err != nil {
		t.Fatal(err)
	}
	if mem.state.Description != "Idle" {
		t.Errorf("Unexpected state, expected: %s, but got: %s", "Idle", mem.state.Description)
	}
	mem, err = careTaker.Memento(-1)

	if err == nil {
		t.Fatal("An error is expected when asking for a negative number")
	}
}
