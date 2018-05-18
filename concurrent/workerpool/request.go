package workerpool

import (
	"fmt"
	"log"
	"sync"
)

type Request struct {
	Data    interface{}
	Handler func(interface{})
}

func NewStringRequest(id int, str string, wg *sync.WaitGroup) Request {
	return Request{
		Data: "Hello",
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
}
