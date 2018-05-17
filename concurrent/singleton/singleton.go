package singleton

import (
	"sync"
)

type singleton struct {
	count int
	sync.RWMutex
}

func (s *singleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}

func (s *singleton) Stop() {

}

// var addCh = make(chan bool)
// var getCountCh = make(chan chan int)
// var quitCh = make(chan bool)

// func init() {
// 	var count int

// 	go func(addCh <-chan bool, getCountCh <-chan chan int, quitCh <-chan bool) {
// 		for {
// 			select {
// 			case <-addCh:
// 				count++
// 			case ch := <-getCountCh:
// 				ch <- count
// 			case <-quitCh:
// 				return
// 			}
// 		}
// 	}(addCh, getCountCh, quitCh)
// }

// type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

// func (s *singleton) AddOne() {
// 	addCh <- true
// }

// func (s *singleton) GetCount() int {
// 	resCh := make(chan int)
// 	defer close(resCh)
// 	getCountCh <- resCh
// 	return <-resCh
// }

// func (s *singleton) Stop() {
// 	quitCh <- true
// 	close(addCh)
// 	close(getCountCh)
// 	close(quitCh)
// }
