package future

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Fatal("Timeout!")
	wg.Done()
}

func setContext(s string) ExecuteStringFunc {
	msg := fmt.Sprintf("%s Closure!\n", s)

	return func() (string, error) {
		return msg, nil
	}
}

func TestExecuteStringOrError(t *testing.T) {
	future := &MaybeString{}

	t.Run("Success", func(t *testing.T) {
		var wg sync.WaitGroup

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		})
		wg.Add(1)

		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})

	t.Run("Error", func(t *testing.T) {
		var wg sync.WaitGroup

		future.Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})
		wg.Add(1)

		future.Execute(func() (string, error) {
			return "", errors.New("Error occured")
		})
		wg.Wait()
	})

	t.Run("Closure", func(t *testing.T) {
		var wg sync.WaitGroup

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		})
		wg.Add(1)
		future.Execute(setContext("Hello"))
		wg.Wait()
	})
}
