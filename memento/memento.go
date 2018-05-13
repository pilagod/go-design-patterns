package memento

import (
	"errors"
)

type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (ct *careTaker) Add(m memento) {
	ct.mementoList = append(ct.mementoList, m)
}

func (ct *careTaker) Memento(i int) (memento, error) {
	if len(ct.mementoList) <= i || i < 0 {
		return memento{}, errors.New("Index not found")
	}
	return ct.mementoList[i], nil
}
