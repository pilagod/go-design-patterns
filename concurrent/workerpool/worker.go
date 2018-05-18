package workerpool

import (
	"fmt"
	"strings"
)

type WorkerLauncher interface {
	LaunchWorker(in <-chan Request)
}

type PrefixSuffixWorker struct {
	prefixStr string
	suffixStr string
}

func (psw *PrefixSuffixWorker) LaunchWorker(in <-chan Request) {
	psw.prefix(psw.append(psw.uppercase(in)))
}

func (psw *PrefixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for req := range in {
			str, ok := req.Data.(string)

			if !ok {
				req.Handler(nil)
				continue
			}
			req.Data = strings.ToUpper(str)
			out <- req
		}
		close(out)
	}()
	return out
}

func (psw *PrefixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for req := range in {
			str, ok := req.Data.(string)

			if !ok {
				req.Handler(nil)
				continue
			}
			req.Data = fmt.Sprintf("%s%s", str, psw.suffixStr)
			out <- req
		}
		close(out)
	}()
	return out
}

func (psw *PrefixSuffixWorker) prefix(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for req := range in {
			str, ok := req.Data.(string)

			if !ok {
				req.Handler(nil)
				continue
			}
			req.Data = fmt.Sprintf("%s%s", psw.prefixStr, str)
			out <- req
		}
		close(out)
	}()
	return out
}
