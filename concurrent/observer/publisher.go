package observer

type publisher struct {
	subscribers []Subscriber
	addSubCh    chan Subscriber
	removeSubCh chan Subscriber
	in          chan interface{}
	stop        chan bool
}

func NewPublisher() Publisher {
	p := &publisher{
		subscribers: make([]Subscriber, 0),
		addSubCh:    make(chan Subscriber),
		removeSubCh: make(chan Subscriber),
		in:          make(chan interface{}),
		stop:        make(chan bool),
	}
	go func() {
		// https://stackoverflow.com/questions/13666253/breaking-out-of-a-select-statement-when-all-channels-are-closed?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
		for {
			select {
			case sub := <-p.addSubCh:
				p.subscribers = append(p.subscribers, sub)

			case sub := <-p.removeSubCh:
				for i, candidtate := range p.subscribers {
					if candidtate == sub {
						p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
						candidtate.Close()
						break
					}
				}

			case msg := <-p.in:
				for _, sub := range p.subscribers {
					sub.Notify(msg)
				}

			case <-p.stop:
				for _, sub := range p.subscribers {
					sub.Close()
				}
				// Don't know why I can't close channel here
				close(p.addSubCh)
				close(p.removeSubCh)
				close(p.in)
				close(p.stop)
				p.stop = nil
			}
			if p.stop == nil {
				return
			}
		}
	}()
	return p
}

func (p *publisher) AddSubscriberCh() chan<- Subscriber {
	return p.addSubCh
}

func (p *publisher) RemoveSubscriberCh() chan<- Subscriber {
	return p.removeSubCh
}

func (p *publisher) PublishingCh() chan<- interface{} {
	return p.in
}

func (p *publisher) Stop() {
	p.stop <- true
}
