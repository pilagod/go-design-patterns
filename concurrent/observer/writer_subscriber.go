package observer

import (
	"fmt"
	"io"
	"os"
	"time"
)

type writerSubscriber struct {
	id     int
	in     chan interface{}
	Writer io.Writer
}

func NewWriterSubscriber(id int, writer io.Writer) Subscriber {
	if writer == nil {
		writer = os.Stdout
	}
	ws := &writerSubscriber{
		id:     id,
		in:     make(chan interface{}),
		Writer: writer,
	}
	go func() {
		for msg := range ws.in {
			fmt.Fprintf(ws.Writer, "(WriterSubscriber %d): %v\n", ws.id, msg)
		}
	}()
	return ws
}

func (ws *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()
	select {
	case ws.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("Timeout")
	}
	return
}

func (ws *writerSubscriber) Close() {
	close(ws.in)
}
