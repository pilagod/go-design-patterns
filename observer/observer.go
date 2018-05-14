package observer

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserverList []Observer
}

func (p *Publisher) AddObserver(o Observer) {
	p.ObserverList = append(p.ObserverList, o)
}

func (p *Publisher) RemoveObserver(o Observer) {
	var target int

	for index, observer := range p.ObserverList {
		if observer == o {
			target = index
			break
		}
	}
	p.ObserverList = append(p.ObserverList[:target], p.ObserverList[target+1:]...)
}

func (p *Publisher) NotifyObservers(msg string) {
	for _, observer := range p.ObserverList {
		observer.Notify(msg)
	}
}
