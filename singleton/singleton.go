package singleton

// Counter is an interface of singleton counter
type Counter interface {
	GetCurrentCount() int
	Increase()
}

type counter struct {
	count int
}

func (counter *counter) GetCurrentCount() int {
	return counter.count
}

// Increase add count in Counter by 1
func (counter *counter) Increase() {
	counter.count++
}

var instance *counter

// GetInstance always return a same counter object
func GetInstance() Counter {
	if instance == nil {
		instance = new(counter)
	}
	return instance
}
