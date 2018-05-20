package observer

type Subscriber interface {
	Notify(interface{}) error
	Close()
}

type Publisher interface {
	AddSubscriberCh() chan<- Subscriber
	RemoveSubscriberCh() chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop()
}
