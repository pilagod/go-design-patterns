package workerpool

import (
	"time"
)

type Dispatcher interface {
	LaunchWorker(w WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	reqch chan Request
}

func NewDispatcher(bufferSize int) Dispatcher {
	return &dispatcher{
		reqch: make(chan Request, bufferSize),
	}
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.reqch)
}

func (d *dispatcher) MakeRequest(req Request) {
	select {
	case d.reqch <- req:
	// introduce a timeout
	case <-time.After(5 * time.Second):
		return
	}
}

func (d *dispatcher) Stop() {
	close(d.reqch)
}
